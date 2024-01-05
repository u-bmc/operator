// SPDX-License-Identifier: BSD-3-Clause

package hardwared

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/u-bmc/operator/pkg/log"
	"github.com/u-bmc/operator/pkg/version"
)

const (
	DefaultName = "hardwared"
	DefaultUUID = "8746603b-2582-44e6-b53a-d5490d3b020a"
)

type Service struct {
	c config
}

func New(opts ...Option) *Service {
	c := config{
		name:    DefaultName,
		id:      uuid.MustParse(DefaultUUID),
		log:     log.NewDefaultLogger(),
		ipcAddr: nats.DefaultURL,
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

	s.c.log.Info("Connecting to ipcd", "service", s.c.name, "uuid", s.c.id.String(), "addr", s.c.ipcAddr)
	var (
		nc  *nats.Conn
		err error
	)
	for {
		nc, err = nats.Connect(s.c.ipcAddr)
		if err != nil {
			if errors.Is(err, nats.ErrNoServers) {
				time.Sleep(time.Second)
				continue
			}
			return err
		}
		break
	}

	srv, err := micro.AddService(nc, micro.Config{
		Name:        s.c.name,
		Version:     version.SemVer,
		Description: "Handles hardware interaction",
	})
	if err != nil {
		return err
	}

	root := srv.AddGroup("hardware")

	if err := root.AddEndpoint("stub", micro.HandlerFunc(s.handleStub)); err != nil {
		return err
	}

	<-ctx.Done()

	return srv.Stop()
}

func (s *Service) handleStub(req micro.Request) {}
