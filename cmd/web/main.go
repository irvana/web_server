package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	obhandler "web_server/app/onboarding/handler"
	"web_server/app/onboarding/repository/legacy"
	obusecase "web_server/app/onboarding/usecase"
	ratehandler "web_server/app/rate/handler"
	"web_server/app/rate/repository"
	rateusecase "web_server/app/rate/usecase"
	"web_server/configs"
	"web_server/pkg/authentication"
	"web_server/pkg/client"
	commonlog "web_server/pkg/logger"
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
	Server     configs.Config        `mapstructure:"server"`
	HttpClient client.Config         `mapstructure:"client"`
}

func init() {
	configPath := flag.String("f", "configs", "config file path")
	flag.Parse()

	log.Info("initiating config file from: ", *configPath)
	v, err := commonlog.InitConfig(*configPath, "config")
	// TODO separate log & config
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

	log.Info("initializing redis timeseries client")
	redisTsCli, err := redis.InitTimeSeriesClient(cfg.Redis)
	if err != nil {
		log.WithError(err).Panic("error initiating authentication module")
	}

	// TODO: data ref from redis usecase

	// init auth module
	// TODO move auth to usecase
	log.Info("initializing authentication module data")
	auth, err := authentication.InitAuthModule(cfg.Auth, redisCli.Client)
	if err != nil {
		log.WithError(err).Panic("error initiating authentication module")
	}

	// init wamp router
	log.Info("initializing WAMP router")
	wampWss, err := wamp.InitWamp(cfg.WampCfg, auth, redisCli)
	if err != nil {
		log.WithError(err).Panic("error initiating websocket")
	}

	log.Info("attaching websocket server to webserver router")
	gin.SetMode(cfg.Server.Level)
	router := gin.New()
	gin.DefaultWriter = log.StandardLogger().Writer()
	router.GET(cfg.WampCfg.Path, func(ctx *gin.Context) {
		wampWss.Wss.ServeHTTP(ctx.Writer, ctx.Request)
	})
	router.GET(cfg.WampCfg.AuthPath, func(ctx *gin.Context) {
		wampWss.AuthWss.ServeHTTP(ctx.Writer, ctx.Request)
	})

	httpClient := client.InitHttpClient(cfg.HttpClient)
	obRepository := legacy.NewOnboardingRepository(httpClient.Client, httpClient.BaseURL)
	obUsecase := obusecase.NewOnboardingUsecase(obRepository)
	obhandler.NewOnboardingHandler(obUsecase, router)

	rateRepo := repository.NewRateRepository(redisCli.Client, wampWss.Client, redisTsCli)
	rateUsecase := rateusecase.NewRateUsecase(rateRepo)
	ratehandler.RunRateHandlers(rateUsecase)

	runServer(router)
}

func runServer(router *gin.Engine) {
	log.Infof("starting webserver at %s:%d", cfg.Server.Address, cfg.Server.Port)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Server.Address, cfg.Server.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatal("listen")
		}
	}()

	<-ctx.Done()

	stop()
	log.Info("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.WithError(err).Panic("server forced to shutdown: ")
	}

	log.Info("server exiting")
}
