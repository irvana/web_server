package usecase

import (
	"context"
	"testing"
	"web_server/domain"
	"web_server/domain/mocks"

	"github.com/stretchr/testify/assert"
)

func Test_accountUsecase_GetList(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	acRepo := new(mocks.AccountRepository)
	u := NewAccountUsecase(acRepo)

	// Mock ATM Get Info method to return expected result
	expectedResult := []domain.AccountResponse{}
	acRepo.On("GetList", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.GetList(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_accountUsecase_Open(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	acRepo := new(mocks.AccountRepository)
	u := NewAccountUsecase(acRepo)

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.AccountOpenResponse{}
	acRepo.On("Open", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.Open(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}
