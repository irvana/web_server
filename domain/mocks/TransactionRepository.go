// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "web_server/domain"

	mock "github.com/stretchr/testify/mock"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

// Deal provides a mock function with given fields: ctx, req
func (_m *TransactionRepository) Deal(ctx context.Context, req *domain.BaseRequest) (*domain.TransactionResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.TransactionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.TransactionResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.TransactionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.BaseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetail provides a mock function with given fields: ctx, req
func (_m *TransactionRepository) GetDetail(ctx context.Context, req *domain.BaseRequest) (*domain.TransactionResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.TransactionResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.TransactionResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.TransactionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.BaseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransactionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionRepository(t mockConstructorTestingTNewTransactionRepository) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
