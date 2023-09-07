// package synchronizer
// Implements the logic to retrieve data from L1 and send it to the synchronizer
//   - multiples etherman to do it in parallel
//   - generate blocks to be retrieved
//   - retrieve blocks (parallel)
//   - when reach the update state:
// 		- send a update to channel and  keep retrieving last block to ask for new rollup info
//
//     To control that:
//   - cte: ttlOfLastBlockDefault
//   - when creating object param renewLastBlockOnL1
//
// TODO:
//   - All the stuff related to update last block on L1 could be moved to another class
//   - Check context usage:
//     It need a context to cancel it self and create another context to cancel workers?
//   - Emit metrics
//   - if nothing to update reduce de code to be executed
//   - Improve the unittest of this object
//   - Check all log.fatals to remove it or add status before the panic
//   - Old syncBlocks method try to ask for blocks over last L1 block, I suppose that is for keep
//     synchronizing even a long the synchronization have new blocks. This is not implemented here
//     This is the behaviour of ethman in that situation:
//   - GetRollupInfoByBlockRange returns no error, zero blocks...
//   - EthBlockByNumber returns error:  "not found"

package synchronizer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/log"
	"golang.org/x/exp/constraints"
)

const (
	ttlOfLastBlockDefault                      = time.Second * 5
	timeOutMainLoop                            = time.Minute * 5
	timeForShowUpStatisticsLog                 = time.Second * 60
	conversionFactorPercentage                 = 100
	maxRetriesForRequestnitialValueOfLastBlock = 10
	timeRequestInitialValueOfLastBlock         = time.Second * 5
)

type filter interface {
	toStringBrief() string
	filter(data l1SyncMessage) []l1SyncMessage
}

type syncStatusInterface interface {
	verify(allowModify bool) error
	toStringBrief() string
	getNextRange() *blockRange
	isNodeFullySynchronizedWithL1() bool
	needToRenewLastBlockOnL1() bool
	getStatus() syncStatusEnum
	getLastBlockOnL1() uint64

	onStartedNewWorker(br blockRange)
	onFinishWorker(br blockRange, successful bool)
	onNewLastBlockOnL1(lastBlock uint64) onNewLastBlockResponse
}

type workersInterface interface {
	// verify test params, if allowModify = true allow to change things or make connections
	verify(allowModify bool) error
	// initialize object
	initialize() error
	// finalize object
	finalize() error
	// waits until all workers have finish the current task
	waitFinishAllWorkers()
	asyncRequestRollupInfoByBlockRange(ctx context.Context, blockRange blockRange) (chan genericResponse[responseRollupInfoByBlockRange], error)
	requestLastBlockWithRetries(ctx context.Context, timeout time.Duration, maxPermittedRetries int) genericResponse[retrieveL1LastBlockResult]
	getResponseChannelForRollupInfo() chan genericResponse[responseRollupInfoByBlockRange]
}

type l1RollupInfoProducer struct {
	mutex              sync.Mutex
	ctx                context.Context
	cancelCtx          context.CancelFunc
	workers            workersInterface
	syncStatus         syncStatusInterface
	outgoingChannel    chan l1SyncMessage
	timeLastBLockOnL1  time.Time
	ttlOfLastBlockOnL1 time.Duration
	// filter is an object that sort l1DataMessage to be send ordered by block number
	filterToSendOrdererResultsToConsumer filter
	statistics                           l1RollupInfoProducerStatistics
}

// l1DataRetrieverStatistics : create an instance of l1RollupInfoProducer
func newL1DataRetriever(ctx context.Context, ethermans []EthermanInterface,
	startingBlockNumber uint64, SyncChunkSize uint64,
	outgoingChannel chan l1SyncMessage, renewLastBlockOnL1 bool) *l1RollupInfoProducer {
	if cap(outgoingChannel) < len(ethermans) {
		log.Warnf("l1RollupInfoProducer: outgoingChannel must have a capacity (%d) of at least equal to number of ether clients (%d)", cap(outgoingChannel), len(ethermans))
	}
	ctx, cancel := context.WithCancel(ctx)
	ttlOfLastBlock := ttlOfLastBlockDefault
	if !renewLastBlockOnL1 {
		ttlOfLastBlock = ttlOfLastBlockDefault
	}
	result := l1RollupInfoProducer{
		ctx:                                  ctx,
		cancelCtx:                            cancel,
		syncStatus:                           newSyncStatus(startingBlockNumber, SyncChunkSize, ttlOfLastBlock),
		workers:                              newWorkers(ctx, ethermans),
		filterToSendOrdererResultsToConsumer: newFilterToSendOrdererResultsToConsumer(startingBlockNumber),
		outgoingChannel:                      outgoingChannel,
		statistics:                           newRollupInfoProducerStatistics(startingBlockNumber),
		ttlOfLastBlockOnL1:                   ttlOfLastBlock,
	}
	err := result.verify(false)
	if err != nil {
		log.Fatal(err)
	}
	return &result
}

// TDOO: There is no min function in golang??
func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// verify: test params and status without if not allowModify avoid doing connection or modification of objects
func (l *l1RollupInfoProducer) verify(allowModify bool) error {
	err := l.workers.verify(allowModify)
	if err != nil {
		return err
	}
	err = l.syncStatus.verify(allowModify)
	if err != nil {
		return err
	}
	return nil
}

func (l *l1RollupInfoProducer) initialize() error {
	// TODO: check that all ethermans have the same chainID and get last block in L1
	err := l.verify(true)
	if err != nil {
		log.Fatal(err)
	}
	err = l.workers.initialize()
	if err != nil {
		log.Fatal(err)
	}
	if l.syncStatus.needToRenewLastBlockOnL1() {
		log.Infof("producer: Need a initial value for Last Block On L1, doing the request (maxRetries:%v, timeRequest:%v)",
			maxRetriesForRequestnitialValueOfLastBlock, timeRequestInitialValueOfLastBlock)
		//result := l.retrieveInitialValueOfLastBlock(maxRetriesForRequestnitialValueOfLastBlock, timeRequestInitialValueOfLastBlock)
		result := l.workers.requestLastBlockWithRetries(l.ctx, timeRequestInitialValueOfLastBlock, maxRetriesForRequestnitialValueOfLastBlock)
		if result.err != nil {
			log.Error(result.err)
			return result.err
		}
		l.onNewLastBlock(result.result.block, false)
	}

	return nil
}

func (l *l1RollupInfoProducer) start() error {
	var waitDuration = time.Duration(0)
	previousStatus := l.syncStatus.getStatus()
	for l.step(&waitDuration) {
		newStatus := l.syncStatus.getStatus()
		if previousStatus != newStatus {
			log.Infof("producer: Status changed from [%s] to [%s]", previousStatus.String(), newStatus.String())
			if newStatus == syncStatusSynchronized {
				log.Infof("producer: send a message to consumer to indicate that we are synchronized")
				l.sendPackages([]l1SyncMessage{*newL1SyncMessageControl(eventProducerIsFullySynced)})
			}
		}
		previousStatus = newStatus
	}
	l.workers.waitFinishAllWorkers()
	return nil
}

func (l *l1RollupInfoProducer) step(waitDuration *time.Duration) bool {
	select {
	case <-l.ctx.Done():
		return false
	// That timeout is not need, but just in case that stop launching request
	case <-time.After(*waitDuration):
		log.Debugf("producer: timeout [%s]", *waitDuration)
	case resultRollupInfo := <-l.workers.getResponseChannelForRollupInfo():
		l.onResponseRollupInfo(resultRollupInfo)
	}
	if l.syncStatus.getStatus() == syncStatusSynchronized {
		// Try to nenew last block on L1 if needed
		log.Debugf("producer: status==syncStatusSynchronized -> getting last block on L1")
		l.renewLastBlockOnL1IfNeeded(false)
	}
	// Try to launch retrieve more rollupInfo from L1
	l.launchWork()
	if time.Since(l.statistics.lastShowUpTime) > timeForShowUpStatisticsLog {
		log.Infof("producer: Statistics:%s", l.statistics.getETA())
		l.statistics.lastShowUpTime = time.Now()
	}
	*waitDuration = l.getNextTimeout()
	log.Debugf("producer: Next timeout: %s status: %s", *waitDuration, l.syncStatus.toStringBrief())
	return true
}

func (l *l1RollupInfoProducer) getNextTimeout() time.Duration {
	switch l.syncStatus.getStatus() {
	case syncStatusIdle:
		return timeOutMainLoop
	case syncStatusWorking:
		return timeOutMainLoop
	case syncStatusSynchronized:
		nextRenewLastBlock := time.Since(l.timeLastBLockOnL1) + l.ttlOfLastBlockOnL1
		return max(nextRenewLastBlock, time.Second)
	default:
		log.Fatalf("producer: Unknown status: %s", l.syncStatus.getStatus().String())
	}
	return time.Second
}

// OnNewLastBlock is called when a new last block on L1 is received
func (l *l1RollupInfoProducer) onNewLastBlock(lastBlock uint64, launchWork bool) onNewLastBlockResponse {
	resp := l.syncStatus.onNewLastBlockOnL1(lastBlock)
	l.statistics.updateLastBlockNumber(resp.fullRange.toBlock)
	l.timeLastBLockOnL1 = time.Now()
	if resp.extendedRange != nil {
		log.Infof("producer: New last block on L1: %v -> %s", resp.fullRange.toBlock, resp.toString())
	}
	if launchWork {
		l.launchWork()
	}
	return resp
}

// launchWork: launch new workers if possible and returns new channels created
// returns the number of workers launched
func (l *l1RollupInfoProducer) launchWork() int {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	launchedWorker := 0
	accDebugStr := ""
	thereAreAnError := false
	for {
		br := l.syncStatus.getNextRange()
		if br == nil {
			// No more work to do
			accDebugStr += "[NoNextRange] "
			break
		}
		_, err := l.workers.asyncRequestRollupInfoByBlockRange(l.ctx, *br)
		if err != nil {
			thereAreAnError = true
			accDebugStr += fmt.Sprintf(" segment %s -> [Error:%s] ", br.toString(), err.Error())
			break
		}
		launchedWorker++
		l.syncStatus.onStartedNewWorker(*br)
	}
	if launchedWorker == 0 {
		log.Debugf("producer: No workers launched because: %s", accDebugStr)
	}
	if thereAreAnError {
		log.Warnf("producer: launched workers: %d , but there are an error: %s", launchedWorker, accDebugStr)
	}
	return launchedWorker
}

func (l *l1RollupInfoProducer) renewLastBlockOnL1IfNeeded(forced bool) {
	l.mutex.Lock()
	elapsed := time.Since(l.timeLastBLockOnL1)
	ttl := l.ttlOfLastBlockOnL1
	oldBlock := l.syncStatus.getLastBlockOnL1()
	l.mutex.Unlock()
	if elapsed > ttl || forced {
		log.Infof("producer: Need a new value for Last Block On L1, doing the request")
		result := l.workers.requestLastBlockWithRetries(l.ctx, timeRequestInitialValueOfLastBlock, maxRetriesForRequestnitialValueOfLastBlock)
		log.Infof("producer: Need a new value for Last Block On L1, doing the request old_block:%v -> new block:%v", oldBlock, result.result.block)
		if result.err != nil {
			log.Error(result.err)
			return
		}
		l.onNewLastBlock(result.result.block, true)
	}
}

func (l *l1RollupInfoProducer) onResponseRollupInfo(result genericResponse[responseRollupInfoByBlockRange]) {
	isOk := (result.err == nil)
	l.syncStatus.onFinishWorker(result.result.blockRange, isOk)
	if isOk {
		l.statistics.updateNumRollupInfoOk(1, result.result.blockRange.len())
		outgoingPackages := l.filterToSendOrdererResultsToConsumer.filter(*newL1SyncMessageData(result.result))
		l.sendPackages(outgoingPackages)
	} else {
		l.statistics.updateNumRollupInfoErrors(1)
		log.Warnf("producer: Error while trying to get rollup info by block range: %v", result.err)
	}
}

func (l *l1RollupInfoProducer) stop() {
	l.cancelCtx()
}

func (l *l1RollupInfoProducer) sendPackages(outgoingPackages []l1SyncMessage) {
	for _, pkg := range outgoingPackages {
		log.Infof("producer: Sending results [data] to consumer:%s: It could block channel [%d/%d]", pkg.toStringBrief(), len(l.outgoingChannel), cap(l.outgoingChannel))
		l.outgoingChannel <- pkg
	}
}

// https://stackoverflow.com/questions/4220745/how-to-select-for-input-on-a-dynamic-list-of-channels-in-go
