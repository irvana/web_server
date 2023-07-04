package repository

import (
	"context"
	"encoding/json"
	"strings"
	"web_server/domain"
	"web_server/pkg/redis"
	"web_server/pkg/wamp"

	gammazero "github.com/gammazero/nexus/v3/wamp"
	"github.com/sirupsen/logrus"
)

type rateRepository struct {
	redisCli *redis.Client
	wampCli  wamp.Wamp
}

func NewRateRepository(redisCli *redis.Client, wampCli wamp.Wamp) domain.RateRepository {
	return &rateRepository{redisCli, wampCli}
}

// ConsumeRate implements domain.RateRepository
func (r *rateRepository) ConsumeRate(ctx context.Context, output chan<- domain.RateResponse) {
	pubSubCli := r.redisCli.RatesPubSub
	if val, ok := ctx.Value("channel").(string); ok && strings.EqualFold("ON_RATE", val) {
		pubSubCli = r.redisCli.AuthRatesPubSub
	}

	for msg := range pubSubCli.Channel() {
		logrus.WithField("payload", msg.Payload).WithField("channel", ctx.Value("channel")).Info("consuming payload")
		var resp domain.RateResponse
		err := json.Unmarshal([]byte(msg.Payload), &resp)
		if err != nil {
			logrus.WithError(err).WithField("payload", msg.Payload).Error("error unmarshal rate payload")
			continue
		}
		output <- resp
	}

}

// PublishRate implements domain.RateRepository
func (r *rateRepository) PublishRate(ctx context.Context, rate domain.RateResponse) error {
	if chn, ok := ctx.Value("channel").(string); ok {
		logrus.WithField("rate", &rate).WithField("channel", ctx.Value("channel")).Info("passing rate")
		str, err := json.Marshal(rate)
		if err != nil {
			logrus.WithField("rate", &rate).WithField("channel", ctx.Value("channel")).Error("error serializing")
			return err
		}

		if strings.EqualFold("ON_RATE", chn) {
			return r.wampCli.AuthClient.Publish(chn, gammazero.Dict{}, gammazero.List{str}, nil)
		}
		return r.wampCli.Client.Publish(chn, gammazero.Dict{}, gammazero.List{str}, nil)
	}
	logrus.WithField("rate", &rate).WithField("channel", ctx.Value("channel")).Info("channel not found")
	return nil
}
