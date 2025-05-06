package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"service_1/internal/ginserver"
	"service_1/internal/helpers"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := helpers.SetupOtel(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to set up OpenTelemetry")
	}
	defer func() {
		err := otelShutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("failed to shut down OpenTelemetry")
		}
	}()

	// Set up modes.
	setupModes()

	// Set up Gin server.
	ginSrv := ginserver.NewServer()
	defer func() {
		err = ginSrv.Shutdown()
		if err != nil {
			log.Error().Err(err).Msg("failed to shut down Gin server")
		}
	}()

	// Start HTTP server.
	httpSrv := &http.Server{
		Addr:        ":8080",
		BaseContext: func(_ net.Listener) context.Context { return ctx },
		Handler:     ginSrv.GinEngine,
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- httpSrv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-srvErr:
		log.Fatal().Err(err).Msg("failed to run HTTP server")
	case <-ctx.Done():
		stop() // Stop receiving signal notifications as soon as possible.
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = httpSrv.Shutdown(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("failed to shut down HTTP server")
	}
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
