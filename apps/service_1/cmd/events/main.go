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
	cfg, err := helpers.NewConfigs()
	if err != nil {
		log.Panic().Err(err).Msg("failed to load configs")
	}

	// Set up logger.
	helpers.SetupLogger(cfg, "events")

	// Set up events app.
	eventsAppShutdown, eventsApp := events.NewApp(cfg)
	defer func() {
		errShut := eventsAppShutdown()
		if errShut != nil {
			log.Error().Err(errShut).Msg("failed to shut down events app")
		}
	}()

	// Start events server.
	errEventsSrv := make(chan error, 1)
	go func() {
		errEventsSrv <- eventsApp.Listen(ctx)
	}()

	// Wait for interruption.
	select {
	case err = <-errEventsSrv:
		log.Panic().Err(err).Msg("events server exited unexpectedly")
	case <-ctx.Done():
		stop()
	}
}
