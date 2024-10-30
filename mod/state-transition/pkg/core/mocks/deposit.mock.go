// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	common "github.com/berachain/beacon-kit/mod/primitives/pkg/common"

	crypto "github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"

	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"

	mock "github.com/stretchr/testify/mock"
)

// Deposit is an autogenerated mock type for the Deposit type
type Deposit[ForkDataT any, WithdrawlCredentialsT interface{ ~[32]byte }] struct {
	mock.Mock
}

type Deposit_Expecter[ForkDataT any, WithdrawlCredentialsT interface{ ~[32]byte }] struct {
	mock *mock.Mock
}

func (_m *Deposit[ForkDataT, WithdrawlCredentialsT]) EXPECT() *Deposit_Expecter[ForkDataT, WithdrawlCredentialsT] {
	return &Deposit_Expecter[ForkDataT, WithdrawlCredentialsT]{mock: &_m.Mock}
}

// GetAmount provides a mock function with given fields:
func (_m *Deposit[ForkDataT, WithdrawlCredentialsT]) GetAmount() math.Gwei {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAmount")
	}

	var r0 math.Gwei
	if rf, ok := ret.Get(0).(func() math.Gwei); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(math.Gwei)
	}

	return r0
}

// Deposit_GetAmount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAmount'
type Deposit_GetAmount_Call[ForkDataT any, WithdrawlCredentialsT interface{ ~[32]byte }] struct {
	*mock.Call
}

// GetAmount is a helper method to define mock.On call
func (_e *Deposit_Expecter[ForkDataT, WithdrawlCredentialsT]) GetAmount() *Deposit_GetAmount_Call[ForkDataT, WithdrawlCredentialsT] {
	return &Deposit_GetAmount_Call[ForkDataT, WithdrawlCredentialsT]{Call: _e.mock.On("GetAmount")}
}

func (_c *Deposit_GetAmount_Call[ForkDataT, WithdrawlCredentialsT]) Run(run func()) *Deposit_GetAmount_Call[ForkDataT, WithdrawlCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Deposit_GetAmount_Call[ForkDataT, WithdrawlCredentialsT]) Return(_a0 math.Gwei) *Deposit_GetAmount_Call[ForkDataT, WithdrawlCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Deposit_GetAmount_Call[ForkDataT, WithdrawlCredentialsT]) RunAndReturn(run func() math.Gwei) *Deposit_GetAmount_Call[ForkDataT, WithdrawlCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// GetPubkey provides a mock function with given fields:
func (_m *Deposit[ForkDataT, WithdrawlCredentialsT]) GetPubkey() crypto.BLSPubkey {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetPubkey")
	}

	var r0 crypto.BLSPubkey
	if rf, ok := ret.Get(0).(func() crypto.BLSPubkey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.BLSPubkey)
		}
	}

	return r0
}

// Deposit_GetPubkey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPubkey'
type Deposit_GetPubkey_Call[ForkDataT any, WithdrawlCredentialsT interface{ ~[32]byte }] struct {
	*mock.Call
}

// GetPubkey is a helper method to define mock.On call
func (_e *Deposit_Expecter[ForkDataT, WithdrawlCredentialsT]) GetPubkey() *Deposit_GetPubkey_Call[ForkDataT, WithdrawlCredentialsT] {
	return &Deposit_GetPubkey_Call[ForkDataT, WithdrawlCredentialsT]{Call: _e.mock.On("GetPubkey")}
}

func (_c *Deposit_GetPubkey_Call[ForkDataT, WithdrawlCredentialsT]) Run(run func()) *Deposit_GetPubkey_Call[ForkDataT, WithdrawlCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Deposit_GetPubkey_Call[ForkDataT, WithdrawlCredentialsT]) Return(_a0 crypto.BLSPubkey) *Deposit_GetPubkey_Call[ForkDataT, WithdrawlCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Deposit_GetPubkey_Call[ForkDataT, WithdrawlCredentialsT]) RunAndReturn(run func() crypto.BLSPubkey) *Deposit_GetPubkey_Call[ForkDataT, WithdrawlCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// GetWithdrawalCredentials provides a mock function with given fields:
func (_m *Deposit[ForkDataT, WithdrawlCredentialsT]) GetWithdrawalCredentials() WithdrawlCredentialsT {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetWithdrawalCredentials")
	}

	var r0 WithdrawlCredentialsT
	if rf, ok := ret.Get(0).(func() WithdrawlCredentialsT); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(WithdrawlCredentialsT)
	}

	return r0
}

// Deposit_GetWithdrawalCredentials_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWithdrawalCredentials'
type Deposit_GetWithdrawalCredentials_Call[ForkDataT any, WithdrawlCredentialsT interface{ ~[32]byte }] struct {
	*mock.Call
}

// GetWithdrawalCredentials is a helper method to define mock.On call
func (_e *Deposit_Expecter[ForkDataT, WithdrawlCredentialsT]) GetWithdrawalCredentials() *Deposit_GetWithdrawalCredentials_Call[ForkDataT, WithdrawlCredentialsT] {
	return &Deposit_GetWithdrawalCredentials_Call[ForkDataT, WithdrawlCredentialsT]{Call: _e.mock.On("GetWithdrawalCredentials")}
}

func (_c *Deposit_GetWithdrawalCredentials_Call[ForkDataT, WithdrawlCredentialsT]) Run(run func()) *Deposit_GetWithdrawalCredentials_Call[ForkDataT, WithdrawlCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Deposit_GetWithdrawalCredentials_Call[ForkDataT, WithdrawlCredentialsT]) Return(_a0 WithdrawlCredentialsT) *Deposit_GetWithdrawalCredentials_Call[ForkDataT, WithdrawlCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Deposit_GetWithdrawalCredentials_Call[ForkDataT, WithdrawlCredentialsT]) RunAndReturn(run func() WithdrawlCredentialsT) *Deposit_GetWithdrawalCredentials_Call[ForkDataT, WithdrawlCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// VerifySignature provides a mock function with given fields: forkData, domainType, signatureVerificationFn
func (_m *Deposit[ForkDataT, WithdrawlCredentialsT]) VerifySignature(forkData ForkDataT, domainType common.DomainType, signatureVerificationFn func(crypto.BLSPubkey, []byte, crypto.BLSSignature) error) error {
	ret := _m.Called(forkData, domainType, signatureVerificationFn)

	if len(ret) == 0 {
		panic("no return value specified for VerifySignature")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(ForkDataT, common.DomainType, func(crypto.BLSPubkey, []byte, crypto.BLSSignature) error) error); ok {
		r0 = rf(forkData, domainType, signatureVerificationFn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Deposit_VerifySignature_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VerifySignature'
type Deposit_VerifySignature_Call[ForkDataT any, WithdrawlCredentialsT interface{ ~[32]byte }] struct {
	*mock.Call
}

// VerifySignature is a helper method to define mock.On call
//   - forkData ForkDataT
//   - domainType common.DomainType
//   - signatureVerificationFn func(crypto.BLSPubkey , []byte , crypto.BLSSignature) error
func (_e *Deposit_Expecter[ForkDataT, WithdrawlCredentialsT]) VerifySignature(forkData interface{}, domainType interface{}, signatureVerificationFn interface{}) *Deposit_VerifySignature_Call[ForkDataT, WithdrawlCredentialsT] {
	return &Deposit_VerifySignature_Call[ForkDataT, WithdrawlCredentialsT]{Call: _e.mock.On("VerifySignature", forkData, domainType, signatureVerificationFn)}
}

func (_c *Deposit_VerifySignature_Call[ForkDataT, WithdrawlCredentialsT]) Run(run func(forkData ForkDataT, domainType common.DomainType, signatureVerificationFn func(crypto.BLSPubkey, []byte, crypto.BLSSignature) error)) *Deposit_VerifySignature_Call[ForkDataT, WithdrawlCredentialsT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ForkDataT), args[1].(common.DomainType), args[2].(func(crypto.BLSPubkey, []byte, crypto.BLSSignature) error))
	})
	return _c
}

func (_c *Deposit_VerifySignature_Call[ForkDataT, WithdrawlCredentialsT]) Return(_a0 error) *Deposit_VerifySignature_Call[ForkDataT, WithdrawlCredentialsT] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Deposit_VerifySignature_Call[ForkDataT, WithdrawlCredentialsT]) RunAndReturn(run func(ForkDataT, common.DomainType, func(crypto.BLSPubkey, []byte, crypto.BLSSignature) error) error) *Deposit_VerifySignature_Call[ForkDataT, WithdrawlCredentialsT] {
	_c.Call.Return(run)
	return _c
}

// NewDeposit creates a new instance of Deposit. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeposit[ForkDataT any, WithdrawlCredentialsT interface{ ~[32]byte }](t interface {
	mock.TestingT
	Cleanup(func())
}) *Deposit[ForkDataT, WithdrawlCredentialsT] {
	mock := &Deposit[ForkDataT, WithdrawlCredentialsT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
