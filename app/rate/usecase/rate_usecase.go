package usecase

import (
	"context"
	"web_server/domain"

	"github.com/sirupsen/logrus"
)

type rateUsecase struct {
	rateRepo domain.RateRepository
}

func NewRateUsecase(rateRepo domain.RateRepository) domain.RateUsecase {
	return &rateUsecase{rateRepo}
}

// PublishRate implements domain.RateUsecase
func (r *rateUsecase) ProcessBackgroundRate(ctx context.Context) {
	logrus.WithField("foo", ctx.Value("foo")).Info("running rate publisher")
	for {
		resp, err := r.rateRepo.ConsumeRate(ctx)
		if err != nil {
			return
		}

		r.rateRepo.PublishRate(ctx, resp)
	}

}

// PublishReference implements domain.RateUsecase
func (*rateUsecase) ProcessBackgroundRef(ctx context.Context) {
	panic("unimplemented")
}
