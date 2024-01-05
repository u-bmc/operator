// SPDX-License-Identifier: BSD-3-Clause

package registryd

import (
	"github.com/go-logr/logr"
	"github.com/google/uuid"
)

type config struct {
	name    string
	id      uuid.UUID
	log     logr.Logger
	ipcAddr string
	dbPath  string
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

type dbPathOption struct {
	dbPath string
}

func (o *dbPathOption) apply(c *config) {
	c.dbPath = o.dbPath
}

func WithDBPath(dbPath string) Option {
	return &dbPathOption{
		dbPath: dbPath,
	}
}
