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

	g.POST(STATEMENT_LIST, orderHandler.StatementList)
	g.POST(TRANSACTION_DEAL, orderHandler.Deal)
	g.POST(TRANSACTION_DETAIL, orderHandler.TransactionDetail)

}

func (oh *OrderHandler) GetStatus(ctx *gin.Context) {
	var req domain.OrderRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := oh.OrderUsecase.GetStatus(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)

}

func (oh *OrderHandler) Cancel(ctx *gin.Context) {
	var req domain.OrderRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := oh.OrderUsecase.Cancel(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (oh *OrderHandler) Amend(ctx *gin.Context) {
	var req domain.OrderRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := oh.OrderUsecase.Amend(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (oh *OrderHandler) GetDetail(ctx *gin.Context) {
	var req domain.OrderRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := oh.OrderUsecase.GetDetail(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (oh *OrderHandler) Create(ctx *gin.Context) {
	var req domain.OrderRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := oh.OrderUsecase.Create(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (oh *OrderHandler) StatementList(ctx *gin.Context) {
	var req domain.StatementRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := oh.OrderUsecase.GetStatementList(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (oh *OrderHandler) TransactionDetail(ctx *gin.Context) {
	var req domain.TransactionRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := oh.OrderUsecase.GetTransactionDetail(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (oh *OrderHandler) Deal(ctx *gin.Context) {
	var req domain.TransactionRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := oh.OrderUsecase.Deal(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
