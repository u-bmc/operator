// SPDX-License-Identifier: BSD-3-Clause

package ipcd

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/u-bmc/operator/pkg/log"
)

const (
	DefaultName     = "ipcd"
	DefaultUUID     = "7d7f58a8-71dd-4e9b-9fb1-fe524f6f9942"
	DefaultHost     = "localhost"
	DefaultStoreDir = "/run/ipcd/storage"
)

type Service struct {
	c config
}

func New(opts ...Option) *Service {
	c := config{
		name: DefaultName,
		id:   uuid.MustParse(DefaultUUID),
		log:  log.NewDefaultLogger(),
		so: server.Options{
			ServerName: fmt.Sprintf("%s-%s", DefaultName, DefaultUUID),
			Host:       DefaultHost,
			Port:       server.DEFAULT_PORT,
			StoreDir:   DefaultStoreDir,
		},
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

func (s *Service) Run(ctx context.Context) error {
	s.c.log.Info("Starting service", "service", s.c.name, "uuid", s.c.id.String())

	s.c.log.Info("Creating IPC server", "addr", s.c.so.Host, "port", s.c.so.Port)
	ns, err := server.NewServer(&s.c.so)
	if err != nil {
		s.c.log.Error(err, "Failed to create IPC server")
		return err
	}

	ns.SetLoggerV2(log.NewNATSLogger(s.c.log), s.c.so.Debug, s.c.so.Trace, s.c.so.TraceVerbose)

	s.c.log.Info("Starting IPC server", "addr", s.c.so.Host, "port", s.c.so.Port)

	go ns.Start()
	defer ns.Shutdown()

	<-ctx.Done()

	return nil
}
