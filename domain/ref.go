package domain

import "context"

type Currency struct {
	ID      string `json:"id,omitempty"`
	Label   string `json:"label,omitempty"`
	Name    string `json:"name,omitempty"`
	Decimal string `json:"decimal,omitempty"`
}

type Pair struct {
	ID         string `json:"id,omitempty"`
	Label      string `json:"label,omitempty"`
	BaseCcyID  string `json:"baseCcyId,omitempty"`
	QuoteCcyID string `json:"quoteCcyId,omitempty"`
	Decimal    string `json:"decimal,omitempty"`
	Tradeable  string `json:"tradeable,omitempty"`
}

type News struct {
	ID        string `json:"id,omitempty"`
	NewsID    string `json:"newsId,omitempty"`
	PairID    string `json:"pairId,omitempty"`
	Stamp     string `json:"stamp,omitempty"`
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
}

type RefRequest struct {
	Type string `json:"type"`
}

type RefResponse struct {
	Currency Currency `json:"ccy,omitempty"`
	Pair     Pair     `json:"pair,omitempty"`
	News     News     `json:"news,omitempty"`
}
type Config struct {
	ID         string `json:"id,omitempty"`
	AutoLocked string `json:"autoLocked,omitempty"`
}

type RefUsecase interface {
	GetAll(ctx context.Context, req RefRequest) (RefResponse, error)
	GetCurrency(ctx context.Context, req RefRequest) (Currency, error)
	GetPair(ctx context.Context, req RefRequest) (Pair, error)
	GetNews(ctx context.Context, req RefRequest) (News, error)
	GetConfig(ctx context.Context, req RefRequest) (Config, error)
}

type RefRepository interface {
	GetAll(ctx context.Context, req RefRequest) (RefResponse, error)
	GetCurrency(ctx context.Context, req RefRequest) (Currency, error)
	GetPair(ctx context.Context, req RefRequest) (Pair, error)
	GetNews(ctx context.Context, req RefRequest) (News, error)
	GetConfig(ctx context.Context, req RefRequest) (Config, error)
}
