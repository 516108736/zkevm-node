// Code generated by mockery v2.22.1. DO NOT EDIT.

package synchronizer

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	etherman "github.com/0xPolygonHermez/zkevm-node/etherman"

	mock "github.com/stretchr/testify/mock"

	state "github.com/0xPolygonHermez/zkevm-node/state"

	types "github.com/ethereum/go-ethereum/core/types"
)

// ethermanMock is an autogenerated mock type for the ethermanInterface type
type ethermanMock struct {
	mock.Mock
}

// EthBlockByNumber provides a mock function with given fields: ctx, blockNumber
func (_m *ethermanMock) EthBlockByNumber(ctx context.Context, blockNumber uint64) (*types.Block, error) {
	ret := _m.Called(ctx, blockNumber)

	var r0 *types.Block
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*types.Block, error)); ok {
		return rf(ctx, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *types.Block); ok {
		r0 = rf(ctx, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetForks provides a mock function with given fields: ctx
func (_m *ethermanMock) GetForks(ctx context.Context) ([]state.ForkIDInterval, error) {
	ret := _m.Called(ctx)

	var r0 []state.ForkIDInterval
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]state.ForkIDInterval, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []state.ForkIDInterval); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]state.ForkIDInterval)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBatchNumber provides a mock function with given fields:
func (_m *ethermanMock) GetLatestBatchNumber() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRollupInfoByBlockRange provides a mock function with given fields: ctx, fromBlock, toBlock
func (_m *ethermanMock) GetRollupInfoByBlockRange(ctx context.Context, fromBlock uint64, toBlock *uint64) ([]etherman.Block, map[common.Hash][]etherman.Order, error) {
	ret := _m.Called(ctx, fromBlock, toBlock)

	var r0 []etherman.Block
	var r1 map[common.Hash][]etherman.Order
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *uint64) ([]etherman.Block, map[common.Hash][]etherman.Order, error)); ok {
		return rf(ctx, fromBlock, toBlock)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *uint64) []etherman.Block); ok {
		r0 = rf(ctx, fromBlock, toBlock)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]etherman.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, *uint64) map[common.Hash][]etherman.Order); ok {
		r1 = rf(ctx, fromBlock, toBlock)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[common.Hash][]etherman.Order)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, uint64, *uint64) error); ok {
		r2 = rf(ctx, fromBlock, toBlock)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTrustedSequencerURL provides a mock function with given fields:
func (_m *ethermanMock) GetTrustedSequencerURL() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeaderByNumber provides a mock function with given fields: ctx, number
func (_m *ethermanMock) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ret := _m.Called(ctx, number)

	var r0 *types.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Header, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyGenBlockNumber provides a mock function with given fields: ctx, genBlockNumber
func (_m *ethermanMock) VerifyGenBlockNumber(ctx context.Context, genBlockNumber uint64) (bool, error) {
	ret := _m.Called(ctx, genBlockNumber)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (bool, error)); ok {
		return rf(ctx, genBlockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) bool); ok {
		r0 = rf(ctx, genBlockNumber)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, genBlockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewEthermanMock interface {
	mock.TestingT
	Cleanup(func())
}

// newEthermanMock creates a new instance of ethermanMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newEthermanMock(t mockConstructorTestingTnewEthermanMock) *ethermanMock {
	mock := &ethermanMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
