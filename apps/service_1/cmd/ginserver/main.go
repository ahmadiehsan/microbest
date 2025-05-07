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
	"service_1/internal/ginserver"
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
	cfg := helpers.NewConfigs()

	// Set up modes.
	setupModes(cfg)

	// Set up Gin server.
	ginShutdown, ginSrv := ginserver.NewServer(cfg)
	defer func() {
		errShut := ginShutdown()
		if errShut != nil {
			log.Error().Err(errShut).Msg("failed to shut down Gin server")
		}
	}()

	// Start HTTP server.
	httpSrv := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: httpSrvReadHeaderTimeout,
		BaseContext:       func(_ net.Listener) context.Context { return ctx },
		Handler:           ginSrv.Handler(),
	}
	errHTTPSrv := make(chan error, 1)
	go func() {
		errHTTPSrv <- httpSrv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-errHTTPSrv:
		log.Panic().Err(err).Msg("failed to run HTTP server")
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
		helpers.SwitchLoggerToHumanReadableMode()
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}
