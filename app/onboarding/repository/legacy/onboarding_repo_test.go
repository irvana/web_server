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

func TestOnboardingLegacyRepo_SimobyVerifyUser(t *testing.T) {
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

	// Create SimobyVerifyUser function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.SimobyVerifyUser(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_SimobyVerifyPhone(t *testing.T) {
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

	// Create SimobyVerifyPhone function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.SimobyVerifyPhone(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_SimobyOTP(t *testing.T) {
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

	// Create SimobyOTP function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.SimobyOTP(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_SimobyRegister(t *testing.T) {
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

	// Create SimobyRegister function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.SimobyRegister(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_AtmOTP(t *testing.T) {
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

	// Create AtmOTP function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.AtmOTP(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_AtmPIN(t *testing.T) {
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

	// Create AtmPIN function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.AtmPIN(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_AtmRegister(t *testing.T) {
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

	// Create AtmRegister function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.AtmRegister(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_NoAtmVerifyUser(t *testing.T) {
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

	// Create NoAtmVerifyUser function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.NoAtmVerifyUser(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_NoAtmOTP(t *testing.T) {
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

	// Create NoAtmOTP function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.NoAtmOTP(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_NoAtmEmail(t *testing.T) {
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

	// Create NoAtmEmail function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.NoAtmEmail(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_NoAtmRegister(t *testing.T) {
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

	// Create NoAtmRegister function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.NoAtmRegister(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_FreshVerifyPhone(t *testing.T) {
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

	// Create FreshVerifyPhone function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.FreshVerifyPhone(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_FreshOTP(t *testing.T) {
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

	// Create FreshOTP function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.FreshOTP(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_FreshPassword(t *testing.T) {
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

	// Create FreshPassword function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.FreshPassword(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_ResetVerifyPhone(t *testing.T) {
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

	// Create ResetVerifyPhone function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.ResetVerifyPhone(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_ResetOTP(t *testing.T) {
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

	// Create ResetOTP function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.ResetOTP(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_ResetPassword(t *testing.T) {
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

	// Create ResetPassword function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.ResetPassword(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
func TestOnboardingLegacyRepo_LoginVerifyPassword(t *testing.T) {
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

	// Create LoginVerifyPassword function with mocked HTTP client pointing to our mocked server
	repo := &onboardingLegacyRepo{client: &http.Client{Transport: &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		return url.Parse(mockServer.URL)
	}}}, baseURL: mockServer.URL}

	// Call function and assert result matches expected value
	result, err := repo.LoginVerifyPassword(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, mockResp, result)
}
