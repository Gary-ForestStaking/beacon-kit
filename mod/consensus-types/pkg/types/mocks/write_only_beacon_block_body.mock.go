// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	bytes "github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	common "github.com/ethereum/go-ethereum/common"

	eip4844 "github.com/berachain/beacon-kit/mod/primitives/pkg/eip4844"

	mock "github.com/stretchr/testify/mock"

	types "github.com/berachain/beacon-kit/mod/consensus-types/pkg/types/v2"
)

// WriteOnlyBeaconBlockBody is an autogenerated mock type for the WriteOnlyBeaconBlockBody type
type WriteOnlyBeaconBlockBody struct {
	mock.Mock
}

type WriteOnlyBeaconBlockBody_Expecter struct {
	mock *mock.Mock
}

func (_m *WriteOnlyBeaconBlockBody) EXPECT() *WriteOnlyBeaconBlockBody_Expecter {
	return &WriteOnlyBeaconBlockBody_Expecter{mock: &_m.Mock}
}

// SetAttestations provides a mock function with given fields: _a0
func (_m *WriteOnlyBeaconBlockBody) SetAttestations(_a0 []*types.AttestationData) {
	_m.Called(_a0)
}

// WriteOnlyBeaconBlockBody_SetAttestations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetAttestations'
type WriteOnlyBeaconBlockBody_SetAttestations_Call struct {
	*mock.Call
}

// SetAttestations is a helper method to define mock.On call
//   - _a0 []*types.AttestationData
func (_e *WriteOnlyBeaconBlockBody_Expecter) SetAttestations(_a0 interface{}) *WriteOnlyBeaconBlockBody_SetAttestations_Call {
	return &WriteOnlyBeaconBlockBody_SetAttestations_Call{Call: _e.mock.On("SetAttestations", _a0)}
}

func (_c *WriteOnlyBeaconBlockBody_SetAttestations_Call) Run(run func(_a0 []*types.AttestationData)) *WriteOnlyBeaconBlockBody_SetAttestations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]*types.AttestationData))
	})
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetAttestations_Call) Return() *WriteOnlyBeaconBlockBody_SetAttestations_Call {
	_c.Call.Return()
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetAttestations_Call) RunAndReturn(run func([]*types.AttestationData)) *WriteOnlyBeaconBlockBody_SetAttestations_Call {
	_c.Call.Return(run)
	return _c
}

// SetBlobKzgCommitments provides a mock function with given fields: _a0
func (_m *WriteOnlyBeaconBlockBody) SetBlobKzgCommitments(_a0 eip4844.KZGCommitments[common.Hash]) {
	_m.Called(_a0)
}

// WriteOnlyBeaconBlockBody_SetBlobKzgCommitments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetBlobKzgCommitments'
type WriteOnlyBeaconBlockBody_SetBlobKzgCommitments_Call struct {
	*mock.Call
}

// SetBlobKzgCommitments is a helper method to define mock.On call
//   - _a0 eip4844.KZGCommitments[common.Hash]
func (_e *WriteOnlyBeaconBlockBody_Expecter) SetBlobKzgCommitments(_a0 interface{}) *WriteOnlyBeaconBlockBody_SetBlobKzgCommitments_Call {
	return &WriteOnlyBeaconBlockBody_SetBlobKzgCommitments_Call{Call: _e.mock.On("SetBlobKzgCommitments", _a0)}
}

func (_c *WriteOnlyBeaconBlockBody_SetBlobKzgCommitments_Call) Run(run func(_a0 eip4844.KZGCommitments[common.Hash])) *WriteOnlyBeaconBlockBody_SetBlobKzgCommitments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(eip4844.KZGCommitments[common.Hash]))
	})
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetBlobKzgCommitments_Call) Return() *WriteOnlyBeaconBlockBody_SetBlobKzgCommitments_Call {
	_c.Call.Return()
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetBlobKzgCommitments_Call) RunAndReturn(run func(eip4844.KZGCommitments[common.Hash])) *WriteOnlyBeaconBlockBody_SetBlobKzgCommitments_Call {
	_c.Call.Return(run)
	return _c
}

// SetDeposits provides a mock function with given fields: _a0
func (_m *WriteOnlyBeaconBlockBody) SetDeposits(_a0 []*types.Deposit) {
	_m.Called(_a0)
}

// WriteOnlyBeaconBlockBody_SetDeposits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetDeposits'
type WriteOnlyBeaconBlockBody_SetDeposits_Call struct {
	*mock.Call
}

// SetDeposits is a helper method to define mock.On call
//   - _a0 []*types.Deposit
func (_e *WriteOnlyBeaconBlockBody_Expecter) SetDeposits(_a0 interface{}) *WriteOnlyBeaconBlockBody_SetDeposits_Call {
	return &WriteOnlyBeaconBlockBody_SetDeposits_Call{Call: _e.mock.On("SetDeposits", _a0)}
}

func (_c *WriteOnlyBeaconBlockBody_SetDeposits_Call) Run(run func(_a0 []*types.Deposit)) *WriteOnlyBeaconBlockBody_SetDeposits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]*types.Deposit))
	})
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetDeposits_Call) Return() *WriteOnlyBeaconBlockBody_SetDeposits_Call {
	_c.Call.Return()
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetDeposits_Call) RunAndReturn(run func([]*types.Deposit)) *WriteOnlyBeaconBlockBody_SetDeposits_Call {
	_c.Call.Return(run)
	return _c
}

// SetEth1Data provides a mock function with given fields: _a0
func (_m *WriteOnlyBeaconBlockBody) SetEth1Data(_a0 *types.Eth1Data) {
	_m.Called(_a0)
}

// WriteOnlyBeaconBlockBody_SetEth1Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetEth1Data'
type WriteOnlyBeaconBlockBody_SetEth1Data_Call struct {
	*mock.Call
}

// SetEth1Data is a helper method to define mock.On call
//   - _a0 *types.Eth1Data
func (_e *WriteOnlyBeaconBlockBody_Expecter) SetEth1Data(_a0 interface{}) *WriteOnlyBeaconBlockBody_SetEth1Data_Call {
	return &WriteOnlyBeaconBlockBody_SetEth1Data_Call{Call: _e.mock.On("SetEth1Data", _a0)}
}

func (_c *WriteOnlyBeaconBlockBody_SetEth1Data_Call) Run(run func(_a0 *types.Eth1Data)) *WriteOnlyBeaconBlockBody_SetEth1Data_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.Eth1Data))
	})
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetEth1Data_Call) Return() *WriteOnlyBeaconBlockBody_SetEth1Data_Call {
	_c.Call.Return()
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetEth1Data_Call) RunAndReturn(run func(*types.Eth1Data)) *WriteOnlyBeaconBlockBody_SetEth1Data_Call {
	_c.Call.Return(run)
	return _c
}

// SetExecutionData provides a mock function with given fields: _a0
func (_m *WriteOnlyBeaconBlockBody) SetExecutionData(_a0 *types.ExecutionPayload) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SetExecutionData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.ExecutionPayload) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteOnlyBeaconBlockBody_SetExecutionData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetExecutionData'
type WriteOnlyBeaconBlockBody_SetExecutionData_Call struct {
	*mock.Call
}

// SetExecutionData is a helper method to define mock.On call
//   - _a0 *types.ExecutionPayload
func (_e *WriteOnlyBeaconBlockBody_Expecter) SetExecutionData(_a0 interface{}) *WriteOnlyBeaconBlockBody_SetExecutionData_Call {
	return &WriteOnlyBeaconBlockBody_SetExecutionData_Call{Call: _e.mock.On("SetExecutionData", _a0)}
}

func (_c *WriteOnlyBeaconBlockBody_SetExecutionData_Call) Run(run func(_a0 *types.ExecutionPayload)) *WriteOnlyBeaconBlockBody_SetExecutionData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.ExecutionPayload))
	})
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetExecutionData_Call) Return(_a0 error) *WriteOnlyBeaconBlockBody_SetExecutionData_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetExecutionData_Call) RunAndReturn(run func(*types.ExecutionPayload) error) *WriteOnlyBeaconBlockBody_SetExecutionData_Call {
	_c.Call.Return(run)
	return _c
}

// SetGraffiti provides a mock function with given fields: _a0
func (_m *WriteOnlyBeaconBlockBody) SetGraffiti(_a0 bytes.B32) {
	_m.Called(_a0)
}

// WriteOnlyBeaconBlockBody_SetGraffiti_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetGraffiti'
type WriteOnlyBeaconBlockBody_SetGraffiti_Call struct {
	*mock.Call
}

// SetGraffiti is a helper method to define mock.On call
//   - _a0 bytes.B32
func (_e *WriteOnlyBeaconBlockBody_Expecter) SetGraffiti(_a0 interface{}) *WriteOnlyBeaconBlockBody_SetGraffiti_Call {
	return &WriteOnlyBeaconBlockBody_SetGraffiti_Call{Call: _e.mock.On("SetGraffiti", _a0)}
}

func (_c *WriteOnlyBeaconBlockBody_SetGraffiti_Call) Run(run func(_a0 bytes.B32)) *WriteOnlyBeaconBlockBody_SetGraffiti_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bytes.B32))
	})
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetGraffiti_Call) Return() *WriteOnlyBeaconBlockBody_SetGraffiti_Call {
	_c.Call.Return()
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetGraffiti_Call) RunAndReturn(run func(bytes.B32)) *WriteOnlyBeaconBlockBody_SetGraffiti_Call {
	_c.Call.Return(run)
	return _c
}

// SetRandaoReveal provides a mock function with given fields: _a0
func (_m *WriteOnlyBeaconBlockBody) SetRandaoReveal(_a0 bytes.B96) {
	_m.Called(_a0)
}

// WriteOnlyBeaconBlockBody_SetRandaoReveal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetRandaoReveal'
type WriteOnlyBeaconBlockBody_SetRandaoReveal_Call struct {
	*mock.Call
}

// SetRandaoReveal is a helper method to define mock.On call
//   - _a0 bytes.B96
func (_e *WriteOnlyBeaconBlockBody_Expecter) SetRandaoReveal(_a0 interface{}) *WriteOnlyBeaconBlockBody_SetRandaoReveal_Call {
	return &WriteOnlyBeaconBlockBody_SetRandaoReveal_Call{Call: _e.mock.On("SetRandaoReveal", _a0)}
}

func (_c *WriteOnlyBeaconBlockBody_SetRandaoReveal_Call) Run(run func(_a0 bytes.B96)) *WriteOnlyBeaconBlockBody_SetRandaoReveal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bytes.B96))
	})
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetRandaoReveal_Call) Return() *WriteOnlyBeaconBlockBody_SetRandaoReveal_Call {
	_c.Call.Return()
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetRandaoReveal_Call) RunAndReturn(run func(bytes.B96)) *WriteOnlyBeaconBlockBody_SetRandaoReveal_Call {
	_c.Call.Return(run)
	return _c
}

// SetSlashingInfo provides a mock function with given fields: _a0
func (_m *WriteOnlyBeaconBlockBody) SetSlashingInfo(_a0 []*types.SlashingInfo) {
	_m.Called(_a0)
}

// WriteOnlyBeaconBlockBody_SetSlashingInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetSlashingInfo'
type WriteOnlyBeaconBlockBody_SetSlashingInfo_Call struct {
	*mock.Call
}

// SetSlashingInfo is a helper method to define mock.On call
//   - _a0 []*types.SlashingInfo
func (_e *WriteOnlyBeaconBlockBody_Expecter) SetSlashingInfo(_a0 interface{}) *WriteOnlyBeaconBlockBody_SetSlashingInfo_Call {
	return &WriteOnlyBeaconBlockBody_SetSlashingInfo_Call{Call: _e.mock.On("SetSlashingInfo", _a0)}
}

func (_c *WriteOnlyBeaconBlockBody_SetSlashingInfo_Call) Run(run func(_a0 []*types.SlashingInfo)) *WriteOnlyBeaconBlockBody_SetSlashingInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]*types.SlashingInfo))
	})
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetSlashingInfo_Call) Return() *WriteOnlyBeaconBlockBody_SetSlashingInfo_Call {
	_c.Call.Return()
	return _c
}

func (_c *WriteOnlyBeaconBlockBody_SetSlashingInfo_Call) RunAndReturn(run func([]*types.SlashingInfo)) *WriteOnlyBeaconBlockBody_SetSlashingInfo_Call {
	_c.Call.Return(run)
	return _c
}

// NewWriteOnlyBeaconBlockBody creates a new instance of WriteOnlyBeaconBlockBody. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWriteOnlyBeaconBlockBody(t interface {
	mock.TestingT
	Cleanup(func())
}) *WriteOnlyBeaconBlockBody {
	mock := &WriteOnlyBeaconBlockBody{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
