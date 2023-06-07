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

type onboardingLegacyRepo struct {
	client  *http.Client
	baseURL string
}

func NewOrderRepository(client *http.Client, baseURL string) domain.OrderRepository {
	return &onboardingLegacyRepo{client: client, baseURL: baseURL}
}

// Amend implements domain.OrderRepository
func (o *onboardingLegacyRepo) Amend(ctx context.Context, req *domain.OrderRequest) (*domain.OrderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ORDER_AMEND, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if reqCtx, ok := ctx.(*gin.Context); ok {
		httpReq.Header = reqCtx.Request.Header
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
		logrus.WithError(err).WithField("resp", string(respByte)).Error("failed unmarshal response")
		return nil, err
	}

	return &result, nil
}

// Cancel implements domain.OrderRepository
func (o *onboardingLegacyRepo) Cancel(ctx context.Context, req *domain.OrderRequest) (*domain.OrderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ORDER_CANCEL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if reqCtx, ok := ctx.(*gin.Context); ok {
		httpReq.Header = reqCtx.Request.Header
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
		logrus.WithError(err).WithField("resp", string(respByte)).Error("failed unmarshal response")
		return nil, err
	}

	return &result, nil
}

// Create implements domain.OrderRepository
func (o *onboardingLegacyRepo) Create(ctx context.Context, req *domain.OrderRequest) (*domain.OrderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ORDER_CREATE, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if reqCtx, ok := ctx.(*gin.Context); ok {
		httpReq.Header = reqCtx.Request.Header
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
		logrus.WithError(err).WithField("resp", string(respByte)).Error("failed unmarshal response")
		return nil, err
	}

	return &result, nil
}

// GetDetail implements domain.OrderRepository
func (o *onboardingLegacyRepo) GetDetail(ctx context.Context, req *domain.OrderRequest) (*domain.OrderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ORDER_DETAIL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if reqCtx, ok := ctx.(*gin.Context); ok {
		httpReq.Header = reqCtx.Request.Header
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
		logrus.WithError(err).WithField("resp", string(respByte)).Error("failed unmarshal response")
		return nil, err
	}

	return &result, nil
}

// GetStatus implements domain.OrderRepository
func (o *onboardingLegacyRepo) GetStatus(ctx context.Context, req *domain.OrderRequest) ([]domain.OrderResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ORDER_STATUS, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if reqCtx, ok := ctx.(*gin.Context); ok {
		httpReq.Header = reqCtx.Request.Header
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

	var result []domain.OrderResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		logrus.WithError(err).WithField("resp", string(respByte)).Error("failed unmarshal response")
		return nil, err
	}

	return result, nil
}
