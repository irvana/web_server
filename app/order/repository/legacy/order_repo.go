package legacy

import (
	"context"
	"web_server/domain"
)

type onboardingLegacyRepo struct {
}

func NewOnboardingRepository() domain.OrderRepository {
	return &onboardingLegacyRepo{}
}

// Amend implements domain.OrderRepository
func (*onboardingLegacyRepo) Amend(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	panic("unimplemented")
}

// Cancel implements domain.OrderRepository
func (*onboardingLegacyRepo) Cancel(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	panic("unimplemented")
}

// Create implements domain.OrderRepository
func (*onboardingLegacyRepo) Create(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	panic("unimplemented")
}

// GetDetail implements domain.OrderRepository
func (*onboardingLegacyRepo) GetDetail(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	panic("unimplemented")
}

// GetStatus implements domain.OrderRepository
func (*onboardingLegacyRepo) GetStatus(ctx context.Context, req domain.BaseRequest) (domain.OrderResponse, error) {
	panic("unimplemented")
}
