package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Config struct {
	}
)

func RegisterHandlers(cfg Config) (http.Handler, error) {
	r := gin.Default()
	r.GET("/rates", wss)
	r.Run()
	return nil, nil
}

func wss(ctx *gin.Context) {

}
