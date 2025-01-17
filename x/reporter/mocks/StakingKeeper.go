// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"

	math "cosmossdk.io/math"
	db "github.com/cosmos/cosmos-db"

	mock "github.com/stretchr/testify/mock"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// StakingKeeper is an autogenerated mock type for the StakingKeeper type
type StakingKeeper struct {
	mock.Mock
}

// Delegate provides a mock function with given fields: ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount
func (_m *StakingKeeper) Delegate(ctx context.Context, delAddr types.AccAddress, bondAmt math.Int, tokenSrc stakingtypes.BondStatus, validator stakingtypes.Validator, subtractAccount bool) (math.LegacyDec, error) {
	ret := _m.Called(ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount)

	var r0 math.LegacyDec
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, math.Int, stakingtypes.BondStatus, stakingtypes.Validator, bool) (math.LegacyDec, error)); ok {
		return rf(ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, math.Int, stakingtypes.BondStatus, stakingtypes.Validator, bool) math.LegacyDec); ok {
		r0 = rf(ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount)
	} else {
		r0 = ret.Get(0).(math.LegacyDec)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.AccAddress, math.Int, stakingtypes.BondStatus, stakingtypes.Validator, bool) error); ok {
		r1 = rf(ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delegation provides a mock function with given fields: _a0, _a1, _a2
func (_m *StakingKeeper) Delegation(_a0 context.Context, _a1 types.AccAddress, _a2 types.ValAddress) (stakingtypes.DelegationI, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 stakingtypes.DelegationI
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, types.ValAddress) (stakingtypes.DelegationI, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, types.ValAddress) stakingtypes.DelegationI); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(stakingtypes.DelegationI)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.AccAddress, types.ValAddress) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDelegation provides a mock function with given fields: ctx, delAddr, valAddr
func (_m *StakingKeeper) GetDelegation(ctx context.Context, delAddr types.AccAddress, valAddr types.ValAddress) (stakingtypes.Delegation, error) {
	ret := _m.Called(ctx, delAddr, valAddr)

	var r0 stakingtypes.Delegation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, types.ValAddress) (stakingtypes.Delegation, error)); ok {
		return rf(ctx, delAddr, valAddr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, types.ValAddress) stakingtypes.Delegation); ok {
		r0 = rf(ctx, delAddr, valAddr)
	} else {
		r0 = ret.Get(0).(stakingtypes.Delegation)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.AccAddress, types.ValAddress) error); ok {
		r1 = rf(ctx, delAddr, valAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRedelegationsFromSrcValidator provides a mock function with given fields: ctx, valAddr
func (_m *StakingKeeper) GetRedelegationsFromSrcValidator(ctx context.Context, valAddr types.ValAddress) ([]stakingtypes.Redelegation, error) {
	ret := _m.Called(ctx, valAddr)

	var r0 []stakingtypes.Redelegation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.ValAddress) ([]stakingtypes.Redelegation, error)); ok {
		return rf(ctx, valAddr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.ValAddress) []stakingtypes.Redelegation); ok {
		r0 = rf(ctx, valAddr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]stakingtypes.Redelegation)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.ValAddress) error); ok {
		r1 = rf(ctx, valAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUnbondingDelegation provides a mock function with given fields: ctx, delAddr, valAddr
func (_m *StakingKeeper) GetUnbondingDelegation(ctx context.Context, delAddr types.AccAddress, valAddr types.ValAddress) (stakingtypes.UnbondingDelegation, error) {
	ret := _m.Called(ctx, delAddr, valAddr)

	var r0 stakingtypes.UnbondingDelegation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, types.ValAddress) (stakingtypes.UnbondingDelegation, error)); ok {
		return rf(ctx, delAddr, valAddr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, types.ValAddress) stakingtypes.UnbondingDelegation); ok {
		r0 = rf(ctx, delAddr, valAddr)
	} else {
		r0 = ret.Get(0).(stakingtypes.UnbondingDelegation)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.AccAddress, types.ValAddress) error); ok {
		r1 = rf(ctx, delAddr, valAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetValidator provides a mock function with given fields: ctx, addr
func (_m *StakingKeeper) GetValidator(ctx context.Context, addr types.ValAddress) (stakingtypes.Validator, error) {
	ret := _m.Called(ctx, addr)

	var r0 stakingtypes.Validator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.ValAddress) (stakingtypes.Validator, error)); ok {
		return rf(ctx, addr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.ValAddress) stakingtypes.Validator); ok {
		r0 = rf(ctx, addr)
	} else {
		r0 = ret.Get(0).(stakingtypes.Validator)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.ValAddress) error); ok {
		r1 = rf(ctx, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetValidators provides a mock function with given fields: ctx, maxRetrieve
func (_m *StakingKeeper) GetValidators(ctx context.Context, maxRetrieve uint32) ([]stakingtypes.Validator, error) {
	ret := _m.Called(ctx, maxRetrieve)

	var r0 []stakingtypes.Validator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) ([]stakingtypes.Validator, error)); ok {
		return rf(ctx, maxRetrieve)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []stakingtypes.Validator); ok {
		r0 = rf(ctx, maxRetrieve)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]stakingtypes.Validator)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, maxRetrieve)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveUnbondingDelegation provides a mock function with given fields: ctx, ubd
func (_m *StakingKeeper) RemoveUnbondingDelegation(ctx context.Context, ubd stakingtypes.UnbondingDelegation) error {
	ret := _m.Called(ctx, ubd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, stakingtypes.UnbondingDelegation) error); ok {
		r0 = rf(ctx, ubd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetUnbondingDelegation provides a mock function with given fields: ctx, ubd
func (_m *StakingKeeper) SetUnbondingDelegation(ctx context.Context, ubd stakingtypes.UnbondingDelegation) error {
	ret := _m.Called(ctx, ubd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, stakingtypes.UnbondingDelegation) error); ok {
		r0 = rf(ctx, ubd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TotalBondedTokens provides a mock function with given fields: _a0
func (_m *StakingKeeper) TotalBondedTokens(_a0 context.Context) (math.Int, error) {
	ret := _m.Called(_a0)

	var r0 math.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (math.Int, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) math.Int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unbond provides a mock function with given fields: ctx, delAddr, valAddr, shares
func (_m *StakingKeeper) Unbond(ctx context.Context, delAddr types.AccAddress, valAddr types.ValAddress, shares math.LegacyDec) (math.Int, error) {
	ret := _m.Called(ctx, delAddr, valAddr, shares)

	var r0 math.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, types.ValAddress, math.LegacyDec) (math.Int, error)); ok {
		return rf(ctx, delAddr, valAddr, shares)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.AccAddress, types.ValAddress, math.LegacyDec) math.Int); ok {
		r0 = rf(ctx, delAddr, valAddr, shares)
	} else {
		r0 = ret.Get(0).(math.Int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.AccAddress, types.ValAddress, math.LegacyDec) error); ok {
		r1 = rf(ctx, delAddr, valAddr, shares)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatorsPowerStoreIterator provides a mock function with given fields: ctx
func (_m *StakingKeeper) ValidatorsPowerStoreIterator(ctx context.Context) (db.Iterator, error) {
	ret := _m.Called(ctx)

	var r0 db.Iterator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (db.Iterator, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) db.Iterator); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Iterator)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStakingKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewStakingKeeper creates a new instance of StakingKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStakingKeeper(t mockConstructorTestingTNewStakingKeeper) *StakingKeeper {
	mock := &StakingKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
