package usecase

import (
	"context"
	"web_server/domain"
)

type onboardingUsecase struct {
	ob      domain.OrderRepository
	stmRepo domain.StatementRepository
	trRepo  domain.TransactionRepository
}

func NewOrderUsecase(ob domain.OrderRepository, stmRepo domain.StatementRepository, trRepo domain.TransactionRepository) domain.OrderUsecase {
	return &onboardingUsecase{ob, stmRepo, trRepo}
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
func (o *onboardingUsecase) GetStatus(ctx context.Context, req *domain.BaseRequest) ([]domain.OrderResponse, error) {
	return o.ob.GetStatus(ctx, req)
}

func (o *onboardingUsecase) GetStatementList(ctx context.Context, req *domain.BaseRequest) ([]domain.StatementResponse, error) {
	return o.stmRepo.List(ctx, req)
}

func (o *onboardingUsecase) GetTransactionDetail(ctx context.Context, req *domain.BaseRequest) (*domain.TransactionResponse, error) {
	return o.trRepo.GetDetail(ctx, req)
}

func (o *onboardingUsecase) Deal(ctx context.Context, req *domain.BaseRequest) (*domain.TransactionResponse, error) {
	return o.trRepo.Deal(ctx, req)
}
