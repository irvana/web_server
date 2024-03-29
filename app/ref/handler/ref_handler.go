package handler

import (
	"net/http"
	"web_server/domain"

	"github.com/gin-gonic/gin"
)

type RefHandler struct {
	RefUsecase domain.RefUsecase
}

const (
	REF_ALL  = "/ref/all"
	REF_CCY  = "/ref/ccy"
	REF_PAIR = "/ref/pair"
	REF_NEWS = "/ref/news"
	REF_CFG  = "/ref/cfg"
)

func NewRefHandler(refUsecase domain.RefUsecase, g *gin.Engine) {
	refHandler := &RefHandler{refUsecase}

	g.POST(REF_ALL, refHandler.GetAll)
	g.POST(REF_CCY, refHandler.GetCurrency)
	g.POST(REF_NEWS, refHandler.GetNews)
	g.POST(REF_PAIR, refHandler.GetPair)
	g.POST(REF_CFG, refHandler.GetConfig)
}

func (ref *RefHandler) GetCurrency(ctx *gin.Context) {
	var req domain.RefRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ref.RefUsecase.GetCurrency(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ref *RefHandler) GetAll(ctx *gin.Context) {
	var req domain.RefRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ref.RefUsecase.GetAll(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ref *RefHandler) GetNews(ctx *gin.Context) {
	var req domain.RefRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ref.RefUsecase.GetNews(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ref *RefHandler) GetPair(ctx *gin.Context) {
	var req domain.RefRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ref.RefUsecase.GetPair(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ref *RefHandler) GetConfig(ctx *gin.Context) {
	var req domain.RefRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ref.RefUsecase.GetConfig(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
