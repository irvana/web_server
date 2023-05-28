package domain

import "context"

type (
	RateResponse struct {
	}

	HistoricRateRequest struct {
	}
)

type RateUsecase interface {
	ProcessBackgroundRate(ctx context.Context)
	ProcessBackgroundRef(ctx context.Context)
	GetChart(ctx context.Context, req *HistoricRateRequest) ([]RateResponse, error)
}

type RateRepository interface {
	ConsumeRate(ctx context.Context) (RateResponse, error)
	PublishRate(ctx context.Context, rate RateResponse) error
}

type HistoricRates interface {
	GetChart(ctx context.Context, req *HistoricRateRequest) ([]RateResponse, error)
}
