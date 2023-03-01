package domain

import "context"

type Currency struct {
	ID      string `json:"id"`
	Label   string `json:"label"`
	Name    string `json:"name"`
	Decimal string `json:"decimal"`
}

type Pair struct {
	ID         string `json:"id"`
	Label      string `json:"label"`
	BaseCcyID  string `json:"baseCcyId"`
	QuoteCcyID string `json:"quoteCcyId"`
	Decimal    string `json:"decimal"`
	Tradeable  string `json:"tradeable"`
}

type News struct {
	ID        string `json:"id"`
	NewsID    string `json:"newsId"`
	PairID    string `json:"pairId"`
	Stamp     string `json:"stamp"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type FeeResponse struct {
	Currency Currency `json:"ccy"`
	Pair     Pair     `json:"pair"`
	News     News     `json:"news"`
}

type FeeUsecase interface {
	GetAll(ctx context.Context, req BaseRequest) (FeeResponse, error)
	GetCurrency(ctx context.Context, req BaseRequest) (Currency, error)
	GetPair(ctx context.Context, req BaseRequest) (Pair, error)
	GetNews(ctx context.Context, req BaseRequest) (News, error)
}

type FeeRepository interface {
	GetAll(ctx context.Context, req BaseRequest) (FeeResponse, error)
	GetCurrency(ctx context.Context, req BaseRequest) (Currency, error)
	GetPair(ctx context.Context, req BaseRequest) (Pair, error)
	GetNews(ctx context.Context, req BaseRequest) (News, error)
}
