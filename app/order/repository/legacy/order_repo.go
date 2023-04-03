package legacy

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"web_server/domain"
)

type onboardingLegacyRepo struct {
	client  *http.Client
	baseURL string
}

func NewOnboardingRepository(client *http.Client, baseURL string) domain.OrderRepository {
	return &onboardingLegacyRepo{client: client, baseURL: baseURL}
}

// Amend implements domain.OrderRepository
func (o *onboardingLegacyRepo) Amend(ctx context.Context, req *domain.BaseRequest) (*domain.OrderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ORDER_AMEND, bytes.NewBuffer(body))
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

	var result domain.OrderResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Cancel implements domain.OrderRepository
func (o *onboardingLegacyRepo) Cancel(ctx context.Context, req *domain.BaseRequest) (*domain.OrderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ORDER_CANCEL, bytes.NewBuffer(body))
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

	var result domain.OrderResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Create implements domain.OrderRepository
func (o *onboardingLegacyRepo) Create(ctx context.Context, req *domain.BaseRequest) (*domain.OrderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ORDER_CREATE, bytes.NewBuffer(body))
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

	var result domain.OrderResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDetail implements domain.OrderRepository
func (o *onboardingLegacyRepo) GetDetail(ctx context.Context, req *domain.BaseRequest) (*domain.OrderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ORDER_DETAIL, bytes.NewBuffer(body))
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

	var result domain.OrderResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetStatus implements domain.OrderRepository
func (o *onboardingLegacyRepo) GetStatus(ctx context.Context, req *domain.BaseRequest) (*domain.OrderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ORDER_STATUS, bytes.NewBuffer(body))
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

	var result domain.OrderResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
