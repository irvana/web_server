package usecase

import (
	"context"
	"testing"

	"web_server/domain"
	"web_server/domain/mocks"
	"web_server/pkg/authentication"

	"github.com/alicebob/miniredis"
	goredis "github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func TestOnboardingUsecase_AtmsGetInfo(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("AtmGetInfo", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.AtmGetInfo(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestOnboardingUsecase_AtmOTP(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("AtmOTP", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.AtmOTP(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestOnboardingUsecase_AtmPIN(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("AtmPIN", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.AtmPIN(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestOnboardingUsecase_AtmRegister(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("AtmRegister", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.AtmRegister(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestOnboardingUsecase_FreshOTP(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("FreshOTP", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.FreshOTP(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestOnboardingUsecase_FreshPassword(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("FreshPassword", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.FreshPassword(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestOnboardingUsecase_FreshVerifyPhone(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("FreshVerifyPhone", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.FreshVerifyPhone(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestOnboardingUsecase_LoginVerifyPassword(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	mr, err := miniredis.Run()
	if err != nil {
		t.Error(err)
		return
	}
	defer mr.Close()
	cli := goredis.NewClient(&goredis.Options{
		Addr: mr.Addr(),
	})

	auth, err := authentication.InitAuthModule(authentication.Config{Secret: "secret"}, cli)
	if err != nil {
		t.Error(err)
	}
	u := NewOnboardingUsecase(obRepo, auth)

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.VerifyPassword{}
	obRepo.On("LoginVerifyPassword", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.LoginVerifyPassword(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestOnboardingUsecase_NoAtmEmail(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("NoAtmEmail", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.NoAtmEmail(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestOnboardingUsecase_NoAtmOTP(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("NoAtmOTP", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.NoAtmOTP(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestOnboardingUsecase_NoAtmRegister(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("NoAtmRegister", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.NoAtmRegister(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestOnboardingUsecase_NoAtmVerifyUser(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("NoAtmVerifyUser", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.NoAtmVerifyUser(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestOnboardingUsecase_ResetOTP(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("ResetOTP", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.ResetOTP(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestOnboardingUsecase_ResetPassword(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("ResetPassword", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.ResetPassword(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestOnboardingUsecase_ResetVerifyPhone(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("ResetVerifyPhone", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.ResetVerifyPhone(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestOnboardingUsecase_SimobiOTP(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("SimobiOTP", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.SimobiOTP(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestOnboardingUsecase_SimobiRegister(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("SimobiRegister", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.SimobiRegister(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestOnboardingUsecase_SimobiVerifyPhone(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("SimobiVerifyPhone", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.SimobiVerifyPhone(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestOnboardingUsecase_SimobiVerifyUser(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OnboardingRepository)
	u := NewOnboardingUsecase(obRepo, &authentication.Authentication{})

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.BaseResponse{}
	obRepo.On("SimobiVerifyUser", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.SimobiVerifyUser(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}
