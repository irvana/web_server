package domain

import (
	"context"

	"github.com/go-redis/redis"
)

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
	Type    string       `json:"type"`
	Results []RefDetails `json:"results"`
}

type RefDetails struct {
	ID         string `json:"id"`
	Label      string `json:"label,omitempty"`
	Name       string `json:"name,omitempty"`
	Decimal    string `json:"decimal,omitempty"`
	BaseCcyID  string `json:"baseCcyId,omitempty"`
	QuoteCcyID string `json:"quoteCcyId,omitempty"`
	Tradeable  string `json:"tradeable,omitempty"`
	NewsID     string `json:"newsId,omitempty"`
	PairID     string `json:"pairId,omitempty"`
	Stamp      string `json:"stamp,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	StartDate  string `json:"startDate,omitempty"`
	EndDate    string `json:"endDate,omitempty"`
	AutoLocked string `json:"autoLocked,omitempty"`
	Channel    string `json:"channel,omitempty"`
}
type Config struct {
	ID         string `json:"id,omitempty"`
	AutoLocked string `json:"autoLocked,omitempty"`
}

type RefUsecase interface {
	GetAll(ctx context.Context, req *RefRequest) ([]RefResponse, error)
	GetCurrency(ctx context.Context, req *RefRequest) ([]Currency, error)
	GetPair(ctx context.Context, req *RefRequest) ([]Pair, error)
	GetNews(ctx context.Context, req *RefRequest) ([]News, error)
	GetConfig(ctx context.Context, req *RefRequest) ([]Config, error)
	UpdateConfig()
}

type RefRepository interface {
	GetAll(ctx context.Context, req *RefRequest) ([]RefResponse, error)
	GetCurrency(ctx context.Context, req *RefRequest) ([]Currency, error)
	GetPair(ctx context.Context, req *RefRequest) ([]Pair, error)
	GetNews(ctx context.Context, req *RefRequest) ([]News, error)
	GetConfig(ctx context.Context, req *RefRequest) (*Config, error)
}

type RefSubscriberRepository interface {
	Consume() <-chan *redis.Message
	GetAll() ([]RefResponse, error)
}
