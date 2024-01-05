// SPDX-License-Identifier: BSD-3-Clause

package netd

import (
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/u-bmc/operator/pkg/iface"
)

type config struct {
	name    string
	id      uuid.UUID
	log     logr.Logger
	ipcAddr string
	ifaces  []iface.Iface
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

type ipcAddrOption struct {
	ipcAddr string
}

func (o *ipcAddrOption) apply(c *config) {
	c.ipcAddr = o.ipcAddr
}

func WithIPCAddr(ipcAddr string) Option {
	return &ipcAddrOption{
		ipcAddr: ipcAddr,
	}
}

type ifacesOption struct {
	ifaces []iface.Iface
}

func (o *ifacesOption) apply(c *config) {
	c.ifaces = o.ifaces
}

func WithIfaces(ifaces []iface.Iface) Option {
	return &ifacesOption{
		ifaces: ifaces,
	}
}
