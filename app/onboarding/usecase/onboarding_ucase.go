package usecase

import (
	"context"
	"web_server/domain"
)

type onboardingUsecase struct {
	obRepo domain.OnboardingRepository
}

func NewOnboardingUsecase(obRepo domain.OnboardingRepository) domain.OnboardingUsecase {
	return &onboardingUsecase{obRepo}
}

// AtmGetInfo implements domain.OnboardingUsecase
func (*onboardingUsecase) AtmGetInfo(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// AtmOTP implements domain.OnboardingUsecase
func (*onboardingUsecase) AtmOTP(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// AtmPIN implements domain.OnboardingUsecase
func (*onboardingUsecase) AtmPIN(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// AtmRegister implements domain.OnboardingUsecase
func (*onboardingUsecase) AtmRegister(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// FreshOTP implements domain.OnboardingUsecase
func (*onboardingUsecase) FreshOTP(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// FreshPassword implements domain.OnboardingUsecase
func (*onboardingUsecase) FreshPassword(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// FreshVerifyPhone implements domain.OnboardingUsecase
func (*onboardingUsecase) FreshVerifyPhone(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// LoginVerifyPassword implements domain.OnboardingUsecase
func (*onboardingUsecase) LoginVerifyPassword(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// NoAtmEmail implements domain.OnboardingUsecase
func (*onboardingUsecase) NoAtmEmail(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// NoAtmOTP implements domain.OnboardingUsecase
func (*onboardingUsecase) NoAtmOTP(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// NoAtmRegister implements domain.OnboardingUsecase
func (*onboardingUsecase) NoAtmRegister(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// NoAtmVerifyUser implements domain.OnboardingUsecase
func (*onboardingUsecase) NoAtmVerifyUser(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// ResetOTP implements domain.OnboardingUsecase
func (*onboardingUsecase) ResetOTP(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// ResetPassword implements domain.OnboardingUsecase
func (*onboardingUsecase) ResetPassword(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// ResetVerifyPhone implements domain.OnboardingUsecase
func (*onboardingUsecase) ResetVerifyPhone(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// SimobyOTP implements domain.OnboardingUsecase
func (*onboardingUsecase) SimobyOTP(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// SimobyRegister implements domain.OnboardingUsecase
func (*onboardingUsecase) SimobyRegister(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// SimobyVerifyPhone implements domain.OnboardingUsecase
func (*onboardingUsecase) SimobyVerifyPhone(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// SimobyVerifyUser implements domain.OnboardingUsecase
func (*onboardingUsecase) SimobyVerifyUser(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}
