package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"service_1/internal/ginserver"
	"service_1/internal/helpers"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func init() {
	http.DefaultTransport = otelhttp.NewTransport(http.DefaultTransport)
}

func main() {
	if err := run(); err != nil {
		log.Fatal().Err(err).Msg("failed to run server")
	}
}

func run() (err error) {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := helpers.SetupOtel(ctx)
	if err != nil {
		return
	}

	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// Start HTTP server.
	srv := &http.Server{
		Addr:        ":8080",
		BaseContext: func(_ net.Listener) context.Context { return ctx },
		Handler:     newHttpHandler(),
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-srvErr:
		return // Error when starting HTTP server
	case <-ctx.Done():
		stop() // Stop receiving signal notifications as soon as possible.
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = srv.Shutdown(context.Background())
	return
}

func newHttpHandler() http.Handler {
	setupModes()
	middlewares := []gin.HandlerFunc{otelgin.Middleware("gin")}
	return ginserver.NewEngine(middlewares...)
}

func setupModes() {
	configs := helpers.GetConfigs()
	if configs.IsDebug {
		gin.SetMode(gin.DebugMode)
		helpers.SwitchLoggerToHumanReadableMode()
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
