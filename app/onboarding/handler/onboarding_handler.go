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
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.SimobyVerifyUser(ctx, &req)
}

func (ob *OnboardingHandler) VerifyPhone(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.SimobyVerifyPhone(ctx, &req)
}

func (ob *OnboardingHandler) VerifyOTP(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.SimobyOTP(ctx, &req)
}

func (ob *OnboardingHandler) Register(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.SimobyRegister(ctx, &req)
}

func (ob *OnboardingHandler) Atm(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.AtmGetInfo(ctx, &req)
}

func (ob *OnboardingHandler) AtmOTP(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.AtmOTP(ctx, &req)
}

func (ob *OnboardingHandler) AtmPIN(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.AtmPIN(ctx, &req)
}

func (ob *OnboardingHandler) AtmRegister(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.AtmRegister(ctx, &req)
}

func (ob *OnboardingHandler) NoAtmVerifyUser(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.NoAtmVerifyUser(ctx, &req)
}

func (ob *OnboardingHandler) NoAtmOTP(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.NoAtmOTP(ctx, &req)
}

func (ob *OnboardingHandler) NoAtmRegister(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.NoAtmRegister(ctx, &req)
}

func (ob *OnboardingHandler) NoAtmEmail(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.NoAtmEmail(ctx, &req)
}

func (ob *OnboardingHandler) FreshVerifyPhone(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.FreshVerifyPhone(ctx, &req)
}

func (ob *OnboardingHandler) FreshOTP(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.FreshOTP(ctx, &req)
}

func (ob *OnboardingHandler) FreshPassword(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.FreshPassword(ctx, &req)
}

func (ob *OnboardingHandler) ResetVerifyPhone(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.ResetVerifyPhone(ctx, &req)
}

func (ob *OnboardingHandler) ResetOTP(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.ResetOTP(ctx, &req)
}

func (ob *OnboardingHandler) ResetPassword(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.SimobyVerifyUser(ctx, &req)
}

func (ob *OnboardingHandler) VerifyPassword(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ob.ObUsecase.LoginVerifyPassword(ctx, &req)
}
