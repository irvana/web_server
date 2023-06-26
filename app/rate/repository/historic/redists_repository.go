package historic

import (
	"context"
	"strconv"
	"web_server/domain"

	"github.com/rueian/rueidis"
	"github.com/sirupsen/logrus"
)

type redistsRepository struct {
	cli rueidis.Client
}

func NewRedistsRepository(redisTsCli rueidis.Client) domain.HistoricRates {
	return &redistsRepository{redisTsCli}
}

func (r *redistsRepository) GetChart(ctx context.Context, req *domain.HistoricParam) (*domain.HistoricRateResponse, error) {
	res := r.cli.Do(ctx, r.cli.B().TsRange().Key(req.Key).Fromtimestamp(req.Start).Totimestamp(req.End).Build())

	c, err := res.ToArray()
	if err != nil {
		logrus.WithError(err).Error("error converting all elements")
		return nil, err
	}

	pairs := []domain.RatePair{}
	for _, v := range c {
		arr, err := v.ToArray()
		if err != nil {
			logrus.WithError(err).Error("error converting single element")
			continue
		}
		if len(arr) == 2 {
			key, err := arr[0].ToInt64()
			if err != nil {
				logrus.WithError(err).Error("error converting stamp element")
				return nil, err
			}
			val, err := arr[1].ToString()
			if err != nil {
				logrus.WithError(err).Error("error converting rate element")
				return nil, err
			}
			pairs = append(pairs, domain.RatePair{Stamp: strconv.FormatInt(key, 10), Rate: val})
		}

	}

	return &domain.HistoricRateResponse{Ts: pairs}, nil
}
