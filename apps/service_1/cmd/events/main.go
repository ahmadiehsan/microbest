package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/agoda-com/opentelemetry-go/otelzap"
	"go.uber.org/zap"
	"service_1/internal/events"
	"service_1/internal/helpers/confighelper"
	"service_1/internal/helpers/loghelper"
	"service_1/internal/helpers/otelhelper"
)

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
	err = loghelper.Setup(cfg, otelSetupper.LoggerProvider, "events")
	if err != nil {
		otelzap.Ctx(ctx).With(zap.Error(err)).Panic("failed to set up logger")
	}

	// Set up events app.
	eventsAppSetupper := events.NewAppSetupper()
	err = eventsAppSetupper.Setup(cfg)
	if err != nil {
		otelzap.Ctx(ctx).With(zap.Error(err)).Panic("failed to create events app")
	}
	defer func() {
		errShut := eventsAppSetupper.Shutdown()
		if errShut != nil {
			otelzap.Ctx(ctx).With(zap.Error(errShut)).Error("failed to shut down events app")
		}
	}()

	// Start events server.
	errEventsSrv := make(chan error, 1)
	go func() {
		errEventsSrv <- eventsAppSetupper.App.Listen(ctx)
	}()

	otelzap.Ctx(ctx).Info("events server is up and running")

	select {
	case err = <-errEventsSrv:
		otelzap.Ctx(ctx).With(zap.Error(err)).Panic("events server exited unexpectedly")
	case <-ctx.Done():
		stop()
	}
}
