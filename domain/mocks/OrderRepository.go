// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "web_server/domain"

	mock "github.com/stretchr/testify/mock"
)

// OrderRepository is an autogenerated mock type for the OrderRepository type
type OrderRepository struct {
	mock.Mock
}

// Amend provides a mock function with given fields: ctx, req
func (_m *OrderRepository) Amend(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.OrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) (domain.OrderResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) domain.OrderResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.OrderResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.BaseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Cancel provides a mock function with given fields: ctx, req
func (_m *OrderRepository) Cancel(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.OrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) (domain.OrderResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) domain.OrderResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.OrderResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.BaseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, req
func (_m *OrderRepository) Create(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.OrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) (domain.OrderResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) domain.OrderResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.OrderResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.BaseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetail provides a mock function with given fields: ctx, req
func (_m *OrderRepository) GetDetail(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.OrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) (domain.OrderResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) domain.OrderResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.OrderResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.BaseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStatus provides a mock function with given fields: ctx, req
func (_m *OrderRepository) GetStatus(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.OrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) (domain.OrderResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) domain.OrderResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(domain.OrderResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.BaseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewOrderRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderRepository creates a new instance of OrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderRepository(t mockConstructorTestingTNewOrderRepository) *OrderRepository {
	mock := &OrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
