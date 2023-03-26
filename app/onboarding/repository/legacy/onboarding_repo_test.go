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

func TestOnboardingLegacyRepo_AtmGetInfo(t *testing.T) {
	// Create sample request
	req := &domain.BaseRequest{}

	// Define mock response
	mockResp := &domain.BaseResponse{Status: "success"}
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
	result, err := repo.AtmGetInfo(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
