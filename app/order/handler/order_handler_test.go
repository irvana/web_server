package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"web_server/domain"
	"web_server/domain/mocks"

	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOrderHandler_GetStatus(t *testing.T) {
	var resp []domain.OrderResponse
	faker.FakeData(&resp)
	var bodyObj domain.OrderRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	orderUc := new(mocks.OrderUsecase)
	orderUc.On("GetStatus", mock.Anything, &bodyObj).Return(resp, nil)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, ORDER_STATUS, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ah := &OrderHandler{orderUc}
	e.POST(ORDER_STATUS, ah.GetStatus)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Greater(t, len(string(respBody)), 10)

	orderUc.AssertExpectations(t)
}

func TestOrderHandler_Cancel(t *testing.T) {
	var resp domain.OrderResponse
	faker.FakeData(&resp)
	var bodyObj domain.OrderRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	orderUc := new(mocks.OrderUsecase)
	orderUc.On("Cancel", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, ORDER_CANCEL, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ah := &OrderHandler{orderUc}
	e.POST(ORDER_CANCEL, ah.Cancel)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Status)

	orderUc.AssertExpectations(t)
}

func TestOrderHandler_Amend(t *testing.T) {
	var resp domain.OrderResponse
	faker.FakeData(&resp)
	var bodyObj domain.OrderRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	orderUc := new(mocks.OrderUsecase)
	orderUc.On("Amend", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, ORDER_AMEND, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ah := &OrderHandler{orderUc}
	e.POST(ORDER_AMEND, ah.Amend)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Status)

	orderUc.AssertExpectations(t)
}

func TestOrderHandler_GetDetail(t *testing.T) {
	var resp domain.OrderResponse
	faker.FakeData(&resp)
	var bodyObj domain.OrderRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	orderUc := new(mocks.OrderUsecase)
	orderUc.On("GetDetail", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, ORDER_DETAIL, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ah := &OrderHandler{orderUc}
	e.POST(ORDER_DETAIL, ah.GetDetail)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Status)

	orderUc.AssertExpectations(t)
}

func TestOrderHandler_Create(t *testing.T) {
	var resp domain.OrderResponse
	faker.FakeData(&resp)
	var bodyObj domain.OrderRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	orderUc := new(mocks.OrderUsecase)
	orderUc.On("Create", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, ORDER_CREATE, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ah := &OrderHandler{orderUc}
	e.POST(ORDER_CREATE, ah.Create)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Status)

	orderUc.AssertExpectations(t)
}
