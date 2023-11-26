// SPDX-License-Identifier: BSD-3-Clause

package apid

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	ipcv1alpha1 "github.com/u-bmc/operator/api/gen/ipc/v1alpha1"
	"github.com/u-bmc/operator/api/gen/ipc/v1alpha1/ipcv1alpha1connect"
	umgmtv1alpha1 "github.com/u-bmc/operator/api/gen/umgmt/v1alpha1"
	"github.com/u-bmc/operator/api/gen/umgmt/v1alpha1/umgmtv1alpha1connect"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type umgmtServiceServer struct {
	umgmtv1alpha1connect.UnimplementedUmgmtServiceHandler
	id   uuid.UUID
	name string
	log  logr.Logger
	c    ipcv1alpha1connect.IPCServiceClient
}

func (s *umgmtServiceServer) GetUsers(ctx context.Context, req *connect.Request[umgmtv1alpha1.GetUsersRequest]) (*connect.Response[umgmtv1alpha1.GetUsersResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	topicName := "get-users"
	publishRes, err := s.c.Publish(ctx, connect.NewRequest(&ipcv1alpha1.PublishRequest{
		Topic:         topicName,
		PublisherName: s.name,
		PublisherId:   s.id.String(),
	}))
	if err != nil {
		return nil, err
	}

	// Check the status of the publish operation
	if publishRes.Msg.Status != ipcv1alpha1.Status_STATUS_SUCCESS {
		return nil, connect.NewError(connect.CodeInternal, errors.New(publishRes.Msg.Status.String()))
	}

	_, err = s.c.Subscribe(ctx, connect.NewRequest(&ipcv1alpha1.SubscribeRequest{
		Topic:          topicName,
		SubscriberName: s.name,
		SubscriberId:   s.id.String(),
	}))
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&umgmtv1alpha1.GetUsersResponse{}), nil
}

func validate(msg protoreflect.ProtoMessage) error {
	v, err := protovalidate.New(protovalidate.WithFailFast(true))
	if err != nil {
		return err
	}

	return v.Validate(msg)
}
