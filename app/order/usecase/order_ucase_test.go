package usecase

import (
	"context"
	"testing"
	"web_server/domain"
	"web_server/domain/mocks"

	"github.com/stretchr/testify/assert"
)

func TestOrderUsecase_Amend(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OrderRepository)
	u := NewOrderUsecase(obRepo)

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.OrderResponse{}
	obRepo.On("Amend", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.Amend(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestOrderUsecase_Cancel(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OrderRepository)
	u := NewOrderUsecase(obRepo)

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.OrderResponse{}
	obRepo.On("Cancel", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.Cancel(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}
func TestOrderUsecase_Create(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OrderRepository)
	u := NewOrderUsecase(obRepo)

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.OrderResponse{}
	obRepo.On("Create", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.Create(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}
func TestOrderUsecase_GetDetail(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OrderRepository)
	u := NewOrderUsecase(obRepo)

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.OrderResponse{}
	obRepo.On("GetDetail", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.GetDetail(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}
func TestOrderUsecase_GetStatus(t *testing.T) {
	// Set up sample request
	req := &domain.BaseRequest{}

	// Create onboarding usecase with mocked repository
	obRepo := new(mocks.OrderRepository)
	u := NewOrderUsecase(obRepo)

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.OrderResponse{}
	obRepo.On("GetStatus", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.GetStatus(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}
