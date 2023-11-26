// SPDX-License-Identifier: BSD-3-Clause

package ipc

import (
	"context"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/otelconnect"
	"github.com/bufbuild/protovalidate-go"
	"github.com/google/uuid"
	ipcv1alpha1 "github.com/u-bmc/operator/api/gen/ipc/v1alpha1"
	"github.com/u-bmc/operator/api/gen/ipc/v1alpha1/ipcv1alpha1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ServiceServer struct {
	ipcv1alpha1connect.UnimplementedIPCServiceHandler
	c *Cache
}

func NewDefaultServer() *http.Server {
	return NewServer(":10984", NewCache(context.Background(), 5*time.Second, 10))
}

func NewServer(addr string, cache *Cache) *http.Server {
	mux := http.NewServeMux()
	mux.Handle(ipcv1alpha1connect.NewIPCServiceHandler(
		&ServiceServer{
			c: cache,
		},
		connect.WithInterceptors(otelconnect.NewInterceptor(
			otelconnect.WithTrustRemote(),
			otelconnect.WithoutServerPeerAttributes(),
		)),
	))

	mux.Handle(grpchealth.NewHandler(grpchealth.NewStaticChecker(ipcv1alpha1connect.IPCServiceName)))

	hs := &http.Server{
		Addr:              addr,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 5 * time.Second,
	}

	return hs
}

func (s *ServiceServer) Publish(
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

	if ok := s.c.Write(req.Msg.Topic, req.Msg.PublisherName, b); !ok {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&ipcv1alpha1.PublishResponse{
		Status: 0,
	}), nil
}

func (s *ServiceServer) Subscribe(
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

		b, ok := s.c.Read(req.Msg.Topic, "test")
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
