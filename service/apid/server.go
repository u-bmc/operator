// SPDX-License-Identifier: BSD-3-Clause

package apid

import (
	"context"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	ipcv1alpha1 "github.com/u-bmc/operator/api/gen/ipc/v1alpha1"
	umgmtv1alpha1 "github.com/u-bmc/operator/api/gen/umgmt/v1alpha1"
	"github.com/u-bmc/operator/api/gen/umgmt/v1alpha1/umgmtv1alpha1connect"
	"github.com/u-bmc/operator/pkg/user"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type umgmtServiceServer struct {
	umgmtv1alpha1connect.UnimplementedUmgmtServiceHandler
	c config
}

func (s *umgmtServiceServer) GetUsers(ctx context.Context, req *connect.Request[umgmtv1alpha1.GetUsersRequest]) (*connect.Response[umgmtv1alpha1.GetUsersResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	resp, err := s.c.ipcClient.Subscribe(ctx, connect.NewRequest(&ipcv1alpha1.SubscribeRequest{
		Topic:          user.UserGet,
		SubscriberName: s.c.name,
		SubscriberId:   s.c.id.String(),
	}))
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	if !resp.Receive() {
		return nil, connect.NewError(connect.CodeInternal, resp.Err())
	}

	msg := resp.Msg()
	s.c.log.Info("Received message", "msg", msg)

	users := make([]*user.User, len(msg.Data))
	for i, data := range msg.Data {
		dm := data.AsMap()
		u := &user.User{
			Username:    dm["username"].(string),
			Description: dm["description"].(string),
			Role:        dm["role"].(user.Role),
		}

		users[i] = u
	}

	userpb := make([]*umgmtv1alpha1.User, len(users))
	for i, u := range users {
		userpb[i] = &umgmtv1alpha1.User{
			Name:        u.Username,
			Description: u.Description,
			Role:        umgmtv1alpha1.Role(u.Role),
		}
	}

	return connect.NewResponse(&umgmtv1alpha1.GetUsersResponse{
		Users:  userpb,
		Status: umgmtv1alpha1.Status_STATUS_SUCCESS,
	}), nil
}

func (s *umgmtServiceServer) GetUserInfo(ctx context.Context, req *connect.Request[umgmtv1alpha1.GetUserInfoRequest]) (*connect.Response[umgmtv1alpha1.GetUserInfoResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&umgmtv1alpha1.GetUserInfoResponse{
		Status: umgmtv1alpha1.Status_STATUS_SUCCESS,
	}), nil
}

func (s *umgmtServiceServer) UpdateUser(ctx context.Context, req *connect.Request[umgmtv1alpha1.UpdateUserRequest]) (*connect.Response[umgmtv1alpha1.UpdateUserResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&umgmtv1alpha1.UpdateUserResponse{
		Status: umgmtv1alpha1.Status_STATUS_SUCCESS,
	}), nil
}

func (s *umgmtServiceServer) GetInventory(ctx context.Context, req *connect.Request[umgmtv1alpha1.GetInventoryRequest]) (*connect.Response[umgmtv1alpha1.GetInventoryResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&umgmtv1alpha1.GetInventoryResponse{
		Status: umgmtv1alpha1.Status_STATUS_SUCCESS,
	}), nil
}

func (s *umgmtServiceServer) GetMachineInfo(ctx context.Context, req *connect.Request[umgmtv1alpha1.GetMachineInfoRequest]) (*connect.Response[umgmtv1alpha1.GetMachineInfoResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&umgmtv1alpha1.GetMachineInfoResponse{
		Status: umgmtv1alpha1.Status_STATUS_SUCCESS,
	}), nil
}

func (s *umgmtServiceServer) GetMachineState(ctx context.Context, req *connect.Request[umgmtv1alpha1.GetMachineStateRequest]) (*connect.Response[umgmtv1alpha1.GetMachineStateResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&umgmtv1alpha1.GetMachineStateResponse{
		Status: umgmtv1alpha1.Status_STATUS_SUCCESS,
	}), nil
}

func (s *umgmtServiceServer) ChangeMachineState(ctx context.Context, req *connect.Request[umgmtv1alpha1.ChangeMachineStateRequest]) (*connect.Response[umgmtv1alpha1.ChangeMachineStateResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&umgmtv1alpha1.ChangeMachineStateResponse{
		Status: umgmtv1alpha1.Status_STATUS_SUCCESS,
	}), nil
}

func (s *umgmtServiceServer) GetSensorList(ctx context.Context, req *connect.Request[umgmtv1alpha1.GetSensorListRequest]) (*connect.Response[umgmtv1alpha1.GetSensorListResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&umgmtv1alpha1.GetSensorListResponse{
		Status: umgmtv1alpha1.Status_STATUS_SUCCESS,
	}), nil
}

func (s *umgmtServiceServer) GetSensorData(ctx context.Context, req *connect.Request[umgmtv1alpha1.GetSensorDataRequest]) (*connect.Response[umgmtv1alpha1.GetSensorDataResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&umgmtv1alpha1.GetSensorDataResponse{
		Status: umgmtv1alpha1.Status_STATUS_SUCCESS,
	}), nil
}

func (s *umgmtServiceServer) StreamHostConsole(ctx context.Context, stream *connect.BidiStream[umgmtv1alpha1.StreamHostConsoleRequest, umgmtv1alpha1.StreamHostConsoleResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, nil)
}

func (s *umgmtServiceServer) ConfigureThermalSetPoints(ctx context.Context, req *connect.Request[umgmtv1alpha1.ConfigureThermalSetPointsRequest]) (*connect.Response[umgmtv1alpha1.ConfigureThermalSetPointsResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&umgmtv1alpha1.ConfigureThermalSetPointsResponse{
		Status: umgmtv1alpha1.Status_STATUS_SUCCESS,
	}), nil
}

func (s *umgmtServiceServer) ConfigureThermalFanProfiles(ctx context.Context, req *connect.Request[umgmtv1alpha1.ConfigureThermalFanProfilesRequest]) (*connect.Response[umgmtv1alpha1.ConfigureThermalFanProfilesResponse], error) {
	if err := validate(req.Msg); err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	return connect.NewResponse(&umgmtv1alpha1.ConfigureThermalFanProfilesResponse{
		Status: umgmtv1alpha1.Status_STATUS_SUCCESS,
	}), nil
}

func validate(msg protoreflect.ProtoMessage) error {
	v, err := protovalidate.New(protovalidate.WithFailFast(true))
	if err != nil {
		return err
	}

	return v.Validate(msg)
}
