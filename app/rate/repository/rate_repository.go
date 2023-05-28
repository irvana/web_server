package repository

import (
	"context"
	"web_server/domain"

	"github.com/gammazero/nexus/v3/client"
	"github.com/go-redis/redis"
	"github.com/rueian/rueidis"
)

type rateRepository struct {
	redisCli   *redis.Client
	wampCli    *client.Client
	redisTsCli rueidis.Client
}

func NewRateRepository(redisCli *redis.Client, wampCli *client.Client, redisTsCli rueidis.Client) domain.RateRepository {
	return &rateRepository{redisCli, wampCli, redisTsCli}
}

// ConsumeRate implements domain.RateRepository
func (r *rateRepository) ConsumeRate(ctx context.Context) (domain.RateResponse, error) {
	return domain.RateResponse{}, nil
}

// PublishRate implements domain.RateRepository
func (r *rateRepository) PublishRate(ctx context.Context, rate domain.RateResponse) error {
	// r.wampCli.Publish()
	return nil
}
