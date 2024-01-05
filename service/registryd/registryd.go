// SPDX-License-Identifier: BSD-3-Clause

package registryd

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	umgmtv1alpha1 "github.com/u-bmc/operator/api/gen/umgmt/v1alpha1"
	"github.com/u-bmc/operator/pkg/log"
	"github.com/u-bmc/operator/pkg/telemetry"
	"github.com/u-bmc/operator/pkg/version"
	bolt "go.etcd.io/bbolt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/crypto/argon2"
	"google.golang.org/protobuf/proto"
)

const (
	DefaultName   = "registryd"
	DefaultUUID   = "7fef9861-a94f-4808-ad84-f90c2b00d848"
	DefaultDBPath = "/var/registry.db"
)

type Service struct {
	c      config
	db     *bolt.DB
	tracer trace.Tracer
}

func New(opts ...Option) *Service {
	c := config{
		name:    DefaultName,
		id:      uuid.MustParse(DefaultUUID),
		log:     log.NewDefaultLogger(),
		ipcAddr: nats.DefaultURL,
		dbPath:  DefaultDBPath,
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
	s.tracer = otel.Tracer(
		fmt.Sprintf("%s/%s", s.c.name, s.c.id.String()),
		trace.WithInstrumentationVersion(version.SemVer),
	)

	s.c.log.Info("Connecting to registry", "path", s.c.dbPath)
	db, err := bolt.Open(s.c.dbPath, 0o600, &bolt.Options{
		Timeout: 5 * time.Second,
	})
	if err != nil {
		s.c.log.Error(err, "Failed to open registry", "path", s.c.dbPath)
		return err
	}
	s.db = db
	defer s.db.Close()

	s.c.log.Info("Connecting to ipcd", "service", s.c.name, "uuid", s.c.id.String(), "addr", s.c.ipcAddr)
	var nc *nats.Conn
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

	srv, err := micro.AddService(nc, micro.Config{
		Name:        s.c.name,
		Version:     version.SemVer,
		Description: "Handles registry operations",
	})
	if err != nil {
		return err
	}

	user := srv.AddGroup("user")

	if err := user.AddEndpoint("create", telemetry.TracedHandler(ctx, s.handleUserCreate)); err != nil {
		return err
	}

	if err := user.AddEndpoint("delete", telemetry.TracedHandler(ctx, s.handleUserDelete)); err != nil {
		return err
	}

	if err := user.AddEndpoint("list", telemetry.TracedHandler(ctx, s.handleUserList)); err != nil {
		return err
	}

	<-ctx.Done()

	return srv.Stop()
}

func (s *Service) handleUserCreate(ctx context.Context, req micro.Request) {
	ctx, span := s.tracer.Start(ctx, "handleUserCreate")

	user := &umgmtv1alpha1.User{}
	if err := proto.Unmarshal(req.Data(), user); err != nil {
		s.c.log.Error(err, "failed to unmarshal request data")
	}

	if user.Authentication.Method == umgmtv1alpha1.AuthenticationMethod_AUTHENTICATION_METHOD_PASSWORD {
		salt := make([]byte, 16)
		if _, err := rand.Read(salt); err != nil {
			s.c.log.Error(err, "failed to create salt", "service", s.c.name, "uuid", s.c.id.String())
		}

		if err := writeData(s.db, "usersalt", user.Name, salt); err != nil {
			s.c.log.Error(err, "failed to write salt", "service", s.c.name, "uuid", s.c.id.String())
		}

		user.Authentication.Data = argon2.IDKey(user.Authentication.Data, salt, 1, 64*1024, 4, 32)
	}

	if err := writeData(s.db, "user", user.Name, user); err != nil {
		s.c.log.Error(err, "failed to write user", "service", s.c.name, "uuid", s.c.id.String())
	}

	span.End()

	if err := req.Respond(nil, micro.WithHeaders(micro.Headers(telemetry.HeaderFromContext(ctx)))); err != nil {
		s.c.log.Error(err, "failed to respond to request")
	}
}

func (s *Service) handleUserDelete(ctx context.Context, req micro.Request) {
	ctx, span := s.tracer.Start(ctx, "handleUserDelete")

	if err := deleteData(s.db, "user", string(req.Data())); err != nil {
		s.c.log.Error(err, "failed to delete user", "service", s.c.name, "uuid", s.c.id.String())
	}

	if err := deleteData(s.db, "usersalt", string(req.Data())); err != nil {
		s.c.log.Error(err, "failed to delete user salt", "service", s.c.name, "uuid", s.c.id.String())
	}

	span.End()

	if err := req.Respond(nil, micro.WithHeaders(micro.Headers(telemetry.HeaderFromContext(ctx)))); err != nil {
		s.c.log.Error(err, "failed to respond to request")
	}
}

func (s *Service) handleUserList(ctx context.Context, req micro.Request) {
	ctx, span := s.tracer.Start(ctx, "handleUserGetAll")

	names, err := getKeys(s.db, "user")
	if err != nil {
		s.c.log.Error(err, "failed to get user list", "service", s.c.name, "uuid", s.c.id.String())
	}

	users := make([]*umgmtv1alpha1.User, len(names))

	for i, name := range names {
		user := &umgmtv1alpha1.User{}
		if err := readData(s.db, "user", name, user); err != nil {
			s.c.log.Error(err, "failed to read user", "service", s.c.name, "uuid", s.c.id.String())
		}

		users[i] = user
	}

	data, err := proto.Marshal(&umgmtv1alpha1.ListUsersResponse{
		Users: users,
	})
	if err != nil {
		s.c.log.Error(err, "failed to marshal user list", "service", s.c.name, "uuid", s.c.id.String())
	}

	span.End()

	if err := req.Respond(data, micro.WithHeaders(micro.Headers(telemetry.HeaderFromContext(ctx)))); err != nil {
		s.c.log.Error(err, "failed to respond to request")
	}
}
