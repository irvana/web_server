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
	anonymousAuth := true
	if withAuth {
		keystore := NewKeyStore(redisCli.Client)
		authenticator := InitAuthenticator(keystore, 1*time.Second, authentication)
		authenticators = append(authenticators, authenticator)
		anonymousAuth = false
	}

	nxr, err := router.NewRouter(&router.Config{
		RealmConfigs: []*router.RealmConfig{
			{
				URI:            wamp.URI(realm),
				AnonymousAuth:  anonymousAuth,
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

	subscribeMetaOnJoin(cli, log.StandardLogger())
	subscribeMetaOnLeave(cli, log.StandardLogger())

	return cli, err
}

func subscribeMetaOnJoin(subscriber *client.Client, logger *log.Logger) {
	// Define function to handle on_join events received.
	onJoin := func(event *wamp.Event) {
		if len(event.Arguments) != 0 {
			if details, ok := wamp.AsDict(event.Arguments[0]); ok {
				onJoinID, _ := wamp.AsID(details["session"])
				authid, _ := wamp.AsString(details["authid"])
				logger.Printf("Client %v joined realm (authid=%s)\n", onJoinID, authid)
				return
			}
		}
		logger.Println("Client joined realm - no data provided")
	}

	// Subscribe to on_join topic.
	err := subscriber.Subscribe(string(wamp.MetaEventSessionOnJoin), onJoin, nil)
	if err != nil {
		logger.Fatal("subscribe error:", err)
	}
	logger.Println("Subscribed to", string(wamp.MetaEventSessionOnJoin))
}

func subscribeMetaOnLeave(subscriber *client.Client, logger *log.Logger) {
	// Define function to handle on_leave events received.
	onLeave := func(event *wamp.Event) {
		if len(event.Arguments) != 0 {
			if id, ok := wamp.AsID(event.Arguments[0]); ok {
				logger.Println("Client", id, "left realm")
				return
			}
		}
		logger.Println("A client left the realm")
	}

	// Subscribe to on_leave topic.
	err := subscriber.Subscribe(string(wamp.MetaEventSessionOnLeave), onLeave, nil)
	if err != nil {
		logger.Fatal("subscribe error:", err)
	}
	logger.Println("Subscribed to", string(wamp.MetaEventSessionOnLeave))
}
