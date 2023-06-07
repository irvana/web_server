package legacy

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"web_server/domain"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

func (s *statementLegacyRepo) List(ctx context.Context, req *domain.StatementRequest) ([]domain.StatementResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, s.baseURL+STATEMENT_LIST, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if reqCtx, ok := ctx.(*gin.Context); ok {
		httpReq.Header = reqCtx.Request.Header
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
		logrus.WithError(err).WithField("resp", string(respByte)).Error("failed unmarshal response")
		return nil, err
	}

	return result, nil
}
