package usecase

import (
	"context"
	"testing"
	"web_server/domain"
	"web_server/domain/mocks"

	"github.com/stretchr/testify/assert"
)

func Test_refUsecase_GetAll(t *testing.T) {
	// Set up sample request
	req := &domain.RefRequest{}

	// Create onboarding usecase with mocked repository
	stmRepo := new(mocks.RefRepository)
	u := NewRefUsecase(stmRepo)

	// Mock ATM Get Info method to return expected result
	expectedResult := []domain.RefResponse{}
	stmRepo.On("GetAll", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.GetAll(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_refUsecase_GetCurrency(t *testing.T) {
	// Set up sample request
	req := &domain.RefRequest{}

	// Create onboarding usecase with mocked repository
	stmRepo := new(mocks.RefRepository)
	u := NewRefUsecase(stmRepo)

	// Mock ATM Get Info method to return expected result
	expectedResult := []domain.Currency{}
	stmRepo.On("GetCurrency", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.GetCurrency(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_refUsecase_GetPair(t *testing.T) {
	// Set up sample request
	req := &domain.RefRequest{}

	// Create onboarding usecase with mocked repository
	stmRepo := new(mocks.RefRepository)
	u := NewRefUsecase(stmRepo)

	// Mock ATM Get Info method to return expected result
	expectedResult := []domain.Pair{}
	stmRepo.On("GetPair", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.GetPair(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_refUsecase_GetNews(t *testing.T) {
	// Set up sample request
	req := &domain.RefRequest{}

	// Create onboarding usecase with mocked repository
	stmRepo := new(mocks.RefRepository)
	u := NewRefUsecase(stmRepo)

	// Mock ATM Get Info method to return expected result
	expectedResult := []domain.News{}
	stmRepo.On("GetNews", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.GetNews(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func Test_refUsecase_GetConfig(t *testing.T) {
	// Set up sample request
	req := &domain.RefRequest{}

	// Create onboarding usecase with mocked repository
	stmRepo := new(mocks.RefRepository)
	u := NewRefUsecase(stmRepo)

	// Mock ATM Get Info method to return expected result
	expectedResult := &domain.Config{}
	stmRepo.On("GetConfig", context.Background(), req).Return(expectedResult, nil)

	// Call function and assert result matches expected value
	result, err := u.GetConfig(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}
