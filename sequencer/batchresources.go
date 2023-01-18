package sequencer

import (
	"fmt"

	"github.com/0xPolygonHermez/zkevm-node/state"
)

// batchResources is a struct that contains the ZKEVM resources used by a batch/tx
type batchResources struct {
	zKCounters state.ZKCounters
	bytes      uint64
}

// sub subtracts the batch resources from other
func (r *batchResources) sub(other batchResources) error {
	// Bytes
	if other.bytes > r.bytes {
		return fmt.Errorf("%w. Resource: Bytes", ErrBatchRemainingResourcesUnderflow)
	}
	bytesBackup := r.bytes
	r.bytes -= other.bytes
	err := r.zKCounters.Sub(other.zKCounters)
	if err != nil {
		r.bytes = bytesBackup
		return fmt.Errorf("%w. %s", ErrBatchRemainingResourcesUnderflow, err)
	}

	return err
}
