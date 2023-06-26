package handler

import (
	"context"
	"net/http"
	"web_server/domain"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type rateHandler struct {
	rateUc domain.RateUsecase
}

func RunRateHandlers(rateUc domain.RateUsecase, g *gin.Engine) {
	rateHandler := &rateHandler{rateUc}

	go rateUc.ProcessBackgroundRate(context.WithValue(context.Background(), "rate", "ref"))
	go rateUc.ProcessBackgroundRef(context.WithValue(context.Background(), "ref", "ref"))

	g.POST("/chart", rateHandler.GetChart)
}

func (ref *rateHandler) GetChart(ctx *gin.Context) {
	var req domain.HistoricRateRequest
	if err := ctx.BindJSON(&req); err != nil {
		logrus.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := ref.rateUc.GetChart(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
