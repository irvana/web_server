package redis_repo

import (
	"encoding/json"
	"web_server/domain"
	rediscfg "web_server/pkg/redis"

	"github.com/go-redis/redis"
	logger "github.com/sirupsen/logrus"
)

type refRedisRepo struct {
	cli       *redis.Client
	pubSubCli *redis.PubSub
	cfg       rediscfg.Config
}

func NewRefRedisRepository(cli *rediscfg.Client, cfg rediscfg.Config) domain.RefSubscriberRepository {
	return &refRedisRepo{cli.Client, cli.MiscPubSub, cfg}
}

func (r *refRedisRepo) Consume() <-chan *redis.Message {
	return r.pubSubCli.Channel()
}

func (r *refRedisRepo) GetAll() ([]domain.RefResponse, error) {
	keys := make([]string, 0, len(r.cfg.RefsMap))
	for key := range r.cfg.RefsMap {
		keys = append(keys, key)
	}

	values, err := r.cli.MGet(keys...).Result()
	if err != nil {
		logger.WithError(err).Error("error getting config")
		return nil, err
	}
	var result []domain.RefResponse

	for i, value := range values {
		var ref []domain.RefDetails
		if str, ok := value.(string); ok {
			err := json.Unmarshal([]byte(str), &ref)
			if err != nil {
				logger.WithField("val", value).Error("error parsing ref")
				continue
			}
		}
		result = append(result, domain.RefResponse{Type: r.cfg.RefsMap[keys[i]], Results: ref})
	}

	return result, nil

}
