// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "web_server/domain"

	mock "github.com/stretchr/testify/mock"
)

// HistoricRates is an autogenerated mock type for the HistoricRates type
type HistoricRates struct {
	mock.Mock
}

// GetChart provides a mock function with given fields: ctx, req
func (_m *HistoricRates) GetChart(ctx context.Context, req *domain.HistoricRateRequest) ([]domain.RateResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 []domain.RateResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.HistoricRateRequest) []domain.RateResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.RateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.HistoricRateRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHistoricRates interface {
	mock.TestingT
	Cleanup(func())
}

// NewHistoricRates creates a new instance of HistoricRates. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHistoricRates(t mockConstructorTestingTNewHistoricRates) *HistoricRates {
	mock := &HistoricRates{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
