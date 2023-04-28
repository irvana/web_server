// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "web_server/domain"

	mock "github.com/stretchr/testify/mock"
)

// RefRepository is an autogenerated mock type for the RefRepository type
type RefRepository struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: ctx, req
func (_m *RefRepository) GetAll(ctx context.Context, req domain.RefRequest) (domain.RefResponse, error) {
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
func (_m *RefRepository) GetConfig(ctx context.Context, req domain.RefRequest) (domain.Config, error) {
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
func (_m *RefRepository) GetCurrency(ctx context.Context, req domain.RefRequest) (domain.Currency, error) {
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
func (_m *RefRepository) GetNews(ctx context.Context, req domain.RefRequest) (domain.News, error) {
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
func (_m *RefRepository) GetPair(ctx context.Context, req domain.RefRequest) (domain.Pair, error) {
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

type mockConstructorTestingTNewRefRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRefRepository creates a new instance of RefRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRefRepository(t mockConstructorTestingTNewRefRepository) *RefRepository {
	mock := &RefRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
