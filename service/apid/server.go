// SPDX-License-Identifier: BSD-3-Clause

package apid

import (
	"context"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"github.com/nats-io/nats.go"
	umgmtv1alpha1 "github.com/u-bmc/operator/api/gen/umgmt/v1alpha1"
	"github.com/u-bmc/operator/api/gen/umgmt/v1alpha1/umgmtv1alpha1connect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type umgmtServiceServer struct {
	umgmtv1alpha1connect.UnimplementedUmgmtServiceHandler
	c  config
	nc *nats.Conn
}

func (s *umgmtServiceServer) ListUsers(ctx context.Context, req *connect.Request[umgmtv1alpha1.ListUsersRequest]) (*connect.Response[umgmtv1alpha1.ListUsersResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	data, err := proto.Marshal(req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	msg, err := s.nc.RequestWithContext(ctx, "registry.user.list", data)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var resp umgmtv1alpha1.ListUsersResponse
	if err := proto.Unmarshal(msg.Data, &resp); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&resp), nil
}

func (s *umgmtServiceServer) CreateUser(ctx context.Context, req *connect.Request[umgmtv1alpha1.CreateUserRequest]) (*connect.Response[umgmtv1alpha1.CreateUserResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	data, err := proto.Marshal(req.Msg)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	msg, err := s.nc.RequestWithContext(ctx, "registry.user.create", data)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	var resp umgmtv1alpha1.CreateUserResponse
	if err := proto.Unmarshal(msg.Data, &resp); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&resp), nil
}

func validate(msg protoreflect.ProtoMessage) error {
	v, err := protovalidate.New(protovalidate.WithFailFast(true))
	if err != nil {
		return err
	}

	return v.Validate(msg)
}
