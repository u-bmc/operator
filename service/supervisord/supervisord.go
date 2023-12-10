// SPDX-License-Identifier: BSD-3-Clause

package supervisord

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"

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
	s.c.log.Info("Starting service", "service", s.c.name, "uuid", s.c.id.String())

	// Find the path to the current executable, ideally it is always /sbin/operator
	// but this approach is more robust
	// self, err := filepath.EvalSymlinks("/proc/self/exe")
	// if err != nil {
	// 	return err
	// }

	var wg sync.WaitGroup
	for name, serv := range s.c.services {
		// link := filepath.Join("/run", name)
		// s.c.log.Info("Creating symlink", "path", link, "target", self)
		// if err := os.Symlink(self, link); err != nil {
		// 	return err
		// }

		wg.Add(1)
		go func(name string, serv service.Service, wg *sync.WaitGroup) {
			delay := 1 * time.Second
			for {
				e := func() bool {
					e := false
					defer func() {
						if r := recover(); r != nil {
							stack := debug.Stack()
							s.c.log.Error(fmt.Errorf("%v", r), "Panic occurred in service and was recovered", "service", name, "stack", string(stack))
						}
					}()

					if err := serv.Run(); err != nil {
						s.c.log.Error(err, "Error occurred during service runtime", "service", name)
						e = false
					} else {
						s.c.log.Info("Service exited successfully", "service", name)
						e = true
					}

					return e
				}()
				if e {
					break
				}
				time.Sleep(delay)
				if delay < 30*time.Second {
					delay *= 2
				}
			}
			wg.Done()
		}(name, serv, &wg)
	}
	wg.Wait()

	return nil
}
