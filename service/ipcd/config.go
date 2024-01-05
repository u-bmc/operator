// SPDX-License-Identifier: BSD-3-Clause

package ipcd

import (
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/nats-io/nats-server/v2/server"
)

type config struct {
	name string
	id   uuid.UUID
	log  logr.Logger
	so   server.Options
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

type serverOption struct {
	so server.Options
}

func (o *serverOption) apply(c *config) {
	c.so = o.so
}

func WithServer(so server.Options) Option {
	return &serverOption{
		so: so,
	}
}
