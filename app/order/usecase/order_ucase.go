package usecase

import (
	"context"
	"web_server/domain"
)

type onboardingUsecase struct {
}

// Amend implements domain.OrderUsecase
func (*onboardingUsecase) Amend(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	panic("unimplemented")
}

// Cancel implements domain.OrderUsecase
func (*onboardingUsecase) Cancel(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	panic("unimplemented")
}

// Create implements domain.OrderUsecase
func (*onboardingUsecase) Create(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	panic("unimplemented")
}

// GetDetail implements domain.OrderUsecase
func (*onboardingUsecase) GetDetail(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	panic("unimplemented")
}

// GetStatus implements domain.OrderUsecase
func (*onboardingUsecase) GetStatus(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	panic("unimplemented")
}

func NewOrderUsecase(obRepo domain.OrderRepository) domain.OrderUsecase {
	return &onboardingUsecase{}
}
