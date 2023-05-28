package legacy

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"web_server/domain"
)

type refLegacyRepo struct {
	client  *http.Client
	baseURL string
}

const (
	REF_ALL  = "/ref/all"
	REF_CCY  = "/ref/ccy"
	REF_PAIR = "/ref/pair"
	REF_NEWS = "/ref/news"
	REF_CFG  = "/ref/cfg"
)

func NewRefRepository(client *http.Client, baseURL string) domain.RefRepository {
	return &refLegacyRepo{client: client, baseURL: baseURL}
}

func (r *refLegacyRepo) GetAll(ctx context.Context, req *domain.RefRequest) ([]domain.RefResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, r.baseURL+REF_ALL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []domain.RefResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *refLegacyRepo) GetCurrency(ctx context.Context, req *domain.RefRequest) ([]domain.Currency, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, r.baseURL+REF_CCY, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []domain.Currency
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *refLegacyRepo) GetPair(ctx context.Context, req *domain.RefRequest) ([]domain.Pair, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, r.baseURL+REF_PAIR, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []domain.Pair
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *refLegacyRepo) GetNews(ctx context.Context, req *domain.RefRequest) ([]domain.News, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, r.baseURL+REF_NEWS, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []domain.News
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *refLegacyRepo) GetConfig(ctx context.Context, req *domain.RefRequest) (*domain.Config, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, r.baseURL+REF_CFG, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result domain.Config
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
