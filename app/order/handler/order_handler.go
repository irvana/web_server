package handler

import (
	"net/http"
	"web_server/domain"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderUsecase domain.OrderUsecase
}

func NewOrderHandler(orderUc domain.OrderUsecase, g *gin.Engine) {
	orderHandler := OrderHandler{orderUc}

	g.POST(ORDER_STATUS, orderHandler.GetStatus)
	g.POST(ORDER_CANCEL, orderHandler.Cancel)
	g.POST(ORDER_AMEND, orderHandler.Amend)
	g.POST(ORDER_DETAIL, orderHandler.GetDetail)
	g.POST(ORDER_CREATE, orderHandler.Create)

}

func (oh *OrderHandler) GetStatus(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := oh.OrderUsecase.GetStatus(ctx, &req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	ctx.JSON(http.StatusOK, res)

}

func (oh *OrderHandler) Cancel(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := oh.OrderUsecase.Cancel(ctx, &req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	ctx.JSON(http.StatusOK, res)
}

func (oh *OrderHandler) Amend(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := oh.OrderUsecase.Amend(ctx, &req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	ctx.JSON(http.StatusOK, res)
}

func (oh *OrderHandler) GetDetail(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := oh.OrderUsecase.GetDetail(ctx, &req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	ctx.JSON(http.StatusOK, res)
}

func (oh *OrderHandler) Create(ctx *gin.Context) {
	var req domain.BaseRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := oh.OrderUsecase.Create(ctx, &req)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	ctx.JSON(http.StatusOK, res)
}
