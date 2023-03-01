package domain

import "context"

type RateResponse struct {
}

type RateUsecase interface {
	ProcessBackgroundRate(ctx context.Context)
	ProcessBackgroundRef(ctx context.Context)
}

type RateRepository interface {
	ConsumeRate(ctx context.Context) (RateResponse, error)
	PublishRate(ctx context.Context, rate RateResponse) error
}
