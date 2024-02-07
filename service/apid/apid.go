// SPDX-License-Identifier: BSD-3-Clause

package apid

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/otelconnect"
	"connectrpc.com/vanguard"
	"github.com/google/uuid"
	compress "github.com/klauspost/connect-compress/v2"
	"github.com/nats-io/nats.go"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"github.com/quic-go/quic-go/logging"
	"github.com/quic-go/quic-go/qlog"
	"github.com/rs/cors"
	"github.com/u-bmc/operator/api/gen/umgmt/v1alpha1/umgmtv1alpha1connect"
	"github.com/u-bmc/operator/pkg/cert"
	"github.com/u-bmc/operator/pkg/log"
	"github.com/unrolled/secure"
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
		name:    DefaultName,
		id:      uuid.MustParse(DefaultUUID),
		log:     log.NewDefaultLogger(),
		ipcAddr: nats.DefaultURL,
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

func (s *Service) Run(ctx context.Context) error { //nolint:funlen
	s.c.log.Info("Starting service", "service", s.c.name, "uuid", s.c.id.String())

	s.c.log.Info("Connecting to ipcd", "service", s.c.name, "uuid", s.c.id.String(), "addr", s.c.ipcAddr)
	var (
		nc  *nats.Conn
		err error
	)
	for {
		nc, err = nats.Connect(s.c.ipcAddr)
		if err != nil {
			if errors.Is(err, nats.ErrNoServers) {
				time.Sleep(time.Second)
				continue
			}
			return err
		}
		break
	}

	s.c.log.Info("Creating u-mgmt server", "service", s.c.name, "uuid", s.c.id.String())
	oi, _ := otelconnect.NewInterceptor() //nolint:errcheck
	rpcRoute, rpcHandler := umgmtv1alpha1connect.NewUmgmtServiceHandler(
		&umgmtServiceServer{
			c:  s.c,
			nc: nc,
		},
		compress.WithAll(compress.LevelFastest),
		connect.WithInterceptors(
			oi,
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

	secureMiddleware := secure.New(secure.Options{
		HostsProxyHeaders:     []string{"X-Forwarded-Host"},
		SSLRedirect:           true,
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "script-src $NONCE",
	})

	handler := secureMiddleware.Handler(mux)

	corsMiddleware := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedOrigins: []string{"example.com"},
		AllowedHeaders: []string{
			"Accept-Encoding",
			"Content-Encoding",
			"Content-Type",
			"Connect-Protocol-Version",
			"Connect-Timeout-Ms",
			"Connect-Accept-Encoding",  // Unused in web browsers, but added for future-proofing
			"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
			"Grpc-Timeout",             // Used for gRPC-web
			"X-Grpc-Web",               // Used for gRPC-web
			"X-User-Agent",             // Used for gRPC-web
		},
		ExposedHeaders: []string{
			"Content-Encoding",         // Unused in web browsers, but added for future-proofing
			"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
			"Grpc-Status",              // Required for gRPC-web
			"Grpc-Message",             // Required for gRPC-web
		},
	})

	handler = corsMiddleware.Handler(handler)

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

	s.c.log.Info("Creating HTTP/3 server", "service", s.c.name, "uuid", s.c.id, "addr", "[::]:443", "protocol", "udp")
	h3 := http3.Server{
		Handler:    handler,
		QuicConfig: qconf,
		TLSConfig:  tconf,
	}

	s.c.log.Info("Creating HTTP/2 server", "service", s.c.name, "uuid", s.c.id, "addr", "[::]:443", "protocol", "tcp")
	h2 := &http.Server{
		Handler:           handler,
		TLSConfig:         tconf,
		ReadHeaderTimeout: 5 * time.Second,
	}

	s.c.log.Info("Starting API server", "service", s.c.name, "uuid", s.c.id.String())

	errChan := make(chan error, 1)

	go func() {
		errChan <- h3.ListenAndServe()
	}()

	go func() {
		// The certificate and key are already provided in the tls.Config
		errChan <- h2.ListenAndServeTLS("", "")
	}()

	select {
	case <-ctx.Done():
		s.c.log.Info("Shutting down API server", "service", s.c.name, "uuid", s.c.id.String())

		// Ignore error here as this best effort only.
		_ = h3.CloseGracefully(5 * time.Second)

		// As the parent context is canceled we'll make a new one that is still valid.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Ignore error here as this best effort only.
		_ = h2.Shutdown(ctx) //nolint:contextcheck
	case err := <-errChan:
		// Ignore error here as this best effort only.
		_ = h3.Close()
		_ = h2.Close()
		return err
	}

	return nil
}
