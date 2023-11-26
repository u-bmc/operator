// SPDX-License-Identifier: BSD-3-Clause

package netd

import (
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/u-bmc/operator/api/gen/ipc/v1alpha1/ipcv1alpha1connect"
	"github.com/u-bmc/operator/pkg/iface"
)

type config struct {
	name      string
	id        uuid.UUID
	log       logr.Logger
	ipcClient ipcv1alpha1connect.IPCServiceClient
	ifaces    []iface.Iface
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

type ipcClientOption struct {
	ipcClient ipcv1alpha1connect.IPCServiceClient
}

func (o *ipcClientOption) apply(c *config) {
	c.ipcClient = o.ipcClient
}

func WithIPCClient(ipcClient ipcv1alpha1connect.IPCServiceClient) Option {
	return &ipcClientOption{
		ipcClient: ipcClient,
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
