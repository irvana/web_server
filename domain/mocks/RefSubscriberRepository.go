// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	domain "web_server/domain"

	redis "github.com/go-redis/redis"
	mock "github.com/stretchr/testify/mock"
)

// RefSubscriberRepository is an autogenerated mock type for the RefSubscriberRepository type
type RefSubscriberRepository struct {
	mock.Mock
}

// Consume provides a mock function with given fields:
func (_m *RefSubscriberRepository) Consume() <-chan *redis.Message {
	ret := _m.Called()

	var r0 <-chan *redis.Message
	if rf, ok := ret.Get(0).(func() <-chan *redis.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *redis.Message)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *RefSubscriberRepository) GetAll() ([]domain.RefResponse, error) {
	ret := _m.Called()

	var r0 []domain.RefResponse
	if rf, ok := ret.Get(0).(func() []domain.RefResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.RefResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRefSubscriberRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRefSubscriberRepository creates a new instance of RefSubscriberRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRefSubscriberRepository(t mockConstructorTestingTNewRefSubscriberRepository) *RefSubscriberRepository {
	mock := &RefSubscriberRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
