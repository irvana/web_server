package handler

import (
	"context"
	"web_server/domain"
)

func RunRateHandlers(rateUc domain.RateUsecase) {
	go rateUc.ProcessBackgroundRate(context.WithValue(context.Background(), "rate", "ref"))
	go rateUc.ProcessBackgroundRef(context.WithValue(context.Background(), "ref", "ref"))
}
