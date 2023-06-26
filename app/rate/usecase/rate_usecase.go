package usecase

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"web_server/domain"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type (
	rateUsecase struct {
		rateRepo domain.RateRepository
		histRepo domain.HistoricRates
		cfg      HistoricConfig
	}
	HistoricConfig struct {
		KeyMap map[int64]string `mapstructure:"mapKey"`
	}
)

func NewRateUsecase(rateRepo domain.RateRepository, histRepo domain.HistoricRates, cfg HistoricConfig) domain.RateUsecase {
	return &rateUsecase{rateRepo, histRepo, cfg}
}

// PublishRate implements domain.RateUsecase
func (r *rateUsecase) ProcessBackgroundRate(ctx context.Context) {
	log.WithField("foo", ctx.Value("foo")).Info("running rate publisher")
	for {
		resp, err := r.rateRepo.ConsumeRate(ctx)
		if err != nil {
			log.WithError(err).Errorln("error consuming rate")
		}

		err = r.rateRepo.PublishRate(ctx, resp)
		if err != nil {
			log.WithError(err).Errorln("error publishing rate")
		}
	}

}

// PublishReference implements domain.RateUsecase
func (r *rateUsecase) ProcessBackgroundRef(ctx context.Context) {
	log.WithField("foo", ctx.Value("foo")).Info("running ref publisher")
	for {
		resp, err := r.rateRepo.ConsumeRate(ctx)
		if err != nil {
			log.WithError(err).Errorln("error consuming ref")
		}

		err = r.rateRepo.PublishRate(ctx, resp)
		if err != nil {
			log.WithError(err).Errorln("error publishing ref")
		}
	}
}

func (r *rateUsecase) GetChart(ctx context.Context, req *domain.HistoricRateRequest) (*domain.HistoricRateResponse, error) {
	now := time.Now().UnixMilli()
	lastMillis, err := convertToDurationMilli(strings.ToLower(req.Period))
	if err != nil {
		return nil, err
	}
	period := now - lastMillis
	key := filterKey(r.cfg.KeyMap, lastMillis)
	formattedKey := fmt.Sprintf(key, "1", req.PairId)

	logrus.WithField("key", formattedKey).WithField("start", period).WithField("end", now).Info("getting historic pairs")
	result, err := r.histRepo.GetChart(ctx, &domain.HistoricParam{Start: period, End: now, Key: formattedKey})
	if err != nil {
		logrus.WithError(err).Error("error getting chart")
		return nil, err
	}

	result.PairID = req.PairId
	result.Period = req.Period
	return result, nil
}

func filterKey(m map[int64]string, input int64) string {
	minValue := int64(math.MaxInt64)
	minKey := ""

	for key, value := range m {
		if key > input && key < minValue {
			minValue = key
			minKey = value
		}
	}

	return minKey
}

func convertToDurationMilli(input string) (int64, error) {
	units := map[string]time.Duration{
		"h": time.Hour,
		"d": time.Hour * 24,
		"w": time.Hour * 24 * 7,
		"m": time.Hour * 24 * 30,
		"y": time.Hour * 24 * 365,
	}

	unit := input[len(input)-1:]
	multiplier := input[:len(input)-1]

	durationUnit, found := units[unit]
	if !found {
		return 0, fmt.Errorf("unsupported duration format: %s", input)
	}

	value, err := strconv.Atoi(multiplier)
	if err != nil {
		return 0, fmt.Errorf("failed to parse duration value: %s", err)
	}

	duration := time.Duration(value) * durationUnit

	return duration.Milliseconds(), nil
}
