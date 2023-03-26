// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "web_server/domain"

	mock "github.com/stretchr/testify/mock"
)

// AccountUsecase is an autogenerated mock type for the AccountUsecase type
type AccountUsecase struct {
	mock.Mock
}

// GetList provides a mock function with given fields: ctx, req
func (_m *AccountUsecase) GetList(ctx context.Context, req domain.BaseRequest) (domain.AccountResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.AccountResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) (domain.AccountResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) domain.AccountResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.AccountResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.BaseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Open provides a mock function with given fields: ctx, req
func (_m *AccountUsecase) Open(ctx context.Context, req domain.BaseRequest) (domain.AccountResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.AccountResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) (domain.AccountResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) domain.AccountResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.AccountResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.BaseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAccountUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAccountUsecase creates a new instance of AccountUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccountUsecase(t mockConstructorTestingTNewAccountUsecase) *AccountUsecase {
	mock := &AccountUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
