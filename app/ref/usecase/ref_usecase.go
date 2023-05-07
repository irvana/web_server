package usecase

import (
	"context"
	"web_server/domain"
)

type refUsecase struct {
	ref domain.RefRepository
}

func NewRefUsecase(refRepo domain.RefRepository) domain.RefUsecase {
	return &refUsecase{refRepo}
}
func (r *refUsecase) GetAll(ctx context.Context, req *domain.RefRequest) ([]domain.RefResponse, error) {

	return r.ref.GetAll(ctx, req)
}

func (r *refUsecase) GetCurrency(ctx context.Context, req *domain.RefRequest) ([]domain.Currency, error) {
	return r.ref.GetCurrency(ctx, req)
}

func (r *refUsecase) GetPair(ctx context.Context, req *domain.RefRequest) ([]domain.Pair, error) {
	return r.ref.GetPair(ctx, req)
}

func (r *refUsecase) GetNews(ctx context.Context, req *domain.RefRequest) ([]domain.News, error) {
	return r.ref.GetNews(ctx, req)
}

func (r *refUsecase) GetConfig(ctx context.Context, req *domain.RefRequest) (*domain.Config, error) {
	return r.ref.GetConfig(ctx, req)
}
