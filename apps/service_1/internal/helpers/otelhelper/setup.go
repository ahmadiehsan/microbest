package otelhelper

import (
	"context"
	"errors"

	otellogs "github.com/agoda-com/opentelemetry-logs-go"
	"github.com/agoda-com/opentelemetry-logs-go/exporters/otlp/otlplogs"
	"github.com/agoda-com/opentelemetry-logs-go/exporters/otlp/otlplogs/otlplogsgrpc"
	"github.com/agoda-com/opentelemetry-logs-go/sdk/logs"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

type OtelSetupper struct {
	LoggerProvider *logs.LoggerProvider
	shutdownFuncs  []func(context.Context) error
}

func NewOtelSetupper() *OtelSetupper {
	return &OtelSetupper{}
}

func (s *OtelSetupper) Setup(ctx context.Context) error {
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	loggerProvider, err := newLoggerProvider(ctx)
	if err != nil {
		return err
	}
	s.LoggerProvider = loggerProvider
	s.shutdownFuncs = append(s.shutdownFuncs, loggerProvider.Shutdown)
	otellogs.SetLoggerProvider(loggerProvider)

	tracerProvider, err := newTracerProvider(ctx)
	if err != nil {
		return err
	}
	s.shutdownFuncs = append(s.shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	meterProvider, err := newMeterProvider(ctx)
	if err != nil {
		return err
	}
	s.shutdownFuncs = append(s.shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	return nil
}

func (s *OtelSetupper) Shutdown(ctx context.Context) error {
	var err error
	for _, fn := range s.shutdownFuncs {
		err = errors.Join(err, fn(ctx))
	}
	s.shutdownFuncs = nil
	return err
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newLoggerProvider(ctx context.Context) (*logs.LoggerProvider, error) {
	exporter, err := otlplogs.NewExporter(ctx, otlplogs.WithClient(otlplogsgrpc.NewClient()))
	if err != nil {
		return nil, err
	}

	provider := logs.NewLoggerProvider(logs.WithBatcher(exporter))
	return provider, nil
}

func newTracerProvider(ctx context.Context) (*trace.TracerProvider, error) {
	exporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return nil, err
	}

	provider := trace.NewTracerProvider(trace.WithBatcher(exporter))
	return provider, nil
}

func newMeterProvider(ctx context.Context) (*metric.MeterProvider, error) {
	exporter, err := otlpmetricgrpc.New(ctx)
	if err != nil {
		return nil, err
	}

	reader := metric.NewPeriodicReader(exporter)
	provider := metric.NewMeterProvider(metric.WithReader(reader))
	return provider, nil
}
