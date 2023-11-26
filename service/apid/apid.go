// SPDX-License-Identifier: BSD-3-Clause

package apid

import (
	"context"
	"crypto/tls"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
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

func (s *Service) Run() error {
	s.c.log.Info("apid: starting")
	s.c.log.Info("apid: creating u-mgmt server")
	comp := compress.WithAll(compress.LevelFastest)
	svc := &umgmtServiceServer{
		name: s.c.name,
		id:   s.c.id,
		log:  s.c.log,
		c:    s.c.ipcClient,
	}
	rpcRoute, rpcHandler := umgmtv1alpha1connect.NewUmgmtServiceHandler(
		svc,
		comp,
		connect.WithInterceptors(
			otelconnect.NewInterceptor(
				otelconnect.WithTrustRemote(),
				otelconnect.WithoutServerPeerAttributes(),
			),
		),
	)

	s.c.log.Info("apid: creating vanguard transcoder")
	services := []*vanguard.Service{vanguard.NewService(rpcRoute, rpcHandler)}
	transcoder, err := vanguard.NewTranscoder(services)
	if err != nil {
		return err
	}

	s.c.log.Info("apid: creating http3 server")
	mux := http.NewServeMux()
	mux.Handle("/", transcoder)
	mux.Handle(grpchealth.NewHandler(grpchealth.NewStaticChecker(umgmtv1alpha1connect.UmgmtServiceName)))

	// TODO: Get from registry or config
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

			l := s.c.log.V(10).WithName("qlog").WithValues("connID", connID, "role", role)

			return qlog.NewConnectionTracer(qlogr{l}, p, connID)
		},
	}

	hs := http3.Server{
		Handler:    mux,
		QuicConfig: qconf,
		TLSConfig:  tconf,
	}

	return hs.ListenAndServe()
}
