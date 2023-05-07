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

func Test_refLegacyRepo_GetAll(t *testing.T) {

	// Create sample request
	req := &domain.RefRequest{}

	// Define mock response
	mockResp := []domain.RefResponse{}
	respBody, _ := json.Marshal(mockResp)

	// Create mocked server to return response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))
	defer mockServer.Close()

	// Create AtmGetInfo function with mocked HTTP client pointing to our mocked server
	repo := &refLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.GetAll(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}

func Test_refLegacyRepo_GetCurrency(t *testing.T) {
	// Create sample request
	req := &domain.RefRequest{}

	// Define mock response
	mockResp := []domain.Currency{}
	respBody, _ := json.Marshal(mockResp)

	// Create mocked server to return response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))
	defer mockServer.Close()

	// Create AtmGetInfo function with mocked HTTP client pointing to our mocked server
	repo := &refLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.GetCurrency(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}

func Test_refLegacyRepo_GetPair(t *testing.T) {
	// Create sample request
	req := &domain.RefRequest{}

	// Define mock response
	mockResp := []domain.Pair{}
	respBody, _ := json.Marshal(mockResp)

	// Create mocked server to return response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))
	defer mockServer.Close()

	// Create AtmGetInfo function with mocked HTTP client pointing to our mocked server
	repo := &refLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.GetPair(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}

func Test_refLegacyRepo_GetNews(t *testing.T) {
	// Create sample request
	req := &domain.RefRequest{}

	// Define mock response
	mockResp := []domain.News{}
	respBody, _ := json.Marshal(mockResp)

	// Create mocked server to return response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))
	defer mockServer.Close()

	// Create AtmGetInfo function with mocked HTTP client pointing to our mocked server
	repo := &refLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.GetNews(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}

func Test_refLegacyRepo_GetConfig(t *testing.T) {
	// Create sample request
	req := &domain.RefRequest{}

	// Define mock response
	mockResp := &domain.Config{}
	respBody, _ := json.Marshal(mockResp)

	// Create mocked server to return response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))
	defer mockServer.Close()

	// Create AtmGetInfo function with mocked HTTP client pointing to our mocked server
	repo := &refLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.GetConfig(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
