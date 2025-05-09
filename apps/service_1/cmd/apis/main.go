package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"service_1/internal/apis"
	"service_1/internal/helpers"
)

const httpSrvReadHeaderTimeout = 5 * time.Second

func main() {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := helpers.SetupOtel(ctx)
	if err != nil {
		log.Panic().Err(err).Msg("failed to set up OpenTelemetry")
	}
	defer func() {
		errShut := otelShutdown(context.Background())
		if errShut != nil {
			log.Error().Err(errShut).Msg("failed to shut down OpenTelemetry")
		}
	}()

	// Set up configs.
	cfg, err := helpers.NewConfigs()
	if err != nil {
		log.Panic().Err(err).Msg("failed to load configs")
	}

	// Set up logger.
	helpers.SetupLogger(cfg, "apis")

	// Set up modes.
	setupModes(cfg)

	// Set up APIs app.
	apisApp, apisAppShutdown, err := apis.NewApp(cfg)
	if err != nil {
		log.Panic().Err(err).Msg("failed to create APIs app")
	}
	defer func() {
		errShut := apisAppShutdown()
		if errShut != nil {
			log.Error().Err(errShut).Msg("failed to shut down APIs app")
		}
	}()

	// Start HTTP server.
	httpSrv := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: httpSrvReadHeaderTimeout,
		BaseContext:       func(_ net.Listener) context.Context { return ctx },
		Handler:           apisApp.Handler(),
	}
	errHTTPSrv := make(chan error, 1)
	go func() {
		errHTTPSrv <- httpSrv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-errHTTPSrv:
		log.Panic().Err(err).Msg("HTTP server exited unexpectedly")
	case <-ctx.Done():
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = httpSrv.Shutdown(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("failed to shut down HTTP server")
	}
}

func setupModes(cfg *helpers.Configs) {
	if cfg.IsDebug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
