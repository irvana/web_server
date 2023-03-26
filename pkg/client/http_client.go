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
	OnboardingTimeout time.Duration `mapstructure:"onboardingTimeoutMillis"`
	LegacyBaseURL     string        `mapstructure:"legacyBaseUrl"`
}

func InitHttpClient(cfg Config) *HttpClient {
	return &HttpClient{
		Client:  &http.Client{Timeout: cfg.OnboardingTimeout * time.Millisecond},
		BaseURL: cfg.LegacyBaseURL,
	}
}
