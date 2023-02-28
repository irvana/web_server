package handler

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type (
	Config struct {
		Address string
		Port    int
	}
)

var (
	onceHandler sync.Once
	handler     http.Handler
)

func SetupHandlers(cfg Config, wssHandler http.Handler) http.Handler {
	onceHandler.Do(func() {
		r := gin.Default()
		handler = registerHandlers(r, wssHandler)
	})
	return handler
}

func registerHandlers(r *gin.Engine, wss http.Handler) http.Handler {
	r.GET(WEBSOCKET_SERVER, func(ctx *gin.Context) { wss.ServeHTTP(ctx.Writer, ctx.Request) })

	r.GET(SIMOBI_VERIFY_USER, simobiVerifyUser)
	r.GET(SIMOBI_VERIFY_PHONE, simobiVerifyPhone)
	r.GET(SIMOBI_OTP, simobiVerifyOTP)
	r.GET(SIMOBI_REGISTER, simobyVerifyRegister)

	return r
}
