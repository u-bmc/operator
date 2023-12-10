// SPDX-License-Identifier: BSD-3-Clause

package ipcd

import (
	"context"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"github.com/google/uuid"
	ipcv1alpha1 "github.com/u-bmc/operator/api/gen/ipc/v1alpha1"
	"github.com/u-bmc/operator/api/gen/ipc/v1alpha1/ipcv1alpha1connect"
	"github.com/u-bmc/operator/pkg/cache"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ipcServiceServer struct {
	ipcv1alpha1connect.UnimplementedIPCServiceHandler
	c     config
	cache *cache.Cache
}

func (s *ipcServiceServer) Publish(
	ctx context.Context,
	req *connect.Request[ipcv1alpha1.PublishRequest],
) (*connect.Response[ipcv1alpha1.PublishResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	b, err := proto.Marshal(req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	s.cache.Write(req.Msg.Topic, b)

	return connect.NewResponse(&ipcv1alpha1.PublishResponse{
		Status: ipcv1alpha1.Status_STATUS_SUCCESS,
	}), nil
}

func (s *ipcServiceServer) Subscribe(
	ctx context.Context,
	req *connect.Request[ipcv1alpha1.SubscribeRequest],
	stream *connect.ServerStream[ipcv1alpha1.SubscribeResponse],
) error {
	if err := validate(req.Msg); err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}

	for {
		if ctx.Err() != nil {
			return connect.NewError(connect.CodeCanceled, ctx.Err())
		}

		b, ok := s.cache.Read(req.Msg.Topic)
		if !ok {
			continue
		}

		m := &ipcv1alpha1.PublishRequest{}
		if err := proto.Unmarshal(b, m); err != nil {
			return connect.NewError(connect.CodeInternal, err)
		}

		id, err := uuid.NewRandom()
		if err != nil {
			return connect.NewError(connect.CodeInternal, err)
		}

		if err := stream.Send(&ipcv1alpha1.SubscribeResponse{
			Timestamp:   timestamppb.Now(),
			Topic:       m.Topic,
			PublisherId: m.PublisherId,
			MessageId:   id.String(),
			Data:        m.Data,
		}); err != nil {
			return connect.NewError(connect.CodeInternal, err)
		}
	}
}

func validate(msg protoreflect.ProtoMessage) error {
	v, err := protovalidate.New(protovalidate.WithFailFast(true))
	if err != nil {
		return err
	}

	return v.Validate(msg)
}
