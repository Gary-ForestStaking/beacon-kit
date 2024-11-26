// Code generated by mockery v2.49.1. DO NOT EDIT.

package mocks

import (
	context "context"

	pruner "github.com/berachain/beacon-kit/mod/storage/pkg/pruner"
	mock "github.com/stretchr/testify/mock"
)

// Pruner is an autogenerated mock type for the Pruner type
type Pruner[PrunableT pruner.Prunable] struct {
	mock.Mock
}

type Pruner_Expecter[PrunableT pruner.Prunable] struct {
	mock *mock.Mock
}

func (_m *Pruner[PrunableT]) EXPECT() *Pruner_Expecter[PrunableT] {
	return &Pruner_Expecter[PrunableT]{mock: &_m.Mock}
}

// Name provides a mock function with given fields:
func (_m *Pruner[PrunableT]) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Pruner_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type Pruner_Name_Call[PrunableT pruner.Prunable] struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *Pruner_Expecter[PrunableT]) Name() *Pruner_Name_Call[PrunableT] {
	return &Pruner_Name_Call[PrunableT]{Call: _e.mock.On("Name")}
}

func (_c *Pruner_Name_Call[PrunableT]) Run(run func()) *Pruner_Name_Call[PrunableT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Pruner_Name_Call[PrunableT]) Return(_a0 string) *Pruner_Name_Call[PrunableT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Pruner_Name_Call[PrunableT]) RunAndReturn(run func() string) *Pruner_Name_Call[PrunableT] {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx
func (_m *Pruner[PrunableT]) Start(ctx context.Context) {
	_m.Called(ctx)
}

// Pruner_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Pruner_Start_Call[PrunableT pruner.Prunable] struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Pruner_Expecter[PrunableT]) Start(ctx interface{}) *Pruner_Start_Call[PrunableT] {
	return &Pruner_Start_Call[PrunableT]{Call: _e.mock.On("Start", ctx)}
}

func (_c *Pruner_Start_Call[PrunableT]) Run(run func(ctx context.Context)) *Pruner_Start_Call[PrunableT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Pruner_Start_Call[PrunableT]) Return() *Pruner_Start_Call[PrunableT] {
	_c.Call.Return()
	return _c
}

func (_c *Pruner_Start_Call[PrunableT]) RunAndReturn(run func(context.Context)) *Pruner_Start_Call[PrunableT] {
	_c.Call.Return(run)
	return _c
}

// NewPruner creates a new instance of Pruner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPruner[PrunableT pruner.Prunable](t interface {
	mock.TestingT
	Cleanup(func())
}) *Pruner[PrunableT] {
	mock := &Pruner[PrunableT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
