package historic

import (
	"context"
	"web_server/domain"

	"github.com/rueian/rueidis"
)

type redistsRepository struct {
	cli rueidis.Client
}

func NewRedistsRepository(redisTsCli rueidis.Client) domain.HistoricRates {
	return &redistsRepository{redisTsCli}
}

func (r *redistsRepository) GetChart(ctx context.Context, req *domain.HistoricRateRequest) ([]domain.RateResponse, error) {
	// y := r.cli.B().TsRange().Key().Fromtimestamp().Totimestamp().
	// y

	return nil, nil
}
