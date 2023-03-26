// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "web_server/domain"

	mock "github.com/stretchr/testify/mock"
)

// FeeRepository is an autogenerated mock type for the FeeRepository type
type FeeRepository struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx, req
func (_m *FeeRepository) GetAll(ctx context.Context, req domain.BaseRequest) (domain.FeeResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.FeeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) (domain.FeeResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) domain.FeeResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.FeeResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.BaseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrency provides a mock function with given fields: ctx, req
func (_m *FeeRepository) GetCurrency(ctx context.Context, req domain.BaseRequest) (domain.Currency, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.Currency
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) (domain.Currency, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) domain.Currency); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.Currency)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.BaseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNews provides a mock function with given fields: ctx, req
func (_m *FeeRepository) GetNews(ctx context.Context, req domain.BaseRequest) (domain.News, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.News
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) (domain.News, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) domain.News); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.News)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.BaseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPair provides a mock function with given fields: ctx, req
func (_m *FeeRepository) GetPair(ctx context.Context, req domain.BaseRequest) (domain.Pair, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.Pair
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) (domain.Pair, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) domain.Pair); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.Pair)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.BaseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFeeRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewFeeRepository creates a new instance of FeeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFeeRepository(t mockConstructorTestingTNewFeeRepository) *FeeRepository {
	mock := &FeeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}