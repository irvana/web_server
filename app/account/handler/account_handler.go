package handler

import (
	"net/http"
	"web_server/domain"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	ac domain.AccountUsecase
}

func NewAccountHandler(ac domain.AccountUsecase, g *gin.Engine) {
	acHandler := AccountHandler{ac}

	g.POST(ACCOUNT_LIST, acHandler.GetList)
	g.POST(ACCOUNT_OPEN, acHandler.Open)
}

func (ah *AccountHandler) GetList(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ah.ac.GetList(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (ah *AccountHandler) Open(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := ah.ac.Open(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)

}
