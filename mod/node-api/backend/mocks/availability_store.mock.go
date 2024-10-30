// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"

	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	mock "github.com/stretchr/testify/mock"
)

// AvailabilityStore is an autogenerated mock type for the AvailabilityStore type
type AvailabilityStore[BeaconBlockBodyT any, BlobSidecarsT any] struct {
	mock.Mock
}

type AvailabilityStore_Expecter[BeaconBlockBodyT any, BlobSidecarsT any] struct {
	mock *mock.Mock
}

func (_m *AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT]) EXPECT() *AvailabilityStore_Expecter[BeaconBlockBodyT, BlobSidecarsT] {
	return &AvailabilityStore_Expecter[BeaconBlockBodyT, BlobSidecarsT]{mock: &_m.Mock}
}

// IsDataAvailable provides a mock function with given fields: _a0, _a1, _a2
func (_m *AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT]) IsDataAvailable(_a0 context.Context, _a1 math.Slot, _a2 BeaconBlockBodyT) bool {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for IsDataAvailable")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, math.Slot, BeaconBlockBodyT) bool); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// AvailabilityStore_IsDataAvailable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsDataAvailable'
type AvailabilityStore_IsDataAvailable_Call[BeaconBlockBodyT any, BlobSidecarsT any] struct {
	*mock.Call
}

// IsDataAvailable is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 math.Slot
//   - _a2 BeaconBlockBodyT
func (_e *AvailabilityStore_Expecter[BeaconBlockBodyT, BlobSidecarsT]) IsDataAvailable(_a0 interface{}, _a1 interface{}, _a2 interface{}) *AvailabilityStore_IsDataAvailable_Call[BeaconBlockBodyT, BlobSidecarsT] {
	return &AvailabilityStore_IsDataAvailable_Call[BeaconBlockBodyT, BlobSidecarsT]{Call: _e.mock.On("IsDataAvailable", _a0, _a1, _a2)}
}

func (_c *AvailabilityStore_IsDataAvailable_Call[BeaconBlockBodyT, BlobSidecarsT]) Run(run func(_a0 context.Context, _a1 math.Slot, _a2 BeaconBlockBodyT)) *AvailabilityStore_IsDataAvailable_Call[BeaconBlockBodyT, BlobSidecarsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(math.Slot), args[2].(BeaconBlockBodyT))
	})
	return _c
}

func (_c *AvailabilityStore_IsDataAvailable_Call[BeaconBlockBodyT, BlobSidecarsT]) Return(_a0 bool) *AvailabilityStore_IsDataAvailable_Call[BeaconBlockBodyT, BlobSidecarsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AvailabilityStore_IsDataAvailable_Call[BeaconBlockBodyT, BlobSidecarsT]) RunAndReturn(run func(context.Context, math.Slot, BeaconBlockBodyT) bool) *AvailabilityStore_IsDataAvailable_Call[BeaconBlockBodyT, BlobSidecarsT] {
	_c.Call.Return(run)
	return _c
}

// Persist provides a mock function with given fields: _a0, _a1
func (_m *AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT]) Persist(_a0 math.Slot, _a1 BlobSidecarsT) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Persist")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(math.Slot, BlobSidecarsT) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AvailabilityStore_Persist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Persist'
type AvailabilityStore_Persist_Call[BeaconBlockBodyT any, BlobSidecarsT any] struct {
	*mock.Call
}

// Persist is a helper method to define mock.On call
//   - _a0 math.Slot
//   - _a1 BlobSidecarsT
func (_e *AvailabilityStore_Expecter[BeaconBlockBodyT, BlobSidecarsT]) Persist(_a0 interface{}, _a1 interface{}) *AvailabilityStore_Persist_Call[BeaconBlockBodyT, BlobSidecarsT] {
	return &AvailabilityStore_Persist_Call[BeaconBlockBodyT, BlobSidecarsT]{Call: _e.mock.On("Persist", _a0, _a1)}
}

func (_c *AvailabilityStore_Persist_Call[BeaconBlockBodyT, BlobSidecarsT]) Run(run func(_a0 math.Slot, _a1 BlobSidecarsT)) *AvailabilityStore_Persist_Call[BeaconBlockBodyT, BlobSidecarsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(math.Slot), args[1].(BlobSidecarsT))
	})
	return _c
}

func (_c *AvailabilityStore_Persist_Call[BeaconBlockBodyT, BlobSidecarsT]) Return(_a0 error) *AvailabilityStore_Persist_Call[BeaconBlockBodyT, BlobSidecarsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AvailabilityStore_Persist_Call[BeaconBlockBodyT, BlobSidecarsT]) RunAndReturn(run func(math.Slot, BlobSidecarsT) error) *AvailabilityStore_Persist_Call[BeaconBlockBodyT, BlobSidecarsT] {
	_c.Call.Return(run)
	return _c
}

// NewAvailabilityStore creates a new instance of AvailabilityStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAvailabilityStore[BeaconBlockBodyT any, BlobSidecarsT any](t interface {
	mock.TestingT
	Cleanup(func())
}) *AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT] {
	mock := &AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
