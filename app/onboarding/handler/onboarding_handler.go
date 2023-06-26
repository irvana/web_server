package handler

import (
	"net/http"
	"web_server/domain"

	"github.com/gin-gonic/gin"
)

type OnboardingHandler struct {
	ObUsecase domain.OnboardingUsecase
}

func NewOnboardingHandler(obUsecase domain.OnboardingUsecase, g *gin.Engine) {
	obHandler := &OnboardingHandler{obUsecase}

	g.POST(SIMOBI_VERIFY_USER, obHandler.SimobiVerifyUser)
	g.POST(SIMOBI_VERIFY_PHONE, obHandler.SimobiVerifyPhone)
	g.POST(SIMOBI_OTP, obHandler.SimobiVerifyOTP)
	g.POST(SIMOBI_REGISTER, obHandler.SimobiRegister)

	g.POST(ATM_CARD, obHandler.AtmGetInfo)
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

func (ob *OnboardingHandler) SimobiVerifyUser(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.SimobiVerifyUser(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) SimobiVerifyPhone(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.SimobiVerifyPhone(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) SimobiVerifyOTP(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.SimobiOTP(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) SimobiRegister(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.SimobiRegister(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) AtmGetInfo(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.AtmGetInfo(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) AtmOTP(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.AtmOTP(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) AtmPIN(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.AtmPIN(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) AtmRegister(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.AtmRegister(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) NoAtmVerifyUser(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.NoAtmVerifyUser(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) NoAtmOTP(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.NoAtmOTP(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) NoAtmRegister(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.NoAtmRegister(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) NoAtmEmail(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.NoAtmEmail(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) FreshVerifyPhone(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.FreshVerifyPhone(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) FreshOTP(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.FreshOTP(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) FreshPassword(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.FreshPassword(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) ResetVerifyPhone(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.ResetVerifyPhone(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) ResetOTP(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.ResetOTP(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) ResetPassword(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.ResetPassword(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ob *OnboardingHandler) VerifyPassword(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ob.ObUsecase.LoginVerifyPassword(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
