package client

import (
	"net/http"
	"time"
)

type HttpClient struct {
	Client *http.Client
}

type Config struct {
	Timeout time.Duration `mapstructure:"timeout"`
}

func InitHttpClient(cfg Config) *HttpClient {
	return &HttpClient{
		Client: &http.Client{Timeout: cfg.Timeout},
	}
}
