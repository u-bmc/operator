// SPDX-License-Identifier: BSD-3-Clause

package ipcd

import (
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/u-bmc/operator/pkg/ipc"
)

type config struct {
	name     string
	id       uuid.UUID
	log      logr.Logger
	addr     string
	addrType ipc.Transport
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

type addrOption struct {
	addr string
}

func (o *addrOption) apply(c *config) {
	c.addr = o.addr
}

func WithAddr(addr string) Option {
	return &addrOption{
		addr: addr,
	}
}

type addrTypeOption struct {
	addrType ipc.Transport
}

func (o *addrTypeOption) apply(c *config) {
	c.addrType = o.addrType
}

func WithAddrType(addrType ipc.Transport) Option {
	return &addrTypeOption{
		addrType: addrType,
	}
}
