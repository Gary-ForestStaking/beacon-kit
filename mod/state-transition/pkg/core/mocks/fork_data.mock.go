// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	common "github.com/berachain/beacon-kit/mod/primitives/pkg/common"

	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

	mock "github.com/stretchr/testify/mock"
)

// ForkData is an autogenerated mock type for the ForkData type
type ForkData[ForkDataT any] struct {
	mock.Mock
}

type ForkData_Expecter[ForkDataT any] struct {
	mock *mock.Mock
}

func (_m *ForkData[ForkDataT]) EXPECT() *ForkData_Expecter[ForkDataT] {
	return &ForkData_Expecter[ForkDataT]{mock: &_m.Mock}
}

// ComputeRandaoSigningRoot provides a mock function with given fields: domainType, epoch
func (_m *ForkData[ForkDataT]) ComputeRandaoSigningRoot(domainType common.DomainType, epoch math.Epoch) common.Root {
	ret := _m.Called(domainType, epoch)

	if len(ret) == 0 {
		panic("no return value specified for ComputeRandaoSigningRoot")
	}

	var r0 common.Root
	if rf, ok := ret.Get(0).(func(common.DomainType, math.Epoch) common.Root); ok {
		r0 = rf(domainType, epoch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Root)
		}
	}

	return r0
}

// ForkData_ComputeRandaoSigningRoot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ComputeRandaoSigningRoot'
type ForkData_ComputeRandaoSigningRoot_Call[ForkDataT any] struct {
	*mock.Call
}

// ComputeRandaoSigningRoot is a helper method to define mock.On call
//   - domainType common.DomainType
//   - epoch math.Epoch
func (_e *ForkData_Expecter[ForkDataT]) ComputeRandaoSigningRoot(domainType interface{}, epoch interface{}) *ForkData_ComputeRandaoSigningRoot_Call[ForkDataT] {
	return &ForkData_ComputeRandaoSigningRoot_Call[ForkDataT]{Call: _e.mock.On("ComputeRandaoSigningRoot", domainType, epoch)}
}

func (_c *ForkData_ComputeRandaoSigningRoot_Call[ForkDataT]) Run(run func(domainType common.DomainType, epoch math.Epoch)) *ForkData_ComputeRandaoSigningRoot_Call[ForkDataT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.DomainType), args[1].(math.Epoch))
	})
	return _c
}

func (_c *ForkData_ComputeRandaoSigningRoot_Call[ForkDataT]) Return(_a0 common.Root) *ForkData_ComputeRandaoSigningRoot_Call[ForkDataT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ForkData_ComputeRandaoSigningRoot_Call[ForkDataT]) RunAndReturn(run func(common.DomainType, math.Epoch) common.Root) *ForkData_ComputeRandaoSigningRoot_Call[ForkDataT] {
	_c.Call.Return(run)
	return _c
}

// New provides a mock function with given fields: _a0, _a1
func (_m *ForkData[ForkDataT]) New(_a0 common.Version, _a1 common.Root) ForkDataT {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for New")
	}

	var r0 ForkDataT
	if rf, ok := ret.Get(0).(func(common.Version, common.Root) ForkDataT); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(ForkDataT)
	}

	return r0
}

// ForkData_New_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'New'
type ForkData_New_Call[ForkDataT any] struct {
	*mock.Call
}

// New is a helper method to define mock.On call
//   - _a0 common.Version
//   - _a1 common.Root
func (_e *ForkData_Expecter[ForkDataT]) New(_a0 interface{}, _a1 interface{}) *ForkData_New_Call[ForkDataT] {
	return &ForkData_New_Call[ForkDataT]{Call: _e.mock.On("New", _a0, _a1)}
}

func (_c *ForkData_New_Call[ForkDataT]) Run(run func(_a0 common.Version, _a1 common.Root)) *ForkData_New_Call[ForkDataT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(common.Version), args[1].(common.Root))
	})
	return _c
}

func (_c *ForkData_New_Call[ForkDataT]) Return(_a0 ForkDataT) *ForkData_New_Call[ForkDataT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ForkData_New_Call[ForkDataT]) RunAndReturn(run func(common.Version, common.Root) ForkDataT) *ForkData_New_Call[ForkDataT] {
	_c.Call.Return(run)
	return _c
}

// NewForkData creates a new instance of ForkData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewForkData[ForkDataT any](t interface {
	mock.TestingT
	Cleanup(func())
}) *ForkData[ForkDataT] {
	mock := &ForkData[ForkDataT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
