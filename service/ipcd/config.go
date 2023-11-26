// SPDX-License-Identifier: BSD-3-Clause

package ipcd

import (
	"net/http"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
)

type config struct {
	name      string
	id        uuid.UUID
	log       logr.Logger
	addr      string
	ipcServer *http.Server
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

func WithSocket(addr string) Option {
	return &addrOption{
		addr: addr,
	}
}

type ipcServerOption struct {
	ipcServer *http.Server
}

func (o *ipcServerOption) apply(c *config) {
	c.ipcServer = o.ipcServer
}

func WithServer(ipcServer *http.Server) Option {
	return &ipcServerOption{
		ipcServer: ipcServer,
	}
}
