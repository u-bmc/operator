// SPDX-License-Identifier: BSD-3-Clause

package ipcd

import (
	"net"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/u-bmc/operator/pkg/ipc"
	"github.com/u-bmc/operator/pkg/log"
)

const (
	DefaultName = "ipcd"
	DefaultUUID = "7d7f58a8-71dd-4e9b-9fb1-fe524f6f9942"
	DefaultAddr = "unix:///run/ipc.sock"
)

type Service struct {
	c config
}

func New(opts ...Option) *Service {
	c := config{
		name:      DefaultName,
		id:        uuid.MustParse(DefaultUUID),
		log:       log.NewDefaultLogger(),
		addr:      DefaultAddr,
		ipcServer: ipc.NewDefaultServer(),
	}

	for _, opt := range opts {
		opt.apply(&c)
	}

	return &Service{
		c: c,
	}
}

func (s *Service) UUID() uuid.UUID {
	return s.c.id
}

func (s *Service) Name() string {
	return s.c.name
}

func (s *Service) Run() error {
	var (
		addr net.Listener
		err  error
	)
	if strings.HasPrefix(s.c.addr, "unix://") {
		path := strings.TrimPrefix(s.c.addr, "unix://")

		if err := os.Remove(path); err != nil {
			return err
		}

		addr, err = net.Listen("unix", path)
		if err != nil {
			return err
		}
	} else {
		addr, err = net.Listen("tcp", s.c.addr)
		if err != nil {
			return err
		}
	}

	return s.c.ipcServer.Serve(addr)
}
