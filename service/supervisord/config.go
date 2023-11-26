// SPDX-License-Identifier: BSD-3-Clause

package supervisord

import (
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/u-bmc/operator/service"
)

type config struct {
	name     string
	id       uuid.UUID
	log      logr.Logger
	services map[string]service.Service
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

type servicesOption struct {
	services map[string]service.Service
}

func (o *servicesOption) apply(c *config) {
	c.services = o.services
}

func WithServices(services ...service.Service) Option {
	serviceMap := make(map[string]service.Service)

	for _, s := range services {
		serviceMap[s.Name()] = s
	}

	return &servicesOption{
		services: serviceMap,
	}
}
