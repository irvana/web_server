package usecase

import (
	"context"
	"encoding/json"
	"sync"
	"web_server/domain"

	log "github.com/sirupsen/logrus"
)

type (
	refUsecase struct {
		ref    domain.RefRepository
		refCfg domain.RefSubscriberRepository
		config []domain.RefResponse
	}
)

var (
	ref     refUsecase
	onceCfg sync.Once
	mtx     sync.RWMutex
)

func NewRefUsecase(refRepo domain.RefRepository, refCf domain.RefSubscriberRepository) (domain.RefUsecase, error) {
	onceCfg.Do(func() {
		ref.ref = refRepo
		ref.refCfg = refCf

		log.Info("initiating REF configuration")
		cfg, err := ref.refCfg.GetAll()
		if err != nil {
			log.WithError(err).Errorln("error initiating new config")
		}
		ref.config = cfg

		go ref.UpdateConfig()
	})
	return &ref, nil
}

func (r *refUsecase) GetAll(ctx context.Context, req *domain.RefRequest) ([]domain.RefResponse, error) {
	return r.config, nil
}

func (r *refUsecase) GetCurrency(ctx context.Context, req *domain.RefRequest) ([]domain.Currency, error) {
	for _, v := range r.config {
		if v.Type == "currency" {
			var result []domain.Currency
			for _, resp := range v.Results {
				result = append(result, domain.Currency{
					ID:      resp.ID,
					Label:   resp.Label,
					Name:    resp.Name,
					Decimal: resp.Decimal,
				})
			}
			return result, nil
		}
	}
	log.Debug("getcurrency not found")
	return nil, nil
}

func (r *refUsecase) GetPair(ctx context.Context, req *domain.RefRequest) ([]domain.Pair, error) {
	for _, v := range r.config {
		if v.Type == "pair" {
			var result []domain.Pair
			for _, resp := range v.Results {
				result = append(result, domain.Pair{
					ID:         resp.ID,
					Label:      resp.Label,
					BaseCcyID:  resp.BaseCcyID,
					QuoteCcyID: resp.QuoteCcyID,
					Decimal:    resp.Decimal,
					Tradeable:  resp.Tradeable,
				})
			}
			return result, nil
		}
	}
	log.Debug("pair config not found")
	return nil, nil
}

func (r *refUsecase) GetNews(ctx context.Context, req *domain.RefRequest) ([]domain.News, error) {
	for _, v := range r.config {
		if v.Type == "mobilenewsitem" {
			var result []domain.News
			for _, resp := range v.Results {
				result = append(result, domain.News{
					ID:        resp.ID,
					NewsID:    resp.NewsID,
					PairID:    resp.PairID,
					Stamp:     resp.Stamp,
					Title:     resp.Title,
					Content:   resp.Content,
					StartDate: resp.StartDate,
					EndDate:   resp.EndDate,
				})
			}
			return result, nil
		}
	}
	log.Debug("news config not found")
	return nil, nil
}

func (r *refUsecase) GetConfig(ctx context.Context, req *domain.RefRequest) ([]domain.Config, error) {
	for _, v := range r.config {
		if v.Type == "config" {
			var result []domain.Config
			for _, resp := range v.Results {
				result = append(result, domain.Config{
					ID:         resp.ID,
					AutoLocked: resp.AutoLocked,
				})
			}
			return result, nil
		}
	}
	log.Debug("news config not found")
	return nil, nil
}

func (r *refUsecase) UpdateConfig() {
	log.Info("setting up go routine for config")
	for rr := range r.refCfg.Consume() {
		log.WithField("message", rr.Payload).Info("getting incoming ref")
		var n domain.RefResponse
		err := json.Unmarshal([]byte(rr.Payload), &n)
		if err != nil {
			log.WithError(err).WithField("payload", rr.Payload).Errorln("error updating new config with value ")
			continue
		}

		for i, cfg := range r.config {
			if cfg.Type == n.Type {
				mtx.Lock()
				r.config[i] = n
				mtx.Unlock()
			}
		}
	}
}
