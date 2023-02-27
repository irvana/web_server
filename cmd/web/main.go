package main

import (
	"flag"
	"web_server/pkg/app"
	"web_server/pkg/app/consumer"
	"web_server/pkg/app/publisher"
	"web_server/pkg/authentication"
	commonlog "web_server/pkg/common/log"
	"web_server/pkg/redis"
	"web_server/pkg/wamp"

	log "github.com/sirupsen/logrus"
)

var cfg Config

type Config struct {
	Redis   redis.Config          `mapstructure:"redis"`
	Auth    authentication.Config `mapstructure:"authentication"`
	WampCfg wamp.Config           `mapstructure:"wamp"`
}

func init() {
	configPath := flag.String("f", "configs", "config file path")
	flag.Parse()

	log.Info("initiating config file from", *configPath)
	v, err := commonlog.InitConfig(*configPath, "config")
	if err != nil {
		log.WithError(err).Panic("Error unmarshal config file")
	}
	err = v.Unmarshal(&cfg)
	if err != nil {
		log.WithError(err).Panic("Error unmarshal config file")
	}
}

func main() {
	// init redis
	log.Info("initializing redis client")
	redisCli := redis.InitClient(cfg.Redis)

	// init data ref
	log.Info("initializing reference data")
	_, err := app.InitReference(redisCli)
	if err != nil {
		log.WithError(err).Panic("Error collecting REF data")
	}

	// init auth for wamp
	log.Info("initializing authentication module data")
	auth, err := authentication.InitAuthModule(cfg.Auth, redisCli.Client)
	if err != nil {
		log.WithError(err).Panic("Error initiating authentication module")
	}

	// init wamp router
	log.Info("initializing WAMP router")
	wampWss, err := wamp.InitWamp(cfg.WampCfg, auth, redisCli)
	if err != nil {
		log.WithError(err).Panic("Error initiating websocket")
	}

	log.Info("starting rates & ref publisher")
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
