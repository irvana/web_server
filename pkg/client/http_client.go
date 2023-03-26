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
	OnboardingTimeout time.Duration `mapstructure:"onboardingTimeout"`
	LegacyBaseURL     string        `mapstructure:"legacyBaseUrl"`
}

func InitHttpClient(cfg Config) *HttpClient {
	return &HttpClient{
		Client:  &http.Client{Timeout: cfg.OnboardingTimeout},
		BaseURL: cfg.LegacyBaseURL,
	}
}
