package legacy

import (
	"context"
	"net/http"
	"web_server/domain"
)

type onboardingLegacyRepo struct {
	client *http.Client
}

func NewOnboardingRepository(client *http.Client) domain.OnboardingRepository {
	return &onboardingLegacyRepo{client}
}

// AtmGetInfo implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) AtmGetInfo(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// AtmOTP implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) AtmOTP(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// AtmPIN implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) AtmPIN(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// AtmRegister implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) AtmRegister(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// FreshOTP implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) FreshOTP(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// FreshPassword implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) FreshPassword(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// FreshVerifyPhone implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) FreshVerifyPhone(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// LoginVerifyPassword implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) LoginVerifyPassword(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// NoAtmEmail implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) NoAtmEmail(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// NoAtmOTP implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) NoAtmOTP(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// NoAtmRegister implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) NoAtmRegister(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// NoAtmVerifyUser implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) NoAtmVerifyUser(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// ResetOTP implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) ResetOTP(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// ResetPassword implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) ResetPassword(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// ResetVerifyPhone implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) ResetVerifyPhone(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// SimobyOTP implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) SimobyOTP(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// SimobyRegister implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) SimobyRegister(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// SimobyVerifyPhone implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) SimobyVerifyPhone(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}

// SimobyVerifyUser implements domain.OnboardingRepository
func (o *onboardingLegacyRepo) SimobyVerifyUser(ctx context.Context, req *domain.BaseRequest) (domain.BaseResponse, error) {
	panic("unimplemented")
}
