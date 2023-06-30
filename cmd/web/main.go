package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	acchandler "web_server/app/account/handler"
	accRepo "web_server/app/account/repository/legacy"
	accusecase "web_server/app/account/usecase"
	obhandler "web_server/app/onboarding/handler"
	obRepo "web_server/app/onboarding/repository/legacy"
	obusecase "web_server/app/onboarding/usecase"
	orhandler "web_server/app/order/handler"
	orrepo "web_server/app/order/repository/legacy"
	orusecase "web_server/app/order/usecase"
	ratehandler "web_server/app/rate/handler"
	"web_server/app/rate/repository"
	"web_server/app/rate/repository/historic"
	rateusecase "web_server/app/rate/usecase"
	refhandler "web_server/app/ref/handler"
	refrepo "web_server/app/ref/repository/legacy"
	redis_repo "web_server/app/ref/repository/redis"
	"web_server/app/ref/usecase"
	stmrepo "web_server/app/statement/repository/legacy"
	trxrepo "web_server/app/transaction/repository/legacy"
	"web_server/configs"
	"web_server/pkg/authentication"
	"web_server/pkg/client"
	commonlog "web_server/pkg/logger"
	"web_server/pkg/redis"
	"web_server/pkg/wamp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rueian/rueidis"
	log "github.com/sirupsen/logrus"
)

var cfg Config

type Config struct {
	Redis      redis.Config               `mapstructure:"redis"`
	Auth       authentication.Config      `mapstructure:"authentication"`
	WampCfg    wamp.Config                `mapstructure:"wamp"`
	Server     configs.Config             `mapstructure:"server"`
	HttpClient client.Config              `mapstructure:"client"`
	Historic   rateusecase.HistoricConfig `mapstructure:"historic"`
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
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowHeaders:  []string{"X-Requested-With", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"*"},
	}))
	router.GET(cfg.WampCfg.Path, func(c *gin.Context) {
		wampWss.Wss.ServeHTTP(c.Writer, c.Request)
	})
	router.GET(cfg.WampCfg.AuthPath, func(ctx *gin.Context) {
		wampWss.AuthWss.ServeHTTP(ctx.Writer, ctx.Request)
	})

	initHandlers(router, redisCli, redisTsCli, *wampWss)
	runServer(router)
}

func initHandlers(router *gin.Engine, redisCli *redis.Client, redisTsCli rueidis.Client, wampWss wamp.Wamp) {
	httpClient := client.InitHttpClient(cfg.HttpClient)
	obRepository := obRepo.NewOnboardingRepository(httpClient.Client, httpClient.BaseURL)
	obUsecase := obusecase.NewOnboardingUsecase(obRepository)
	obhandler.NewOnboardingHandler(obUsecase, router)

	accRepository := accRepo.NewOnboardingRepository(httpClient.Client, httpClient.BaseURL)
	acUsecase := accusecase.NewAccountUsecase(accRepository)
	acchandler.NewAccountHandler(acUsecase, router)

	orderRepo := orrepo.NewOrderRepository(httpClient.Client, httpClient.BaseURL)
	stmRepo := stmrepo.NewStatementRepository(httpClient.Client, httpClient.BaseURL)
	trxRepo := trxrepo.NewTrxRepository(httpClient.Client, httpClient.BaseURL)
	orUsecase := orusecase.NewOrderUsecase(orderRepo, stmRepo, trxRepo)
	orhandler.NewOrderHandler(orUsecase, router)

	refRepo := refrepo.NewRefRepository(httpClient.Client, httpClient.BaseURL)
	refRedisRepo := redis_repo.NewRefRedisRepository(redisCli, cfg.Redis)
	refUsecase, err := usecase.NewRefUsecase(refRepo, refRedisRepo)
	if err != nil {
		log.WithError(err).Panic("error initiating ref usecase config")
	}
	refhandler.NewRefHandler(refUsecase, router)

	rateRepo := repository.NewRateRepository(redisCli.Client, wampWss.Client, redisTsCli)
	histRepo := historic.NewRedistsRepository(redisTsCli)
	rateUsecase := rateusecase.NewRateUsecase(rateRepo, histRepo, cfg.Historic)
	ratehandler.RunRateHandlers(rateUsecase, router)
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
