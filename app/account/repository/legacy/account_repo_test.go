package legacy

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"web_server/domain"

	"github.com/stretchr/testify/assert"
)

func Test_accountLegacyRepo_GetList(t *testing.T) {
	// Create sample request
	req := &domain.BaseRequest{}

	// Define mock response
	mockResp := []domain.AccountResponse{}
	respBody, _ := json.Marshal(mockResp)

	// Create mocked server to return response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))

	defer mockServer.Close()

	repo := &accountLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	result, err := repo.GetList(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}

func Test_accountLegacyRepo_Open(t *testing.T) {
	// Create sample request
	req := &domain.BaseRequest{}

	// Define mock response
	mockResp := &domain.AccountOpenResponse{}
	respBody, _ := json.Marshal(mockResp)

	// Create mocked server to return response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))

	defer mockServer.Close()

	repo := &accountLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	result, err := repo.Open(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
