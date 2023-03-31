package usecase

import (
	"context"
	"web_server/domain"
)

type onboardingUsecase struct {
	ob domain.OrderRepository
}

func NewOrderUsecase(ob domain.OrderRepository) domain.OrderUsecase {
	return &onboardingUsecase{ob}
}

// Amend implements domain.OrderUsecase
func (o *onboardingUsecase) Amend(ctx context.Context, req *domain.BaseRequest) (*domain.OrderResponse, error) {
	return o.ob.Amend(ctx, req)
}

// Cancel implements domain.OrderUsecase
func (o *onboardingUsecase) Cancel(ctx context.Context, req *domain.BaseRequest) (*domain.OrderResponse, error) {
	return o.ob.Cancel(ctx, req)
}

// Create implements domain.OrderUsecase
func (o *onboardingUsecase) Create(ctx context.Context, req *domain.BaseRequest) (*domain.OrderResponse, error) {
	return o.ob.Create(ctx, req)
}

// GetDetail implements domain.OrderUsecase
func (o *onboardingUsecase) GetDetail(ctx context.Context, req *domain.BaseRequest) (*domain.OrderResponse, error) {
	return o.ob.GetDetail(ctx, req)
}

// GetStatus implements domain.OrderUsecase
func (o *onboardingUsecase) GetStatus(ctx context.Context, req *domain.BaseRequest) (*domain.OrderResponse, error) {
	return o.ob.GetStatus(ctx, req)
}
