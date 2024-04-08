package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/zero/pkg/log"
)

func NewHttpServer(logger *log.Logger) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	return r
}
