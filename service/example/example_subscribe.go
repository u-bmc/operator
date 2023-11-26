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
)

const ExampleSubscriberUUID = "94bf3066-f9f8-41f5-a701-54f77527c593"

type SubscribeService struct {
	name string
	id   uuid.UUID
}

func (s *SubscribeService) UUID() uuid.UUID {
	return s.id
}

func (s *SubscribeService) Name() string {
	return s.name
}

func (s *SubscribeService) Run(ctx context.Context) error {
	client := ipc.NewDefaultClient()

	log := logr.FromContextOrDiscard(ctx)

	topicName := "example-topic"

	res, err := client.Subscribe(ctx, connect.NewRequest(&ipcv1alpha1.SubscribeRequest{
		Topic:          topicName,
		SubscriberName: s.name,
		SubscriberId:   s.id.String(),
	}))
	if err != nil {
		log.Error(err, "Failed to subscribe to topic", "topic", topicName)

		return err
	}

	for {
		if !res.Receive() {
			time.Sleep(time.Second)
			continue
		}

		log.Info("received message", "msg", res.Msg())
	}
}
