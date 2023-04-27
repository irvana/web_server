package redis

import (
	"github.com/rueian/rueidis"
)

func InitTimeSeriesClient(cfg Config) (*rueidis.Client, error) {
	cli, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{cfg.Address},
	})
	if err != nil {
		return nil, err
	}
	return &cli, err
}
