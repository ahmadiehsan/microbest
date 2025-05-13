package apis

import (
	"errors"
	"time"

	"github.com/agoda-com/opentelemetry-go/otelzap"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"service_1/internal/helpers/confighelper"
	"service_1/internal/pb/service2pb"
)

type AppSetupper struct {
	App        *App
	closeFuncs []func() error
}

func NewAppSetupper() *AppSetupper {
	return &AppSetupper{}
}

func (s *AppSetupper) Setup(cfg *confighelper.Configs) error {
	service2RpcConn, err := createRPCConn(cfg.Service2GrpcAddress)
	if err != nil {
		return err
	}
	s.closeFuncs = append(s.closeFuncs, service2RpcConn.Close)

	app := &App{
		configs:           cfg,
		engine:            gin.New(),
		service2RpcClient: service2pb.NewEchoClient(service2RpcConn),
	}
	app.setupMiddlewares()
	app.setupRoutes()
	app.setupMode()

	s.App = app
	return nil
}

func (s *AppSetupper) Shutdown() error {
	var err error
	for _, fn := range s.closeFuncs {
		err = errors.Join(err, fn())
	}
	s.closeFuncs = nil
	return err
}

type App struct {
	configs           *confighelper.Configs
	engine            *gin.Engine
	service2RpcClient service2pb.EchoClient
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
	a.engine.Use(otelgin.Middleware("gin"))
	a.engine.Use(requestLogger)
}

func (a *App) setupMode() {
	if a.configs.IsDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
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

func requestLogger(ginCtx *gin.Context) {
	ctx := ginCtx.Request.Context()
	startTime := time.Now()

	otelzap.Ctx(ctx).With(
		zap.String("path", ginCtx.Request.URL.Path),
		zap.String("method", ginCtx.Request.Method),
	).Info("req start")

	ginCtx.Next()

	processTime := time.Since(startTime)
	status := ginCtx.Writer.Status()
	otelzap.Ctx(ctx).With(
		zap.String("completed_in", processTime.String()),
		zap.Int("status", status),
	).Info("req end")
}
