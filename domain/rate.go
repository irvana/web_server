package domain

import (
	"context"
)

type (
	RateResponse struct {
		Stamp   string   `json:"stamp"`
		Results []Result `json:"results"`
	}

	Result struct {
		PairID int64  `json:"pairId"`
		Bid    string `json:"bid"`
		Ask    string `json:"ask"`
	}

	HistoricRateResponse struct {
		PairID string     `json:"pairId"`
		Period string     `json:"period"`
		Ts     []RatePair `json:"ts"`
	}

	RatePair struct {
		Stamp string `json:"stamp"`
		Rate  string `json:"rate"`
	}

	HistoricRateRequest struct {
		PairId string `json:"pairId,omitempty"`
		Period string `json:"period,omitempty"`
	}

	HistoricParam struct {
		Key   string
		Start int64
		End   int64
	}
)

type RateUsecase interface {
	ProcessBackgroundRate(ctx context.Context)
	GetChart(ctx context.Context, req *HistoricRateRequest) (*HistoricRateResponse, error)
}

type RateRepository interface {
	ConsumeRate(ctx context.Context, chn chan<- RateResponse)
	PublishRate(ctx context.Context, rate RateResponse) error
}

type HistoricRates interface {
	GetChart(ctx context.Context, req *HistoricParam) (*HistoricRateResponse, error)
}
