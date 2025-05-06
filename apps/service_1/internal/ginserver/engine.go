package ginserver

import (
	"github.com/gin-gonic/gin"
)

func NewEngine() *gin.Engine {
	e := gin.Default()
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
