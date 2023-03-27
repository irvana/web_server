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
	faker "github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOnboardingHandler_Verifyuser(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("SimobiVerifyUser", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, SIMOBI_VERIFY_USER, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(SIMOBI_VERIFY_USER, ob.SimobiVerifyUser)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_SimobiVerifyPhone(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("SimobiVerifyPhone", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, SIMOBI_VERIFY_PHONE, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(SIMOBI_VERIFY_PHONE, ob.SimobiVerifyPhone)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_SimobiVerifyOTP(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("SimobiOTP", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, SIMOBI_OTP, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(SIMOBI_OTP, ob.SimobiVerifyOTP)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_SimobiRegister(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("SimobiRegister", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, SIMOBI_REGISTER, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(SIMOBI_REGISTER, ob.SimobiRegister)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}

func TestOnboardingHandler_AtmGetInfo(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("AtmGetInfo", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, ATM_CARD, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(ATM_CARD, ob.AtmGetInfo)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_AtmOTP(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("AtmOTP", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, ATM_OTP, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(ATM_OTP, ob.AtmOTP)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_AtmPIN(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("AtmPIN", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, ATM_PIN, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(ATM_PIN, ob.AtmPIN)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_AtmRegister(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("AtmRegister", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, ATM_REGISTER, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(ATM_REGISTER, ob.AtmRegister)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_NoAtmVerifyUser(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("NoAtmVerifyUser", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, NO_ATM_VERIFY_USER, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(NO_ATM_VERIFY_USER, ob.NoAtmVerifyUser)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_NoAtmOTP(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("NoAtmOTP", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, NO_ATM_OTP, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(NO_ATM_OTP, ob.NoAtmOTP)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_NoAtmEmail(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("NoAtmEmail", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, NO_ATM_EMAIL, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(NO_ATM_EMAIL, ob.NoAtmEmail)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_NoAtmRegister(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("NoAtmRegister", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, NO_ATM_REGISTER, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(NO_ATM_REGISTER, ob.NoAtmRegister)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_FreshVerifyPhone(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("FreshVerifyPhone", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, FRESH_VERIFY_PHONE, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(FRESH_VERIFY_PHONE, ob.FreshVerifyPhone)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_FreshOTP(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("FreshOTP", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, FRESH_OTP, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(FRESH_OTP, ob.FreshOTP)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_FreshPassword(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("FreshPassword", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, FRESH_PASSWORD, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(FRESH_PASSWORD, ob.FreshPassword)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_ResetVerifyPhone(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("ResetVerifyPhone", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, RESET_VERIFY_PHONE, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(RESET_VERIFY_PHONE, ob.ResetVerifyPhone)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_ResetOTP(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("ResetOTP", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, RESET_OTP, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(RESET_OTP, ob.ResetOTP)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_ResetPassword(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("ResetPassword", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, RESET_PASSWORD, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(RESET_PASSWORD, ob.ResetPassword)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
func TestOnboardingHandler_VerifyPassword(t *testing.T) {
	var resp domain.BaseResponse
	faker.FakeData(&resp)
	var bodyObj domain.BaseRequest
	faker.FakeData(&bodyObj)

	body, err := json.Marshal(bodyObj)
	assert.NoError(t, err)

	onboardingUc := new(mocks.OnboardingUsecase)
	onboardingUc.On("LoginVerifyPassword", mock.Anything, &bodyObj).Return(&resp, nil)

	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, LOGIN_VERIFY_PASSWORD, bytes.NewBuffer(body))
	assert.NoError(t, err)

	e := gin.Default()
	ob := &OnboardingHandler{onboardingUc}
	e.POST(LOGIN_VERIFY_PASSWORD, ob.VerifyPassword)

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	respBody, err := io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(respBody), resp.Mail)

	onboardingUc.AssertExpectations(t)
}
