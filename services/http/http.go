package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hugiot/driver-sdk/services/http/middleware"
	"go.uber.org/zap"
	http2 "net/http"
)

func New(l *zap.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.Logger(l), gin.Recovery())
	r.GET("/ping", ping())
	return r
}

// ping handle ping request
func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http2.StatusOK)
	}
}
