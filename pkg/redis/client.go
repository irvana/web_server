package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type (
	Config struct {
		Address      string `mapstructure:"address"`
		DialTimeout  int    `mapstructure:"timeout"`
		RatesChannel string `mapstructure:"ratesChannel"`
		MiscChannel  string `mapstructure:"miscChannel"`
	}

	// Client do we need this?
	Client struct {
		Client      *redis.Client
		RatesPubSub *redis.PubSub
		MiscPubSub  *redis.PubSub
	}
)

// InitClient will return new sentinel client & also set the default pubsub
func InitClient(cfg Config) *Client {
	client := redis.NewClient(&redis.Options{
		Addr:        cfg.Address,
		DialTimeout: time.Duration(cfg.DialTimeout) * time.Second,
		DB:          1,
	})
	ratesPb := client.Subscribe(cfg.RatesChannel)
	miscPb := client.Subscribe(cfg.MiscChannel)
	return &Client{client, ratesPb, miscPb}
}
