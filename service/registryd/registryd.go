// SPDX-License-Identifier: BSD-3-Clause

package registryd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	ipcv1alpha1 "github.com/u-bmc/operator/api/gen/ipc/v1alpha1"
	"github.com/u-bmc/operator/pkg/ipc"
	"github.com/u-bmc/operator/pkg/log"
	"github.com/u-bmc/operator/pkg/user"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	DefaultName   = "registryd"
	DefaultUUID   = "7fef9861-a94f-4808-ad84-f90c2b00d848"
	DefaultDBPath = "/var/registry.db"
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
		dbPath:    DefaultDBPath,
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
	s.c.log.Info("registryd: starting")
	s.c.log.Info("registryd: connecting to registry")
	db, err := bolt.Open(s.c.dbPath, 0o600, &bolt.Options{
		Timeout: 5 * time.Second,
	})
	if err != nil {
		s.c.log.Error(err, "registryd: failed to open registry")
		return err
	}
	defer db.Close()

	ctx := context.Background()
	stream, err := s.c.ipcClient.Subscribe(ctx, connect.NewRequest(&ipcv1alpha1.SubscribeRequest{
		Topic:          "registry",
		SubscriberName: s.c.name,
		SubscriberId:   s.c.id.String(),
	}))
	if err != nil {
		s.c.log.Error(err, "Failed to subscribe to topic", "topic", "registry")

		return err
	}

	for stream.Receive() {
		if stream.Msg().Topic != "registry" {
			continue
		}

		recipient := stream.Msg().PublisherName

		s.c.log.V(10).Info("registryd: received message", "ID", stream.Msg().MessageId, "from", stream.Msg().PublisherName)

		msg := stream.Msg().GetData()
		for _, v := range msg {
			s.createUserHelper(ctx, db, v, recipient)
			s.deleteUserHelper(ctx, db, v, recipient)
			s.updateUserHelper(ctx, db, v, recipient)
			s.getUsersHelper(ctx, db, v, recipient)
			s.checkPasswordHelper(ctx, db, v, recipient)
			s.checkRoleHelper(ctx, db, v, recipient)
		}
	}

	s.c.log.Error(stream.Err(), "registryd: stream ended unexpectedly")

	return fmt.Errorf("unexpected stream end: %w", stream.Err())
}

// createUserHelper is a helper function that creates a new user.
func (s *Service) createUserHelper(ctx context.Context, db *bolt.DB, st *structpb.Struct, recipient string) {
	// Check if UserCreate field exists in the struct
	if f, ok := st.Fields[user.UserCreate]; ok {
		// Initialize status map with status set to true
		status := map[string]any{
			"status": true,
		}

		// Attempt to create user with provided data
		if err := createUser(db, f.GetStructValue()); err != nil {
			// Log error and set status to false if user creation fails
			s.c.log.Error(err, "registryd: failed to create user")
			status["status"] = false
		}

		// Marshal status into a structpb.Struct
		ss, err := structpb.NewStruct(status)
		if err != nil {
			// Log error and return if marshaling fails
			s.c.log.Error(err, "registryd: failed to marshal status")
			return
		}

		// Publish the status to the recipient
		res, err := s.c.ipcClient.Publish(ctx, connect.NewRequest(&ipcv1alpha1.PublishRequest{
			Topic:         recipient,
			PublisherName: s.c.name,
			PublisherId:   s.c.id.String(),
			Data:          []*structpb.Struct{ss},
		}))
		if err != nil {
			// Log error and return if publishing fails
			s.c.log.Error(err, "registryd: failed to publish response", "topic", recipient)
			return
		}

		// Log the status of the published response
		s.c.log.Info("registryd: published response", "status", res.Msg.Status.String())
	}
}

func (s *Service) deleteUserHelper(ctx context.Context, db *bolt.DB, st *structpb.Struct, recipient string) {
	if f, ok := st.Fields[user.UserDelete]; ok {
		status := map[string]any{
			"status": true,
		}

		if err := deleteUser(db, f.GetStringValue()); err != nil {
			s.c.log.Error(err, "registryd: failed to delete user")
			status["status"] = false
		}

		ss, err := structpb.NewStruct(status)
		if err != nil {
			s.c.log.Error(err, "registryd: failed to marshal status")
		}

		res, err := s.c.ipcClient.Publish(ctx, connect.NewRequest(&ipcv1alpha1.PublishRequest{
			Topic:         recipient,
			PublisherName: s.c.name,
			PublisherId:   s.c.id.String(),
			Data:          []*structpb.Struct{ss},
		}))
		if err != nil {
			s.c.log.Error(err, "registryd: failed to publish response", "topic", recipient)
			return
		}

		s.c.log.Info("registryd: published response", "status", res.Msg.Status.String())
	}
}

func (s *Service) updateUserHelper(ctx context.Context, db *bolt.DB, st *structpb.Struct, recipient string) {
	if f, ok := st.Fields[user.UserUpdate]; ok {
		status := map[string]any{
			"status": true,
		}

		if err := updateUser(db, f.GetStructValue()); err != nil {
			s.c.log.Error(err, "registryd: failed to update user")
			status["status"] = false
		}

		ss, err := structpb.NewStruct(status)
		if err != nil {
			s.c.log.Error(err, "registryd: failed to marshal status")
			return
		}

		res, err := s.c.ipcClient.Publish(ctx, connect.NewRequest(&ipcv1alpha1.PublishRequest{
			Topic:         recipient,
			PublisherName: s.c.name,
			PublisherId:   s.c.id.String(),
			Data:          []*structpb.Struct{ss},
		}))
		if err != nil {
			s.c.log.Error(err, "registryd: failed to publish response", "topic", recipient)
			return
		}

		s.c.log.Info("registryd: published response", "status", res.Msg.Status.String())
	}
}

func (s *Service) getUsersHelper(ctx context.Context, db *bolt.DB, st *structpb.Struct, recipient string) {
	if _, ok := st.Fields[user.UserGet]; ok {
		users, err := getUsers(db)
		if err != nil {
			s.c.log.Error(err, "registryd: failed to get users")
		}

		data := make([]*structpb.Struct, 0, len(users))
		for _, u := range users {
			um := make(map[string]any)
			um["username"] = u.Username
			um["description"] = u.Description
			um["role"] = u.Role
			us, err := structpb.NewStruct(um)
			if err != nil {
				s.c.log.Error(err, "registryd: failed to marshal user", "user", u)
			}
			data = append(data, us)
		}

		res, err := s.c.ipcClient.Publish(ctx, connect.NewRequest(&ipcv1alpha1.PublishRequest{
			Topic:         recipient,
			PublisherName: s.c.name,
			PublisherId:   s.c.id.String(),
			Data:          data,
		}))
		if err != nil {
			s.c.log.Error(err, "registryd: failed to publish response", "topic", recipient)
			return
		}

		s.c.log.Info("registryd: published response", "status", res.Msg.Status.String())
	}
}

func (s *Service) checkPasswordHelper(ctx context.Context, db *bolt.DB, st *structpb.Struct, recipient string) {
	if f, ok := st.Fields[user.UserCheckPassword]; ok {
		status := map[string]any{
			"status": true,
		}

		data := strings.Split(f.GetStringValue(), ":")

		if len(data) != 2 {
			s.c.log.Error(fmt.Errorf("invalid data"), "registryd: invalid data", "data", data)
			return
		}

		if ok, err := checkPassword(db, data[0], data[1]); err != nil || !ok {
			s.c.log.Error(err, "registryd: failed to check password")
			status["status"] = false
		}

		ss, err := structpb.NewStruct(status)
		if err != nil {
			s.c.log.Error(err, "registryd: failed to marshal status")
			return
		}

		res, err := s.c.ipcClient.Publish(ctx, connect.NewRequest(&ipcv1alpha1.PublishRequest{
			Topic:         recipient,
			PublisherName: s.c.name,
			PublisherId:   s.c.id.String(),
			Data:          []*structpb.Struct{ss},
		}))
		if err != nil {
			s.c.log.Error(err, "registryd: failed to publish response", "topic", recipient)
			return
		}

		s.c.log.Info("registryd: published response", "status", res.Msg.Status.String())
	}
}

func (s *Service) checkRoleHelper(ctx context.Context, db *bolt.DB, st *structpb.Struct, recipient string) { //nolint:cyclop
	if f, ok := st.Fields[user.UserCheckRole]; ok {
		status := map[string]any{
			"status": true,
		}

		data := strings.Split(f.GetStringValue(), ":")

		if len(data) != 2 {
			s.c.log.Error(fmt.Errorf("invalid data"), "registryd: invalid data", "data", data)
			return
		}

		var role user.Role
		switch data[1] {
		case "debug":
			role = user.RoleDebug
		case "admin":
			role = user.RoleAdmin
		case "user":
			role = user.RoleUser
		default:
			s.c.log.Error(fmt.Errorf("invalid role"), "registryd: invalid role", "role", data[1])
			return
		}

		if ok, err := checkRole(db, data[0], role); err != nil || !ok {
			s.c.log.Error(err, "registryd: failed to check role")
			status["status"] = false
		}

		ss, err := structpb.NewStruct(status)
		if err != nil {
			s.c.log.Error(err, "registryd: failed to marshal status")
			return
		}

		res, err := s.c.ipcClient.Publish(ctx, connect.NewRequest(&ipcv1alpha1.PublishRequest{
			Topic:         recipient,
			PublisherName: s.c.name,
			PublisherId:   s.c.id.String(),
			Data:          []*structpb.Struct{ss},
		}))
		if err != nil {
			s.c.log.Error(err, "registryd: failed to publish response", "topic", recipient)
			return
		}

		s.c.log.Info("registryd: published response", "status", res.Msg.Status.String())
	}
}
