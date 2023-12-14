// SPDX-License-Identifier: BSD-3-Clause

package ipcd

import (
	"context"
	"net"
	"net/http"
	"os"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/otelconnect"
	"github.com/google/uuid"
	"github.com/u-bmc/operator/api/gen/ipc/v1alpha1/ipcv1alpha1connect"
	"github.com/u-bmc/operator/pkg/cache"
	"github.com/u-bmc/operator/pkg/ipc"
	"github.com/u-bmc/operator/pkg/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const (
	DefaultName     = "ipcd"
	DefaultUUID     = "7d7f58a8-71dd-4e9b-9fb1-fe524f6f9942"
	DefaultAddr     = "localhost:10984"
	DefaultAddrType = ipc.TCP
)

type Service struct {
	c config
}

func New(opts ...Option) *Service {
	c := config{
		name:     DefaultName,
		id:       uuid.MustParse(DefaultUUID),
		log:      log.NewDefaultLogger(),
		addr:     DefaultAddr,
		addrType: DefaultAddrType,
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

func (s *Service) Run(ctx context.Context) error {
	s.c.log.Info("Starting service", "service", s.c.name, "uuid", s.c.id.String())
	s.c.log.Info("Creating IPC server", "addr", s.c.addr, "addrType", s.c.addrType)
	var (
		l    net.Listener
		addr string
		err  error
	)
	if s.c.addrType == ipc.Unix {
		if err := os.Remove(s.c.addr); err != nil {
			return err
		}

		l, err = net.Listen("unix", s.c.addr)
		if err != nil {
			return err
		}

		addr = ""
	} else {
		l, err = net.Listen("tcp", s.c.addr)
		if err != nil {
			return err
		}

		addr = s.c.addr
	}

	s.c.log.Info("Creating message cache", "ttl", 5*time.Second, "maxEntries", 10)
	ca := cache.NewCache(ctx, 5*time.Second, 10)

	s.c.log.Info("Creating HTTP/s multiplexer", "service", s.c.name, "uuid", s.c.id)
	mux := http.NewServeMux()
	mux.Handle(ipcv1alpha1connect.NewIPCServiceHandler(
		&ipcServiceServer{
			c:     s.c,
			cache: ca,
		},
		connect.WithInterceptors(otelconnect.NewInterceptor(
			otelconnect.WithTrustRemote(),
			otelconnect.WithoutServerPeerAttributes(),
		)),
	))

	s.c.log.Info("Adding gRPC health check to mux", "service", s.c.name, "uuid", s.c.id)
	mux.Handle(grpchealth.NewHandler(grpchealth.NewStaticChecker(ipcv1alpha1connect.IPCServiceName)))

	s.c.log.Info("Creating HTTP/2 server", "service", s.c.name, "uuid", s.c.id)
	hs := &http.Server{ //nolint:gosec
		Addr:    addr,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	s.c.log.Info("Starting IPC server", "addr", s.c.addr, "addrType", s.c.addrType)

	return hs.Serve(l)
}
