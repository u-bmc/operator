// SPDX-License-Identifier: BSD-3-Clause

package telemetry

import (
	"context"
	"errors"
	"io"

	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func NewResource(serviceName, serviceVersion string) *resource.Resource {
	var res *resource.Resource
	res, err := resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
		))
	if err != nil {
		res = resource.Default()
	}

	return res
}

func NewPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func NewTraceNoopProvider(res *resource.Resource) *trace.TracerProvider {
	traceExporter, err := stdouttrace.New(stdouttrace.WithWriter(io.Discard))
	if err != nil {
		return trace.NewTracerProvider(trace.WithResource(res))
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter),
		trace.WithResource(res),
	)

	return traceProvider
}

func NewTracegRPCProvider(ctx context.Context, res *resource.Resource) *trace.TracerProvider {
	traceExporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return trace.NewTracerProvider(trace.WithResource(res))
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter),
		trace.WithResource(res),
	)

	return traceProvider
}

func NewMeterNoopProvider(res *resource.Resource) *metric.MeterProvider {
	metricExporter, err := stdoutmetric.New(stdoutmetric.WithWriter(io.Discard))
	if err != nil {
		return metric.NewMeterProvider(metric.WithResource(res))
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
	)

	return meterProvider
}

func NewMetergRPCProvider(ctx context.Context, res *resource.Resource) *metric.MeterProvider {
	metricExporter, err := otlpmetricgrpc.New(ctx)
	if err != nil {
		return metric.NewMeterProvider(metric.WithResource(res))
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
	)

	return meterProvider
}

func SetupOtelSDK(ctx context.Context, res *resource.Resource, log logr.Logger) func(context.Context) error {
	var shutdownFuncs []func(ctx context.Context) error

	// shutdown calls cleanup functions registered via shutdownFuncs.
	// The errors from the calls are joined.
	// Each registered cleanup will be invoked once.
	shutdown := func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	// Set up Logger.
	otel.SetLogger(log)

	// Set up propagator.
	otel.SetTextMapPropagator(NewPropagator())

	// Set up trace provider.
	tracerProvider := NewTraceNoopProvider(res)
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	// Set up meter provider.
	meterProvider := NewMeterNoopProvider(res)
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)
	otel.SetMeterProvider(meterProvider)

	return shutdown
}
