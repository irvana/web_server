// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "web_server/domain"

	mock "github.com/stretchr/testify/mock"
)

// StatementRepository is an autogenerated mock type for the StatementRepository type
type StatementRepository struct {
	mock.Mock
}

// List provides a mock function with given fields: ctx, req
func (_m *StatementRepository) List(ctx context.Context, req domain.BaseRequest) (domain.StatementResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 domain.StatementResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) (domain.StatementResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.BaseRequest) domain.StatementResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.StatementResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.BaseRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStatementRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewStatementRepository creates a new instance of StatementRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStatementRepository(t mockConstructorTestingTNewStatementRepository) *StatementRepository {
	mock := &StatementRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
