// SPDX-License-Identifier: BSD-3-Clause

package apid

import (
	"context"
	"crypto/tls"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/otelconnect"
	"connectrpc.com/vanguard"
	"github.com/google/uuid"
	compress "github.com/klauspost/connect-compress/v2"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"github.com/quic-go/quic-go/logging"
	"github.com/quic-go/quic-go/qlog"
	"github.com/u-bmc/operator/api/gen/umgmt/v1alpha1/umgmtv1alpha1connect"
	"github.com/u-bmc/operator/pkg/cert"
	"github.com/u-bmc/operator/pkg/ipc"
	"github.com/u-bmc/operator/pkg/log"
)

const (
	DefaultName = "apid"
	DefaultUUID = "67d6ae85-1c0c-4e26-9cc8-841ef53b3ba0"
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

func (s *Service) Run(ctx context.Context) error {
	s.c.log.Info("Starting service", "service", s.c.name, "uuid", s.c.id.String())
	s.c.log.Info("Creating u-mgmt server", "service", s.c.name, "uuid", s.c.id.String())
	rpcRoute, rpcHandler := umgmtv1alpha1connect.NewUmgmtServiceHandler(
		&umgmtServiceServer{
			c: s.c,
		},
		compress.WithAll(compress.LevelFastest),
		connect.WithInterceptors(
			otelconnect.NewInterceptor(),
		),
	)

	s.c.log.Info("Creating vanguard transcoder", "service", s.c.name, "uuid", s.c.id.String())
	services := []*vanguard.Service{vanguard.NewService(rpcRoute, rpcHandler)}
	transcoder, err := vanguard.NewTranscoder(services)
	if err != nil {
		return err
	}

	s.c.log.Info("Creating HTTP/s multiplexer", "service", s.c.name, "uuid", s.c.id)
	mux := http.NewServeMux()
	mux.Handle("/", transcoder)

	s.c.log.Info("Adding gRPC health check to mux", "service", s.c.name, "uuid", s.c.id)
	mux.Handle(grpchealth.NewHandler(grpchealth.NewStaticChecker(umgmtv1alpha1connect.UmgmtServiceName)))

	s.c.log.Info("Adding gRPC reflector to mux", "service", s.c.name, "uuid", s.c.id)
	reflector := grpcreflect.NewStaticReflector(umgmtv1alpha1connect.UmgmtServiceName)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	s.c.log.Info("Generating certificates", "service", s.c.name, "uuid", s.c.id)
	// TODO: Get from registry or config
	// TODO: Change self signed generate function to behave the same as proper signed generate function
	certPem, keyPem, err := cert.GenerateSelfsigned("localhost")
	if err != nil {
		return err
	}
	c, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		return err
	}
	tconf := &tls.Config{
		Certificates: []tls.Certificate{c},
		MinVersion:   tls.VersionTLS13,
	}

	qconf := &quic.Config{
		Tracer: func(ctx context.Context, p logging.Perspective, connID quic.ConnectionID) *logging.ConnectionTracer {
			role := "server"
			if p == logging.PerspectiveClient {
				role = "client"
			}
			// Make this log level really high to not spam the logs
			l := s.c.log.V(10).WithName("qlog").WithValues("connID", connID, "role", role)

			return qlog.NewConnectionTracer(qlogr{l}, p, connID)
		},
	}

	s.c.log.Info("Creating HTTP/3 server", "service", s.c.name, "uuid", s.c.id, "addr", "[::]:443")
	hs := http3.Server{
		Handler:    mux,
		QuicConfig: qconf,
		TLSConfig:  tconf,
	}

	s.c.log.Info("Starting API server", "service", s.c.name, "uuid", s.c.id.String())

	return hs.ListenAndServe()
}
