package http

import (
	"service_1/internal/helpers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	App     *gin.Engine
	Configs *helpers.Configs
}

func NewServer() *Server {
	server := &Server{
		App:     gin.Default(),
		Configs: helpers.LoadConfigs(),
	}
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	api := s.App.Group("/api")
	api.GET("", s.hello)
	api.GET("/external-api-http", s.externalApiHttp)
	api.GET("/service-2-ping-http", s.service2PingHttp)
	api.GET("/service-2-event-http", s.service2EventHttp)
	api.GET("/service-2-echo-grpc", s.service2EchoGrpc)
}
