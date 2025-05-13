package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/agoda-com/opentelemetry-go/otelzap"
	"go.uber.org/zap"
	"service_1/internal/apis"
	"service_1/internal/helpers/confighelper"
	"service_1/internal/helpers/loghelper"
	"service_1/internal/helpers/otelhelper"
)

const httpSrvReadHeaderTimeout = 5 * time.Second

func main() {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelSetupper := otelhelper.NewOtelSetupper()
	err := otelSetupper.Setup(ctx)
	if err != nil {
		otelzap.Ctx(ctx).With(zap.Error(err)).Panic("failed to set up OpenTelemetry")
	}
	defer func() {
		errShut := otelSetupper.Shutdown(ctx)
		if errShut != nil {
			otelzap.Ctx(ctx).With(zap.Error(errShut)).Error("failed to shut down OpenTelemetry")
		}
	}()

	// Set up configs.
	cfg, err := confighelper.NewConfigs()
	if err != nil {
		otelzap.Ctx(ctx).With(zap.Error(err)).Panic("failed to load configs")
	}

	// Set up logger.
	err = loghelper.Setup(cfg, otelSetupper.LoggerProvider, "apis")
	if err != nil {
		otelzap.Ctx(ctx).With(zap.Error(err)).Panic("failed to set up logger")
	}

	// Set up APIs app.
	apisAppSetupper := apis.NewAppSetupper()
	err = apisAppSetupper.Setup(cfg)
	if err != nil {
		otelzap.Ctx(ctx).With(zap.Error(err)).Panic("failed to create APIs app")
	}
	defer func() {
		errShut := apisAppSetupper.Shutdown()
		if errShut != nil {
			otelzap.Ctx(ctx).With(zap.Error(errShut)).Error("failed to shut down APIs app")
		}
	}()

	// Start HTTP server.
	httpSrv := &http.Server{
		Addr:              cfg.HTTPServerAddress,
		ReadHeaderTimeout: httpSrvReadHeaderTimeout,
		BaseContext:       func(_ net.Listener) context.Context { return ctx },
		Handler:           apisAppSetupper.App.Handler(),
	}
	errHTTPSrv := make(chan error, 1)
	go func() {
		errHTTPSrv <- httpSrv.ListenAndServe()
	}()

	otelzap.Ctx(ctx).Info("APIs server is up and running")

	select {
	case err = <-errHTTPSrv:
		otelzap.Ctx(ctx).With(zap.Error(err)).Panic("HTTP server exited unexpectedly")
	case <-ctx.Done():
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = httpSrv.Shutdown(context.Background())
	if err != nil {
		otelzap.Ctx(ctx).With(zap.Error(err)).Error("failed to shut down HTTP server")
	}
}
