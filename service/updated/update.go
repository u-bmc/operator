// SPDX-License-Identifier: BSD-3-Clause

package updated

import (
	"context"

	"github.com/google/uuid"
	"github.com/u-bmc/operator/pkg/ipc"
	"github.com/u-bmc/operator/pkg/log"
)

const (
	DefaultName = "telemtryd"
	DefaultUUID = "e163a422-d06e-4c78-9f91-cdc060db530b"
)

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

type Service struct {
	c config
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
