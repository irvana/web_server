// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "web_server/domain"

	mock "github.com/stretchr/testify/mock"
)

// OnboardingRepository is an autogenerated mock type for the OnboardingRepository type
type OnboardingRepository struct {
	mock.Mock
}

// AtmGetInfo provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) AtmGetInfo(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// AtmOTP provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) AtmOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// AtmPIN provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) AtmPIN(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// AtmRegister provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) AtmRegister(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// FreshOTP provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) FreshOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// FreshPassword provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) FreshPassword(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// FreshVerifyPhone provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) FreshVerifyPhone(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// LoginVerifyPassword provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) LoginVerifyPassword(ctx context.Context, req *domain.BaseRequest) (*domain.VerifyPassword, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.VerifyPassword
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.VerifyPassword); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.VerifyPassword)
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

// NoAtmEmail provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) NoAtmEmail(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// NoAtmOTP provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) NoAtmOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// NoAtmRegister provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) NoAtmRegister(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// NoAtmVerifyUser provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) NoAtmVerifyUser(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// ResetOTP provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) ResetOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// ResetPassword provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) ResetPassword(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// ResetVerifyPhone provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) ResetVerifyPhone(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// SimobiOTP provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) SimobiOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// SimobiRegister provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) SimobiRegister(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// SimobiVerifyPhone provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) SimobiVerifyPhone(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

// SimobiVerifyUser provides a mock function with given fields: ctx, req
func (_m *OnboardingRepository) SimobiVerifyUser(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BaseRequest) *domain.BaseResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BaseResponse)
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

type mockConstructorTestingTNewOnboardingRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewOnboardingRepository creates a new instance of OnboardingRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOnboardingRepository(t mockConstructorTestingTNewOnboardingRepository) *OnboardingRepository {
	mock := &OnboardingRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
