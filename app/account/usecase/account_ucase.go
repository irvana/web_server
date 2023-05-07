package usecase

import (
	"context"
	"web_server/domain"
)

type accountUsecase struct {
	ar domain.AccountRepository
}

func NewAccountUsecase(ar domain.AccountRepository) domain.AccountUsecase {
	return &accountUsecase{ar}
}
func (ac *accountUsecase) GetList(ctx context.Context, req *domain.BaseRequest) ([]domain.AccountResponse, error) {
	return ac.ar.GetList(ctx, req)
}

func (ac *accountUsecase) Open(ctx context.Context, req *domain.BaseRequest) (*domain.AccountOpenResponse, error) {
	return ac.ar.Open(ctx, req)
}
