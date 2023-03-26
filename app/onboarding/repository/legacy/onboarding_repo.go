package legacy

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"web_server/domain"
)

type onboardingLegacyRepo struct {
	client  *http.Client
	baseURL string
}

func NewOnboardingRepository(client *http.Client, baseURL string) domain.OnboardingRepository {
	return &onboardingLegacyRepo{client, baseURL}
}

// AtmGetInfo implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) AtmGetInfo(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ATM_OTP, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// AtmOTP implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) AtmOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ATM_OTP, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// AtmPIN implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) AtmPIN(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ATM_PIN, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// AtmRegister implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) AtmRegister(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+ATM_REGISTER, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// FreshOTP implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) FreshOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+FRESH_OTP, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// FreshPassword implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) FreshPassword(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+FRESH_PASSWORD, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// FreshVerifyPhone implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) FreshVerifyPhone(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+FRESH_VERIFY_PHONE, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// LoginVerifyPassword implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) LoginVerifyPassword(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+LOGIN_VERIFY_PASSWORD, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// NoAtmEmail implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) NoAtmEmail(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+NO_ATM_EMAIL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// NoAtmOTP implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) NoAtmOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+NO_ATM_OTP, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// NoAtmRegister implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) NoAtmRegister(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+NO_ATM_REGISTER, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// NoAtmVerifyUser implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) NoAtmVerifyUser(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+NO_ATM_VERIFY_USER, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ResetOTP implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) ResetOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+RESET_OTP, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ResetPassword implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) ResetPassword(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+RESET_PASSWORD, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ResetVerifyPhone implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) ResetVerifyPhone(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+RESET_VERIFY_PHONE, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// SimobyOTP implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) SimobyOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+SIMOBI_OTP, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// SimobyRegister implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) SimobyRegister(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+SIMOBI_REGISTER, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// SimobyVerifyPhone implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) SimobyVerifyPhone(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+SIMOBI_VERIFY_PHONE, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// SimobyVerifyUser implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) SimobyVerifyUser(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, o.baseURL+SIMOBI_VERIFY_USER, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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

	var result domain.BaseResponse
	err = json.Unmarshal(respByte, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
