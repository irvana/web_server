package legacy

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"web_server/domain"
)

const (
	TRANSACTION_DETAIL = "/transaction/detail"
	TRANSACTION_DEAL   = "/transaction/deal"
)

type trxLegacyRepo struct {
	client  *http.Client
	baseURL string
}

func NewTrxRepository(client *http.Client, baseURL string) domain.TransactionRepository {
	return &trxLegacyRepo{client: client, baseURL: baseURL}
}
func (s *trxLegacyRepo) GetDetail(ctx context.Context, req *domain.BaseRequest) (*domain.TransactionResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.baseURL+TRANSACTION_DETAIL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result domain.TransactionResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *trxLegacyRepo) Deal(ctx context.Context, req *domain.BaseRequest) (*domain.TransactionResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.baseURL+TRANSACTION_DEAL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result domain.TransactionResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
