package apis

import (
	"errors"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"service_1/internal/helpers"
	"service_1/internal/pb/service2pb"
)

type App struct {
	configs           *helpers.Configs
	engine            *gin.Engine
	service2RpcClient service2pb.EchoClient
}

func NewApp(cfg *helpers.Configs) (*App, func() error, error) {
	var closeFuncs []func() error

	service2RpcConn, err := createRPCConn(cfg.Service2GrpcAddress)
	if err != nil {
		return nil, nil, err
	}
	closeFuncs = append(closeFuncs, service2RpcConn.Close)

	app := &App{
		configs:           cfg,
		engine:            gin.New(),
		service2RpcClient: service2pb.NewEchoClient(service2RpcConn),
	}
	app.setupMiddlewares()
	app.setupRoutes()

	shutdown := func() error {
		var errShut error
		for _, fn := range closeFuncs {
			errShut = errors.Join(errShut, fn())
		}
		closeFuncs = nil
		return errShut
	}

	return app, shutdown, nil
}

func (a *App) Handler() *gin.Engine {
	return a.engine
}

func (a *App) setupRoutes() {
	api := a.engine.Group("/api")
	api.GET("", a.hello)
	api.GET("/external-api-http", a.externalAPIHTTP)
	api.GET("/service-2-ping-http", a.service2PingHTTP)
	api.GET("/service-2-event-http", a.service2EventHTTP)
	api.GET("/service-2-echo-grpc", a.service2EchoGrpc)
}

func (a *App) setupMiddlewares() {
	a.engine.Use(gin.Recovery())
	a.engine.Use(logger.SetLogger())
	a.engine.Use(otelgin.Middleware("gin"))
}

func createRPCConn(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
