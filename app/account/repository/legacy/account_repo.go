package legacy

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"web_server/domain"
)

type accountLegacyRepo struct {
	client  *http.Client
	baseURL string
}

func NewOnboardingRepository(client *http.Client, baseURL string) domain.AccountRepository {
	return &accountLegacyRepo{client, baseURL}
}

// AtmGetInfo implements domain.OnboardingRepository
func (o *accountLegacyRepo) GetList(ctx context.Context, req *domain.BaseRequest) ([]domain.AccountResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ACCOUNT_LIST, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := o.client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []domain.AccountResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// AtmOTP implements domain.OnboardingRepository
func (o *accountLegacyRepo) Open(ctx context.Context, req *domain.BaseRequest) (*domain.AccountOpenResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ACCOUNT_OPEN, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := o.client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result domain.AccountOpenResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
