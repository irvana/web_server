package handler

import (
	"web_server/domain"

	"github.com/gin-gonic/gin"
)

type OnboardingHandler struct {
	ObUsecase domain.OnboardingUsecase
}

func NewOnboardingHandler(obUsecase domain.OnboardingUsecase, g *gin.Engine) {
	obHandler := &OnboardingHandler{obUsecase}

	g.POST(SIMOBI_VERIFY_USER, obHandler.Verifyuser)
	g.POST(SIMOBI_VERIFY_PHONE, obHandler.VerifyPhone)
	g.POST(SIMOBI_OTP, obHandler.VerifyOTP)
	g.POST(SIMOBI_REGISTER, obHandler.Register)

	g.POST(ATM_CARD, obHandler.Atm)
	g.POST(ATM_OTP, obHandler.AtmOTP)
	g.POST(ATM_PIN, obHandler.AtmPIN)
	g.POST(ATM_REGISTER, obHandler.AtmRegister)

	g.POST(NO_ATM_VERIFY_USER, obHandler.NoAtmVerifyUser)
	g.POST(NO_ATM_OTP, obHandler.NoAtmOTP)
	g.POST(NO_ATM_EMAIL, obHandler.NoAtmEmail)
	g.POST(NO_ATM_REGISTER, obHandler.NoAtmRegister)

	g.POST(FRESH_VERIFY_PHONE, obHandler.FreshVerifyPhone)
	g.POST(FRESH_OTP, obHandler.FreshOTP)
	g.POST(FRESH_PASSWORD, obHandler.FreshPassword)

	g.POST(RESET_VERIFY_PHONE, obHandler.ResetVerifyPhone)
	g.POST(RESET_OTP, obHandler.ResetOTP)
	g.POST(RESET_PASSWORD, obHandler.ResetPassword)

	g.POST(LOGIN_VERIFY_PASSWORD, obHandler.VerifyPassword)
}

func (ob *OnboardingHandler) Verifyuser(ctx *gin.Context) {

}
func (ob *OnboardingHandler) VerifyPhone(ctx *gin.Context) {

}
func (ob *OnboardingHandler) VerifyOTP(ctx *gin.Context) {

}
func (ob *OnboardingHandler) Register(ctx *gin.Context) {

}
func (ob *OnboardingHandler) Atm(ctx *gin.Context) {

}
func (ob *OnboardingHandler) AtmOTP(ctx *gin.Context) {

}
func (ob *OnboardingHandler) AtmPIN(ctx *gin.Context) {

}
func (ob *OnboardingHandler) AtmRegister(ctx *gin.Context) {

}

func (ob *OnboardingHandler) NoAtmVerifyUser(ctx *gin.Context) {

}
func (ob *OnboardingHandler) NoAtmOTP(ctx *gin.Context) {

}
func (ob *OnboardingHandler) NoAtmRegister(ctx *gin.Context) {

}
func (ob *OnboardingHandler) NoAtmEmail(ctx *gin.Context) {

}

func (ob *OnboardingHandler) FreshVerifyPhone(ctx *gin.Context) {

}
func (ob *OnboardingHandler) FreshOTP(ctx *gin.Context) {

}
func (ob *OnboardingHandler) FreshPassword(ctx *gin.Context) {

}

func (ob *OnboardingHandler) ResetVerifyPhone(ctx *gin.Context) {

}
func (ob *OnboardingHandler) ResetOTP(ctx *gin.Context) {

}
func (ob *OnboardingHandler) ResetPassword(ctx *gin.Context) {

}

func (ob *OnboardingHandler) VerifyPassword(ctx *gin.Context) {

}
