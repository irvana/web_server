package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"web_server/app/consumer"
	"web_server/app/handler"
	"web_server/app/onboarding/handler"
	"web_server/app/onboarding/repository/legacy"
	"web_server/app/onboarding/usecase"
	"web_server/app/publisher"
	"web_server/pkg/authentication"
	"web_server/pkg/client"
	commonlog "web_server/pkg/log"
	"web_server/pkg/redis"
	"web_server/pkg/wamp"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var cfg Config

type Config struct {
	Redis      redis.Config          `mapstructure:"redis"`
	Auth       authentication.Config `mapstructure:"authentication"`
	WampCfg    wamp.Config           `mapstructure:"wamp"`
	Server     handler.Config        `mapstructure:"server"`
	HttpClient client.Config         `mapstructure:"client"`
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

	// // init data ref
	// log.Info("initializing reference data")
	// _, err := app.InitReference(redisCli)
	// if err != nil {
	// 	log.WithError(err).Panic("Error collecting REF data")
	// }

	// init auth module
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

	httpClient := client.InitHttpClient(cfg.HttpClient)
	gin := gin.Default()

	obRepository := legacy.NewOnboardingRepository(httpClient.Client)
	obUsecase := usecase.NewOnboardingUsecase(obRepository)
	handler.NewOnboardingHandler(obUsecase, gin)

	// log.Info("starting rates & ref publisher")
	// initAndRunPublishers(wampWss, redisCli)

	// router := handler.SetupHandlers(cfg.Server, wampWss.Wss)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Server.Address, cfg.Server.Port),
		Handler: gin,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func initAndRunPublishers(wampWss *wamp.Wamp, redis *redis.Client) {
	ratesConsumer := consumer.InitRedisConsumer(redis.RatesPubSub, "rates")
	ratesPublisher := publisher.InitPublisher(ratesConsumer, wampWss)
	go ratesPublisher.Publish()

	refConsumer := consumer.InitRedisConsumer(redis.MiscPubSub, "ref")
	refPublisher := publisher.InitPublisher(refConsumer, wampWss)
	go refPublisher.Publish()

}
