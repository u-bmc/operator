// SPDX-License-Identifier: BSD-3-Clause

package netd

import (
	"context"

	"github.com/google/uuid"
	"github.com/u-bmc/operator/pkg/ipc"
	"github.com/u-bmc/operator/pkg/log"
)

const (
	DefaultName = "netd"
	DefaultUUID = "d8c941f9-c191-404c-ae42-d3c9bf933d0e"
)

type Service struct {
	c config
}

func New(opts ...Option) *Service {
	c := config{
		name:      DefaultName,
		id:        uuid.MustParse(DefaultUUID),
		log:       log.NewDefaultLogger(),
		ipcClient: ipc.NewDefaultClient(),
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
	return nil
}