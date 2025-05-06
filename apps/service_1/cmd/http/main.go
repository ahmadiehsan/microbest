package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"service_1/api/http"
	"service_1/internal/helpers"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
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

	// Set up logging.
	helpers.SwitchLoggerToHumanReadableMode()

	// Set up server.
	server := http.NewServer()
	server.App.Use(otelgin.Middleware("gin"))

	// Start HTTP server.
	err = server.App.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run server")
	}
}
