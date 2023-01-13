// Code generated by mockery v2.14.0. DO NOT EDIT.

package sequencer

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v4"

	state "github.com/0xPolygonHermez/zkevm-node/state"

	time "time"

	types "github.com/ethereum/go-ethereum/core/types"
)

// StateMock is an autogenerated mock type for the stateInterface type
type StateMock struct {
	mock.Mock
}

// Begin provides a mock function with given fields: ctx
func (_m *StateMock) Begin(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	var r0 pgx.Tx
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BeginStateTransaction provides a mock function with given fields: ctx
func (_m *StateMock) BeginStateTransaction(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	var r0 pgx.Tx
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBalance provides a mock function with given fields: ctx, address, root
func (_m *StateMock) GetBalance(ctx context.Context, address common.Address, root []byte) (*big.Int, error) {
	ret := _m.Called(ctx, address, root)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, []byte) *big.Int); ok {
		r0 = rf(ctx, address, root)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, []byte) error); ok {
		r1 = rf(ctx, address, root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBatchByNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateMock) GetBatchByNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) (*state.Batch, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 *state.Batch
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) *state.Batch); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Batch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastStateRoot provides a mock function with given fields: ctx
func (_m *StateMock) GetLastStateRoot(ctx context.Context) ([]byte, error) {
	ret := _m.Called(ctx)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context) []byte); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastVirtualBatchNum provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetLastVirtualBatchNum(ctx context.Context, dbTx pgx.Tx) (uint64, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) uint64); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNonce provides a mock function with given fields: ctx, address, root
func (_m *StateMock) GetNonce(ctx context.Context, address common.Address, root []byte) (*big.Int, error) {
	ret := _m.Called(ctx, address, root)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, []byte) *big.Int); ok {
		r0 = rf(ctx, address, root)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, []byte) error); ok {
		r1 = rf(ctx, address, root)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTimeForLatestBatchVirtualization provides a mock function with given fields: ctx, dbTx
func (_m *StateMock) GetTimeForLatestBatchVirtualization(ctx context.Context, dbTx pgx.Tx) (time.Time, error) {
	ret := _m.Called(ctx, dbTx)

	var r0 time.Time
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx) time.Time); ok {
		r0 = rf(ctx, dbTx)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx) error); ok {
		r1 = rf(ctx, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionsByBatchNumber provides a mock function with given fields: ctx, batchNumber, dbTx
func (_m *StateMock) GetTransactionsByBatchNumber(ctx context.Context, batchNumber uint64, dbTx pgx.Tx) ([]types.Transaction, error) {
	ret := _m.Called(ctx, batchNumber, dbTx)

	var r0 []types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) []types.Transaction); ok {
		r0 = rf(ctx, batchNumber, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNumber, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxsOlderThanNL1Blocks provides a mock function with given fields: ctx, nL1Blocks, dbTx
func (_m *StateMock) GetTxsOlderThanNL1Blocks(ctx context.Context, nL1Blocks uint64, dbTx pgx.Tx) ([]common.Hash, error) {
	ret := _m.Called(ctx, nL1Blocks, dbTx)

	var r0 []common.Hash
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) []common.Hash); ok {
		r0 = rf(ctx, nL1Blocks, dbTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, nL1Blocks, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsBatchClosed provides a mock function with given fields: ctx, batchNum, dbTx
func (_m *StateMock) IsBatchClosed(ctx context.Context, batchNum uint64, dbTx pgx.Tx) (bool, error) {
	ret := _m.Called(ctx, batchNum, dbTx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, uint64, pgx.Tx) bool); ok {
		r0 = rf(ctx, batchNum, dbTx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, pgx.Tx) error); ok {
		r1 = rf(ctx, batchNum, dbTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessBatch provides a mock function with given fields: ctx, request
func (_m *StateMock) ProcessBatch(ctx context.Context, request state.ProcessRequest) (*state.ProcessBatchResponse, error) {
	ret := _m.Called(ctx, request)

	var r0 *state.ProcessBatchResponse
	if rf, ok := ret.Get(0).(func(context.Context, state.ProcessRequest) *state.ProcessBatchResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.ProcessBatchResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, state.ProcessRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStateMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewStateMock creates a new instance of StateMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStateMock(t mockConstructorTestingTNewStateMock) *StateMock {
	mock := &StateMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}