// SPDX-License-Identifier: BSD-3-Clause

package telemetry

import (
	"context"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
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

func NewTraceStdoutProvider(res *resource.Resource) *trace.TracerProvider {
	traceExporter, err := stdouttrace.New()
	if err != nil {
		return trace.NewTracerProvider(trace.WithResource(res))
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter),
		trace.WithResource(res),
	)

	return traceProvider
}

func NewTraceHTTPProvider(ctx context.Context, res *resource.Resource) *trace.TracerProvider {
	traceExporter, err := otlptracehttp.New(ctx)
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

func NewMeterStdoutProvider(res *resource.Resource) *metric.MeterProvider {
	metricExporter, err := stdoutmetric.New()
	if err != nil {
		return metric.NewMeterProvider(metric.WithResource(res))
	}

	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
	)

	return meterProvider
}

func NewMeterHTTPProvider(ctx context.Context, res *resource.Resource) *metric.MeterProvider {
	metricExporter, err := otlpmetrichttp.New(ctx)
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
