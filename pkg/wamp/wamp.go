package wamp

import (
	"net/http"
	"sync"
	"time"
	auth "web_server/pkg/authentication"
	"web_server/pkg/redis"

	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/router"
	wampauth "github.com/gammazero/nexus/v3/router/auth"
	"github.com/gammazero/nexus/v3/wamp"
	log "github.com/sirupsen/logrus"
)

type (
	Wamp struct {
		Wss        *router.WebsocketServer
		AuthWss    *router.WebsocketServer
		Client     *client.Client
		AuthClient *client.Client
	}

	Config struct {
		Realm              string `mapstructure:"realm"`
		AuthenticatedRealm string `mapstructure:"authenticatedRealm"`
		Path               string `mapstructure:"path"`
		AuthPath           string `mapstructure:"authPath"`
	}
)

var (
	onceWamp sync.Once
	wmp      *Wamp
)

func InitWamp(cfg Config, authentication *auth.Authentication, redisCli *redis.Client) (wmp *Wamp, errs error) {
	onceWamp.Do(func() {
		authWss, authCli, errs := initModules(authentication, redisCli, true, cfg.AuthenticatedRealm)
		if errs != nil {
			return
		}

		wss, cli, errs := initModules(nil, nil, false, cfg.Realm)
		if errs != nil {
			return
		}
		wmp = &Wamp{Wss: wss, Client: cli, AuthWss: authWss, AuthClient: authCli}
	})

	return wmp, errs
}

func initModules(authentication *auth.Authentication, redisCli *redis.Client, withAuth bool, realm string) (*router.WebsocketServer, *client.Client, error) {
	var authenticators []wampauth.Authenticator
	if withAuth {
		keystore := NewKeyStore(redisCli.Client)
		authenticator := InitAuthenticator(keystore, 1*time.Second, authentication)
		authenticators = append(authenticators, authenticator)
	}

	nxr, err := router.NewRouter(&router.Config{
		RealmConfigs: []*router.RealmConfig{
			{
				URI:            wamp.URI(realm),
				AnonymousAuth:  true,
				EnableMetaKill: true,
				Authenticators: authenticators,
			},
		},
	}, log.StandardLogger())
	if err != nil {
		log.WithError(err).Errorln("Error initiating router")
		return nil, nil, err

	}

	wss := router.NewWebsocketServer(nxr)
	wss.Upgrader.CheckOrigin = func(req *http.Request) bool {
		return true
	}

	cli, err := initClient(nxr, realm)
	if err != nil {
		log.WithError(err).Errorln("Error initiating client")
		return nil, nil, err
	}
	return wss, cli, nil
}

func initClient(router router.Router, realm string) (*client.Client, error) {
	cfg := client.Config{Realm: realm, Logger: log.StandardLogger()}
	cli, err := client.ConnectLocal(router, cfg)
	if err != nil {
		log.WithError(err).Errorln("Error connecting through local client")
		return nil, err
	}

	return cli, err
}
