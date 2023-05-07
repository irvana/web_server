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
func (o *onboardingUsecase) AtmGetInfo(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.AtmGetInfo(ctx, req)
}

// AtmOTP implements domain.OnboardingUsecase
func (o *onboardingUsecase) AtmOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.AtmOTP(ctx, req)
}

// AtmPIN implements domain.OnboardingUsecase
func (o *onboardingUsecase) AtmPIN(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.AtmPIN(ctx, req)
}

// AtmRegister implements domain.OnboardingUsecase
func (o *onboardingUsecase) AtmRegister(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.AtmRegister(ctx, req)
}

// FreshOTP implements domain.OnboardingUsecase
func (o *onboardingUsecase) FreshOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.FreshOTP(ctx, req)
}

// FreshPassword implements domain.OnboardingUsecase
func (o *onboardingUsecase) FreshPassword(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.FreshPassword(ctx, req)
}

// FreshVerifyPhone implements domain.OnboardingUsecase
func (o *onboardingUsecase) FreshVerifyPhone(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.FreshVerifyPhone(ctx, req)
}

// LoginVerifyPassword implements domain.OnboardingUsecase
func (o *onboardingUsecase) LoginVerifyPassword(ctx context.Context, req *domain.BaseRequest) (*domain.VerifyPassword, error) {
	// TODO reassign websocket client (change topic or role for connected client)
	return o.obRepo.LoginVerifyPassword(ctx, req)
}

// NoAtmEmail implements domain.OnboardingUsecase
func (o *onboardingUsecase) NoAtmEmail(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.NoAtmEmail(ctx, req)
}

// NoAtmOTP implements domain.OnboardingUsecase
func (o *onboardingUsecase) NoAtmOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.NoAtmOTP(ctx, req)
}

// NoAtmRegister implements domain.OnboardingUsecase
func (o *onboardingUsecase) NoAtmRegister(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.NoAtmRegister(ctx, req)
}

// NoAtmVerifyUser implements domain.OnboardingUsecase
func (o *onboardingUsecase) NoAtmVerifyUser(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.NoAtmVerifyUser(ctx, req)
}

// ResetOTP implements domain.OnboardingUsecase
func (o *onboardingUsecase) ResetOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.ResetOTP(ctx, req)
}

// ResetPassword implements domain.OnboardingUsecase
func (o *onboardingUsecase) ResetPassword(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.ResetPassword(ctx, req)
}

// ResetVerifyPhone implements domain.OnboardingUsecase
func (o *onboardingUsecase) ResetVerifyPhone(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.ResetVerifyPhone(ctx, req)
}

// SimobiOTP implements domain.OnboardingUsecase
func (o *onboardingUsecase) SimobiOTP(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.SimobiOTP(ctx, req)
}

// SimobiRegister implements domain.OnboardingUsecase
func (o *onboardingUsecase) SimobiRegister(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.SimobiRegister(ctx, req)
}

// SimobiVerifyPhone implements domain.OnboardingUsecase
func (o *onboardingUsecase) SimobiVerifyPhone(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.SimobiVerifyPhone(ctx, req)
}

// SimobiVerifyUser implements domain.OnboardingUsecase
func (o *onboardingUsecase) SimobiVerifyUser(ctx context.Context, req *domain.BaseRequest) (*domain.BaseResponse, error) {
	return o.obRepo.SimobiVerifyUser(ctx, req)
}
