// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "web_server/domain"

	mock "github.com/stretchr/testify/mock"
)

// RefUsecase is an autogenerated mock type for the RefUsecase type
type RefUsecase struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx, req
func (_m *RefUsecase) GetAll(ctx context.Context, req domain.RefRequest) (domain.RefResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.RefResponse
	if rf, ok := ret.Get(0).(func(context.Context, domain.RefRequest) domain.RefResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.RefResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.RefRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfig provides a mock function with given fields: ctx, req
func (_m *RefUsecase) GetConfig(ctx context.Context, req domain.RefRequest) (domain.Config, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.Config
	if rf, ok := ret.Get(0).(func(context.Context, domain.RefRequest) domain.Config); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.Config)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.RefRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrency provides a mock function with given fields: ctx, req
func (_m *RefUsecase) GetCurrency(ctx context.Context, req domain.RefRequest) (domain.Currency, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.Currency
	if rf, ok := ret.Get(0).(func(context.Context, domain.RefRequest) domain.Currency); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.Currency)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.RefRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNews provides a mock function with given fields: ctx, req
func (_m *RefUsecase) GetNews(ctx context.Context, req domain.RefRequest) (domain.News, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.News
	if rf, ok := ret.Get(0).(func(context.Context, domain.RefRequest) domain.News); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.News)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.RefRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPair provides a mock function with given fields: ctx, req
func (_m *RefUsecase) GetPair(ctx context.Context, req domain.RefRequest) (domain.Pair, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.Pair
	if rf, ok := ret.Get(0).(func(context.Context, domain.RefRequest) domain.Pair); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.Pair)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.RefRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRefUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewRefUsecase creates a new instance of RefUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRefUsecase(t mockConstructorTestingTNewRefUsecase) *RefUsecase {
	mock := &RefUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
