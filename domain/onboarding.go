package domain

import "context"

// Generated by https://quicktype.io

type VerifyPassword struct {
	BaseResponse
	UserInfo UserInfo `json:"userInfo,omitempty"`
}

type UserInfo struct {
	ID    string `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Name  string `json:"name,omitempty"`
	Ph    string `json:"ph,omitempty"`
	Email string `json:"email,omitempty"`
	Ktp   string `json:"ktp,omitempty"`
	Dob   string `json:"dob,omitempty"`
	Cif   string `json:"cif,omitempty"`
	Reset string `json:"reset,omitempty"`
}

type BaseResponse struct {
	Status      string `json:"status"`
	Ph          string `json:"ph,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	Mail        string `json:"mail,omitempty"`
	IsUserExist string `json:"isUserExist,omitempty"`
}

type BaseRequest struct {
	Ph          string `json:"ph,omitempty"`
	Otp         string `json:"otp,omitempty"`
	Cif         string `json:"cif,omitempty"`
	Pwd         string `json:"pwd,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	ATM         string `json:"atm,omitempty"`
	Pin         string `json:"pin,omitempty"`
	Acc         string `json:"acc,omitempty"`
	Ktp         string `json:"ktp,omitempty"`
	Dob         string `json:"dob,omitempty"`
	Mail        string `json:"mail,omitempty"`
	Token       string `json:"token,omitempty"`
	Name        string `json:"name,omitempty"`
	AccountType string `json:"accountType,omitempty"`
	Currency    string `json:"ccy,omitempty"`
}

type (
	OnboardingUsecase interface {
		SimobiVerifyUser(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		SimobiVerifyPhone(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		SimobiOTP(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		SimobiRegister(ctx context.Context, req *BaseRequest) (*BaseResponse, error)

		AtmGetInfo(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		AtmOTP(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		AtmPIN(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		AtmRegister(ctx context.Context, req *BaseRequest) (*BaseResponse, error)

		NoAtmVerifyUser(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		NoAtmOTP(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		NoAtmEmail(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		NoAtmRegister(ctx context.Context, req *BaseRequest) (*BaseResponse, error)

		FreshVerifyPhone(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		FreshOTP(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		FreshPassword(ctx context.Context, req *BaseRequest) (*BaseResponse, error)

		ResetVerifyPhone(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		ResetOTP(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		ResetPassword(ctx context.Context, req *BaseRequest) (*BaseResponse, error)

		LoginVerifyPassword(ctx context.Context, req *BaseRequest) (*VerifyPassword, error)
	}

	OnboardingRepository interface {
		SimobiVerifyUser(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		SimobiVerifyPhone(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		SimobiOTP(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		SimobiRegister(ctx context.Context, req *BaseRequest) (*BaseResponse, error)

		AtmGetInfo(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		AtmOTP(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		AtmPIN(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		AtmRegister(ctx context.Context, req *BaseRequest) (*BaseResponse, error)

		NoAtmVerifyUser(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		NoAtmOTP(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		NoAtmEmail(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		NoAtmRegister(ctx context.Context, req *BaseRequest) (*BaseResponse, error)

		FreshVerifyPhone(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		FreshOTP(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		FreshPassword(ctx context.Context, req *BaseRequest) (*BaseResponse, error)

		ResetVerifyPhone(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		ResetOTP(ctx context.Context, req *BaseRequest) (*BaseResponse, error)
		ResetPassword(ctx context.Context, req *BaseRequest) (*BaseResponse, error)

		LoginVerifyPassword(ctx context.Context, req *BaseRequest) (*VerifyPassword, error)
	}

	SessionRepository interface {
		GetSessionMeta(ctx context.Context, authId string)
	}
)
