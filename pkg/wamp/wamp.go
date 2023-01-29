package wamp

import (
	"net/http"
	"sync"
	auth "web_server/pkg/authentication"

	"github.com/gammazero/nexus/v3/client"
	"github.com/gammazero/nexus/v3/router"
	"github.com/gammazero/nexus/v3/wamp"
	log "github.com/sirupsen/logrus"
)

type (
	Wamp struct {
		wss    *router.WebsocketServer
		client *client.Client
	}

	Config struct {
		Realm string
	}
)

var (
	onceWamp sync.Once
	wmp      *Wamp
)

func InitWamp(cfg Config, authentication *auth.Authentication) (*Wamp, error) {
	var errs error
	onceWamp.Do(func() {
		nxr, err := router.NewRouter(&router.Config{
			RealmConfigs: []*router.RealmConfig{
				{
					URI:            wamp.URI(cfg.Realm),
					AnonymousAuth:  true,
					EnableMetaKill: true,
				},
			},
		}, nil)
		if err != nil {
			log.WithError(err).Errorln("Error initiating router")
			wmp, errs = nil, err
			return
		}

		wss := router.NewWebsocketServer(nxr)
		wss.Upgrader.CheckOrigin = func(req *http.Request) bool {
			return true
		}

		cli, err := initClient(nxr, cfg.Realm)
		if err != nil {
			log.WithError(err).Errorln("Error initiating router")
			wmp, errs = nil, err
			return
		}

		wmp = &Wamp{wss, cli}
	})

	return wmp, errs
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
