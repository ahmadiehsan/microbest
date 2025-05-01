package http

import (
	"github.com/gin-gonic/gin"
)

func NewGinApp() *gin.Engine {
	app := gin.Default()
	setupRoutes(app)
	return app
}

func setupRoutes(app *gin.Engine) {
	api := app.Group("/api")

	api.GET("", hello)
	api.GET("/external-api-http", externalApiHttp)
	api.GET("/service-2-ping-http", service2PingHttp)
	api.GET("/service-2-event-http", service2EventHttp)
	api.GET("/service-2-echo-grpc", service2EchoGrpc)
}
