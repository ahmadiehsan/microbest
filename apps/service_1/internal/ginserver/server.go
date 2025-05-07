package ginserver

import (
	"errors"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"service_1/internal/helpers"
	"service_1/internal/pb/service2pb"
)

type Server struct {
	configs           *helpers.Configs
	GinEngine         *gin.Engine
	Service2RpcClient service2pb.EchoClient
}

func NewServer(cfg *helpers.Configs) (func() error, *Server) {
	var closeFuncs []func() error

	service2RpcConn := mustCreateRPCConn(cfg.Service2GrpcAddress)
	closeFuncs = append(closeFuncs, service2RpcConn.Close)

	srv := &Server{
		configs:           cfg,
		GinEngine:         gin.New(),
		Service2RpcClient: service2pb.NewEchoClient(service2RpcConn),
	}
	srv.setupMiddlewares()
	srv.setupRoutes()

	shutdown := func() error {
		var shutErr error
		for _, fn := range closeFuncs {
			shutErr = errors.Join(shutErr, fn())
		}
		closeFuncs = nil
		return shutErr
	}

	return shutdown, srv
}

func (s *Server) setupRoutes() {
	api := s.GinEngine.Group("/api")
	api.GET("", s.hello)
	api.GET("/external-api-http", s.externalAPIHTTP)
	api.GET("/service-2-ping-http", s.service2PingHTTP)
	api.GET("/service-2-event-http", s.service2EventHTTP)
	api.GET("/service-2-echo-grpc", s.service2EchoGrpc)
}

func (s *Server) setupMiddlewares() {
	s.GinEngine.Use(gin.Recovery())
	s.GinEngine.Use(logger.SetLogger())
	s.GinEngine.Use(otelgin.Middleware("gin"))
}

func mustCreateRPCConn(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		log.Panic().Err(err).Msgf("could not connect to service address %q", addr)
	}
	return conn
}
