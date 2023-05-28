package usecase

import (
	"context"
	"web_server/domain"

	log "github.com/sirupsen/logrus"
)

type rateUsecase struct {
	rateRepo domain.RateRepository
}

func NewRateUsecase(rateRepo domain.RateRepository) domain.RateUsecase {
	return &rateUsecase{rateRepo}
}

// PublishRate implements domain.RateUsecase
func (r *rateUsecase) ProcessBackgroundRate(ctx context.Context) {
	log.WithField("foo", ctx.Value("foo")).Info("running rate publisher")
	for {
		resp, err := r.rateRepo.ConsumeRate(ctx)
		if err != nil {
			log.WithError(err).Errorln("error consuming rate")
		}

		err = r.rateRepo.PublishRate(ctx, resp)
		if err != nil {
			log.WithError(err).Errorln("error publishing rate")
		}
	}

}

// PublishReference implements domain.RateUsecase
func (r *rateUsecase) ProcessBackgroundRef(ctx context.Context) {
	log.WithField("foo", ctx.Value("foo")).Info("running ref publisher")
	for {
		resp, err := r.rateRepo.ConsumeRate(ctx)
		if err != nil {
			log.WithError(err).Errorln("error consuming ref")

		}

		err = r.rateRepo.PublishRate(ctx, resp)
		if err != nil {
			log.WithError(err).Errorln("error publishing ref")
		}
	}
}

func (r *rateUsecase) GetChart(ctx context.Context, req *domain.HistoricRateRequest) ([]domain.RateResponse, error) {
	panic("not implemented") // TODO: Implement
}
