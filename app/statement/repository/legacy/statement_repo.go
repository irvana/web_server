package legacy

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"web_server/domain"
)

type statementLegacyRepo struct {
	client  *http.Client
	baseURL string
}

const (
	STATEMENT_LIST = "/statement/list"
)

func NewStatementRepository(client *http.Client, baseURL string) domain.StatementRepository {
	return &statementLegacyRepo{client: client, baseURL: baseURL}
}

func (s *statementLegacyRepo) List(ctx context.Context, req *domain.BaseRequest) ([]domain.StatementResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.baseURL+STATEMENT_LIST, bytes.NewBuffer(body))
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

	var result []domain.StatementResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
