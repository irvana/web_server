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

func Test_OrderRepo_Amend(t *testing.T) {

	// Create sample request
	req := &domain.BaseRequest{}

	// Define mock response
	mockResp := &domain.OrderResponse{Status: "success"}
	respBody, _ := json.Marshal(mockResp)

	// Create mocked server to return response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))
	defer mockServer.Close()

	// Create AtmGetInfo function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.Amend(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)

}

func Test_OrderRepo_Cancel(t *testing.T) {

	// Create sample request
	req := &domain.BaseRequest{}

	// Define mock response
	mockResp := &domain.OrderResponse{Status: "success"}
	respBody, _ := json.Marshal(mockResp)

	// Create mocked server to return response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))
	defer mockServer.Close()

	// Create AtmGetInfo function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.Cancel(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)

}

func Test_OrderRepo_Create(t *testing.T) {

	// Create sample request
	req := &domain.BaseRequest{}

	// Define mock response
	mockResp := &domain.OrderResponse{Status: "success"}
	respBody, _ := json.Marshal(mockResp)

	// Create mocked server to return response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))
	defer mockServer.Close()

	// Create AtmGetInfo function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.Create(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)

}

func Test_OrderRepo_GetDetail(t *testing.T) {

	// Create sample request
	req := &domain.BaseRequest{}

	// Define mock response
	mockResp := &domain.OrderResponse{Status: "success"}
	respBody, _ := json.Marshal(mockResp)

	// Create mocked server to return response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))
	defer mockServer.Close()

	// Create AtmGetInfo function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.GetDetail(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)

}

func Test_OrderRepo_GetStatus(t *testing.T) {

	// Create sample request
	req := &domain.BaseRequest{}

	// Define mock response
	mockResp := &domain.OrderResponse{Status: "success"}
	respBody, _ := json.Marshal(mockResp)

	// Create mocked server to return response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(respBody)
	}))
	defer mockServer.Close()

	// Create AtmGetInfo function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.GetStatus(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)

}
