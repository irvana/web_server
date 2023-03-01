package usecase

import (
	"context"
	"web_server/domain"
)

type rateUsecase struct {
	rateRepo domain.RateRepository
}

func NewRateUsecase(rateRepo domain.RateRepository) domain.RateUsecase {
	return &rateUsecase{rateRepo}
}

// PublishRate implements domain.RateUsecase
func (r *rateUsecase) PublishRate(ctx context.Context) {
	for {
		resp, err := r.rateRepo.ConsumeRate(ctx)
		if err != nil {
			return
		}

		r.rateRepo.PublishRate(ctx, resp)
	}

}
