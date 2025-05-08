package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
	"service_1/internal/events"
	"service_1/internal/helpers"
)

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

	// Set up events listener.
	eventsShutdown, eventsSrv := events.NewServer(cfg)
	defer func() {
		errShut := eventsShutdown()
		if errShut != nil {
			log.Error().Err(errShut).Msg("failed to shut down events server")
		}
	}()

	// Start events server.
	errEventsSrv := make(chan error, 1)
	go func() {
		errEventsSrv <- eventsSrv.Listen(ctx)
	}()

	// Wait for interruption.
	select {
	case err = <-errEventsSrv:
		log.Panic().Err(err).Msg("failed to run events server")
	case <-ctx.Done():
		stop()
	}
}

func setupModes(cfg *helpers.Configs) {
	if cfg.IsDebug {
		helpers.SwitchLoggerToHumanReadableMode()
	}
}
