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

func TestAccountHandler_GetList(t *testing.T) {
	var resp []domain.AccountResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	accountUc := new(mocks.AccountUsecase)
	accountUc.On("GetList", mock.Anything, &bodyObj).Return(resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, ACCOUNT_LIST, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ah := &AccountHandler{accountUc}
	e.POST(ACCOUNT_LIST, ah.GetList)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Greater(t, len(string(respBody)), 10)

	accountUc.AssertExpectations(t)
}

func TestAccountHandler_Open(t *testing.T) {
	var resp domain.AccountOpenResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	accountUc := new(mocks.AccountUsecase)
	accountUc.On("Open", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, ACCOUNT_OPEN, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ah := &AccountHandler{accountUc}
	e.POST(ACCOUNT_OPEN, ah.Open)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Status)

	accountUc.AssertExpectations(t)
}
