// SPDX-License-Identifier: BSD-3-Clause

package supervisord

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/u-bmc/operator/pkg/log"
	"github.com/u-bmc/operator/service"
)

const (
	DefaultName = "supervisord"
	DefaultUUID = "3a491ef2-9ef2-454c-8a76-fa425910504c"
)

type Service struct {
	c config
}

func New(opts ...Option) *Service {
	cfg := config{
		name: DefaultName,
		id:   uuid.MustParse(DefaultUUID),
		log:  log.NewDefaultLogger(),
	}

	for _, opt := range opts {
		opt.apply(&cfg)
	}

	return &Service{
		c: cfg,
	}
}

func (s *Service) UUID() uuid.UUID {
	return s.c.id
}

func (s *Service) Name() string {
	return s.c.name
}

func (s *Service) Run() error {
	s.c.log.Info("Running supervisord service")

	exitCh := make(chan error)
	self, err := filepath.EvalSymlinks("/proc/self/exe")
	if err != nil {
		return err
	}

	for name, serv := range s.c.services {
		if err := os.Link(self, filepath.Join("/run", name)); err != nil {
			return err
		}

		go func(serv service.Service, c chan error) {
			defer func(serv service.Service, c chan error) {
				if r := recover(); r != nil {
					err := fmt.Errorf("%v", r)
					s.c.log.Error(err, "Panic occurred in service and was recovered", "service", serv.Name())
					c <- err
				}
			}(serv, c)

			if err := serv.Run(); err != nil {
				s.c.log.Error(err, "Error occurred during service runtime", "service", serv.Name())
				exitCh <- err
			} else {
				s.c.log.Info("Service exited successfully", "service", serv.Name())
				exitCh <- nil
			}
		}(serv, exitCh)
	}

	return nil
}
