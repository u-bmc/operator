// SPDX-License-Identifier: BSD-3-Clause

package example

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	ipcv1alpha1 "github.com/u-bmc/operator/api/gen/ipc/v1alpha1"
	"github.com/u-bmc/operator/pkg/ipc"
	"google.golang.org/protobuf/types/known/structpb"
)

const ExamplePublisherUUID = "1822fb8a-24dc-4df8-bf24-4abfb6fcd11c"

type PublishService struct {
	name string
	id   uuid.UUID
}

func (s *PublishService) UUID() uuid.UUID {
	return s.id
}

func (s *PublishService) Name() string {
	return s.name
}

func (s *PublishService) Run(ctx context.Context) error {
	client := ipc.NewDefaultClient()

	log := logr.FromContextOrDiscard(ctx)

	topicName := "example-topic"

	data := make(map[string]any)
	data["example-int"] = 1
	data["example-str"] = "example"

	stpb, err := structpb.NewStruct(data)
	if err != nil {
		log.Error(err, "Failed to marshal data", "data", data)

		return err
	}

	for {
		res, err := client.Publish(ctx, connect.NewRequest(&ipcv1alpha1.PublishRequest{
			Topic:         topicName,
			PublisherName: s.name,
			PublisherId:   s.id.String(),
			Data:          []*structpb.Struct{stpb},
		}))
		if err != nil {
			log.Error(err, "Failed to publish message", "msg", stpb, "topic", topicName)

			return err
		}

		log.Info("received response", "status", res.Msg.Status.String())

		time.Sleep(500 * time.Microsecond)
	}
}
