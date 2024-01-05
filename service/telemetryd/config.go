// SPDX-License-Identifier: BSD-3-Clause

package telemetryd

import (
	"github.com/go-logr/logr"
	"github.com/google/uuid"
)

type config struct {
	name    string
	id      uuid.UUID
	log     logr.Logger
	tracing bool
	metrics bool
}

type Option interface {
	apply(*config)
}

type nameOption struct {
	name string
}

func (o *nameOption) apply(c *config) {
	c.name = o.name
}

func WithName(name string) Option {
	return &nameOption{
		name: name,
	}
}

type idOption struct {
	id uuid.UUID
}

func (o *idOption) apply(c *config) {
	c.id = o.id
}

func WithInit(id uuid.UUID) Option {
	return &idOption{
		id: id,
	}
}

type logOption struct {
	log logr.Logger
}

func (o *logOption) apply(c *config) {
	c.log = o.log
}

func WithLogger(log logr.Logger) Option {
	return &logOption{
		log: log,
	}
}

type traceProviderOption struct {
	tracing bool
}

func (o *traceProviderOption) apply(c *config) {
	c.tracing = o.tracing
}

func WithTracing(tracing bool) Option {
	return &traceProviderOption{
		tracing: tracing,
	}
}

type metricProviderOption struct {
	metrics bool
}

func (o *metricProviderOption) apply(c *config) {
	c.metrics = o.metrics
}

func WithMetrics(metrics bool) Option {
	return &metricProviderOption{
		metrics: metrics,
	}
}
