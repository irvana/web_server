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

func TestRefHandler_GetCurrency(t *testing.T) {
	var resp []domain.Currency
	faker.FakeData(&resp)
	var bodyObj domain.RefRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	refUc := new(mocks.RefUsecase)
	refUc.On("GetCurrency", mock.Anything, &bodyObj).Return(resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, REF_CCY, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ah := &RefHandler{refUc}
	e.POST(REF_CCY, ah.GetCurrency)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Greater(t, len(string(respBody)), 1)

	refUc.AssertExpectations(t)
}

func TestRefHandler_GetAll(t *testing.T) {
	var resp []domain.RefResponse
	faker.FakeData(&resp)
	var bodyObj domain.RefRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	refUc := new(mocks.RefUsecase)
	refUc.On("GetAll", mock.Anything, &bodyObj).Return(resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, REF_ALL, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ah := &RefHandler{refUc}
	e.POST(REF_ALL, ah.GetAll)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Greater(t, len(string(respBody)), 1)

	refUc.AssertExpectations(t)
}

func TestRefHandler_GetNews(t *testing.T) {
	var resp []domain.News
	faker.FakeData(&resp)
	var bodyObj domain.RefRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	refUc := new(mocks.RefUsecase)
	refUc.On("GetNews", mock.Anything, &bodyObj).Return(resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, REF_NEWS, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ah := &RefHandler{refUc}
	e.POST(REF_NEWS, ah.GetNews)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Greater(t, len(string(respBody)), 10)

	refUc.AssertExpectations(t)
}

func TestRefHandler_GetPair(t *testing.T) {
	var resp []domain.Pair
	faker.FakeData(&resp)
	var bodyObj domain.RefRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	refUc := new(mocks.RefUsecase)
	refUc.On("GetPair", mock.Anything, &bodyObj).Return(resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, REF_PAIR, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ah := &RefHandler{refUc}
	e.POST(REF_PAIR, ah.GetPair)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Greater(t, len(string(respBody)), 10)

	refUc.AssertExpectations(t)
}

func TestRefHandler_GetConfig(t *testing.T) {
	var resp []domain.Config
	faker.FakeData(&resp)
	var bodyObj domain.RefRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	refUc := new(mocks.RefUsecase)
	refUc.On("GetConfig", mock.Anything, &bodyObj).Return(resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, REF_CFG, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ah := &RefHandler{refUc}
	e.POST(REF_CFG, ah.GetConfig)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.NotNil(t, string(respBody))

	refUc.AssertExpectations(t)
}
