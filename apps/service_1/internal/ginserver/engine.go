package ginserver

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

func NewEngine(middlewares ...gin.HandlerFunc) *gin.Engine {
	e := gin.New()
	setupMiddlewares(e, middlewares...)
	setupRoutes(e)
	return e
}

func setupRoutes(e *gin.Engine) {
	api := e.Group("/api")
	api.GET("", hello)
	api.GET("/external-api-http", externalApiHttp)
	api.GET("/service-2-ping-http", service2PingHttp)
	api.GET("/service-2-event-http", service2EventHttp)
	api.GET("/service-2-echo-grpc", service2EchoGrpc)
}

func setupMiddlewares(e *gin.Engine, middlewares ...gin.HandlerFunc) {
	e.Use(gin.Recovery())
	e.Use(logger.SetLogger())
	e.Use(middlewares...)
}
