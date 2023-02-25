package main

import (
	"web_server/pkg/app"
	"web_server/pkg/app/consumer"
	"web_server/pkg/app/publisher"
	"web_server/pkg/authentication"
	commonlog "web_server/pkg/common/log"
	"web_server/pkg/redis"
	"web_server/pkg/wamp"

	log "github.com/sirupsen/logrus"
)

var cfg MainConfig

type MainConfig struct {
	redis   redis.Config          `mapstructure:"redis"`
	auth    authentication.Config `mapstructure:"authentication"`
	wampCfg wamp.Config           `mapstructure:"wamp"`
}

func init() {
	v, _ := commonlog.InitConfig("configs", "config")
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.WithError(err).Panic("Error unmarshal config file")
	}
}

func main() {
	// init redis
	log.Info("initializing redis client")
	redisCli := redis.InitClient(cfg.redis)

	// init data ref
	log.Info("initializing reference data")
	_, err := app.InitReference(redisCli)
	if err != nil {
		log.WithError(err).Panic("Error collecting REF data")
	}

	// init auth for wamp
	auth, err := authentication.InitAuthModule(cfg.auth, redisCli.Client)
	if err != nil {
		log.WithError(err).Panic("Error initiating authentication module")
	}

	// init wamp router
	wampWss, err := wamp.InitWamp(cfg.wampCfg, auth, redisCli)
	if err != nil {
		log.WithError(err).Panic("Error initiating websocket")
	}

	initAndRunPublisher(wampWss, redisCli)

	// listen and serve server
	println("running")
}

func initAndRunPublisher(wampWss *wamp.Wamp, redis *redis.Client) {
	ratesConsumer := consumer.InitRedisConsumer(redis.RatesPubSub, "rates")
	ratesPublisher := publisher.InitPublisher(ratesConsumer, wampWss)
	go ratesPublisher.Publish()

	refConsumer := consumer.InitRedisConsumer(redis.MiscPubSub, "ref")
	refPublisher := publisher.InitPublisher(refConsumer, wampWss)
	go refPublisher.Publish()

}
