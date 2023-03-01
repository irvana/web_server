package repository

import (
	"context"
	"web_server/domain"

	"github.com/gammazero/nexus/v3/client"
	"github.com/go-redis/redis"
)

type rateRepository struct {
	redisCli *redis.Client
	wampCli  *client.Client
}

func NewRateRepository(redisCli *redis.Client, wampCli *client.Client) domain.RateRepository {
	return &rateRepository{redisCli, wampCli}
}

// ConsumeRate implements domain.RateRepository
func (*rateRepository) ConsumeRate(ctx context.Context, rate domain.RateResponse) error {
	panic("unimplemented")
}

// PublishRate implements domain.RateRepository
func (*rateRepository) PublishRate(ctx context.Context) (domain.RateResponse, error) {
	panic("unimplemented")
}
