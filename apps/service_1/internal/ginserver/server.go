package ginserver

import (
	"service_1/internal/helpers"
	"service_1/internal/pb/service2pb"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/rs/zerolog/log"
)

type Server struct {
	configs   *helpers.Configs
	GinEngine *gin.Engine

	service2RpcConn   *grpc.ClientConn
	Service2RpcClient service2pb.EchoClient
}

func NewServer(middlewares ...gin.HandlerFunc) *Server {
	configs := helpers.GetConfigs()
	engine := gin.New()

	server := &Server{
		configs:   configs,
		GinEngine: engine,
	}
	server.setupMiddlewares()
	server.setupRoutes()

	server.service2RpcConn = mustCreateRpcConn(configs.Service2GrpcAddress)
	server.Service2RpcClient = service2pb.NewEchoClient(server.service2RpcConn)

	return server
}

func (s *Server) Shutdown() error {
	return s.service2RpcConn.Close()
}

func (s *Server) setupRoutes() {
	api := s.GinEngine.Group("/api")
	api.GET("", s.hello)
	api.GET("/external-api-http", s.externalApiHttp)
	api.GET("/service-2-ping-http", s.service2PingHttp)
	api.GET("/service-2-event-http", s.service2EventHttp)
	api.GET("/service-2-echo-grpc", s.service2EchoGrpc)
}

func (s *Server) setupMiddlewares() {
	s.GinEngine.Use(gin.Recovery())
	s.GinEngine.Use(logger.SetLogger())
	s.GinEngine.Use(otelgin.Middleware("gin"))
}

func mustCreateRpcConn(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("could not connect to service")
	}
	return conn
}
