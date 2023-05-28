package redis_repo

import (
	"encoding/json"
	"web_server/domain"
	rediscfg "web_server/pkg/redis"

	"github.com/go-redis/redis"
	logger "github.com/sirupsen/logrus"
)

type refRedisRepo struct {
	cli *redis.Client
	cfg rediscfg.Config
}

func NewRefRedisRepository(cli *redis.Client, cfg rediscfg.Config) domain.RefSubscriberRepository {
	return &refRedisRepo{cli, cfg}
}

func (r *refRedisRepo) Consume(channel string) <-chan *redis.Message {
	return r.cli.Subscribe(channel).Channel()
}

func (r *refRedisRepo) GetAll() ([]domain.RefResponse, error) {
	values, err := r.cli.MGet(r.cfg.PairKey, r.cfg.CcyKey, r.cfg.CfgKey, r.cfg.NewsKey).Result()
	if err != nil {
		return nil, err
	}

	var result []domain.RefResponse
	for _, val := range values {
		var ref domain.RefResponse
		if str, ok := val.(string); ok {
			err := json.Unmarshal([]byte(str), &ref)
			if err != nil {
				logger.WithField("val", val).Error("error parsing ref")
				continue
			}
			result = append(result, ref)
		}
	}
	return result, nil

}
