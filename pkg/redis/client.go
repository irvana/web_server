package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type (
	Config struct {
		Address         string            `mapstructure:"address"`
		DialTimeout     int               `mapstructure:"timeout"`
		MiscChannel     string            `mapstructure:"miscChannel"`
		RateChannel     string            `mapstructure:"offRatesChannel"`
		AuthRateChannel string            `mapstructure:"onRatesChannel"`
		ConfigKey       string            `mapstructure:"configKey"`
		RefsMap         map[string]string `mapstructure:"refsMap"`
	}

	// Client do we need this?
	Client struct {
		Client          *redis.Client
		RatesPubSub     *redis.PubSub
		AuthRatesPubSub *redis.PubSub
		MiscPubSub      *redis.PubSub
	}
)

// InitClient will return new sentinel client & also set the default pubsub
func InitClient(cfg Config) *Client {
	client := redis.NewClient(&redis.Options{
		Addr:        cfg.Address,
		DialTimeout: time.Duration(cfg.DialTimeout) * time.Second,
	})
	ratesPb := client.Subscribe(cfg.RateChannel)
	authRatesPb := client.Subscribe(cfg.AuthRateChannel)
	miscPb := client.Subscribe(cfg.MiscChannel)
	return &Client{client, ratesPb, authRatesPb, miscPb}
}
