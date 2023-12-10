// SPDX-License-Identifier: BSD-3-Clause

package kvmd

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	ipcv1alpha1 "github.com/u-bmc/operator/api/gen/ipc/v1alpha1"
	"github.com/u-bmc/operator/pkg/ipc"
	"github.com/u-bmc/operator/pkg/log"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	DefaultName = "kvmd"
	DefaultUUID = "4421ddba-bdac-4350-9769-78803cf97215"
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

func (s *Service) Run() error {
	s.c.log.Info("Starting service", "service", s.c.name, "uuid", s.c.id.String())
	ctx := context.Background()

	for {
		time.Sleep(5 * time.Second)
		spb, err := structpb.NewStruct(map[string]interface{}{
			"foo": "bar",
		})
		if err != nil {
			continue
		}

		_, err = s.c.ipcClient.Publish(ctx, connect.NewRequest(&ipcv1alpha1.PublishRequest{
			Topic:         "kvmd",
			PublisherName: s.c.name,
			PublisherId:   s.c.id.String(),
			Data:          []*structpb.Struct{spb},
		}))
		if err != nil {
			s.c.log.Error(err, "Failed to publish response", "topic", "kvmd", "service", s.c.name, "uuid", s.c.id.String())
			continue
		}
	}
}
