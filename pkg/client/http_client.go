package client

import (
	"net/http"
	"time"
)

type HttpClient struct {
	Client  *http.Client
	BaseURL string
}

type Config struct {
	Timeout       time.Duration `mapstructure:"timeout"`
	LegacyBaseURL string        `mapstructure:"legacyBaseUrl"`
}

func InitHttpClient(cfg Config) *HttpClient {
	return &HttpClient{
		Client:  &http.Client{Timeout: cfg.Timeout},
		BaseURL: cfg.LegacyBaseURL,
	}
}
