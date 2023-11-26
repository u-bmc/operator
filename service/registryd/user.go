// SPDX-License-Identifier: BSD-3-Clause

package registryd

import (
	"crypto/rand"
	"fmt"

	"github.com/u-bmc/operator/pkg/user"
	bolt "go.etcd.io/bbolt"
	"golang.org/x/crypto/argon2"
	"google.golang.org/protobuf/types/known/structpb"
)

func createUser(db *bolt.DB, userData *structpb.Struct) error {
	username, ok := userData.Fields["username"]
	if !ok {
		return fmt.Errorf("username not provided")
	}

	password, ok := userData.Fields["password"]
	if !ok {
		return fmt.Errorf("password not provided")
	}

	description, ok := userData.Fields["description"]
	if !ok {
		description = structpb.NewStringValue("")
	}

	role, ok := userData.Fields["role"]
	if !ok {
		role = structpb.NewNumberValue(float64(user.RoleUser))
	}

	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return fmt.Errorf("failed to generate salt: %w", err)
	}

	u := user.User{
		Username:     username.GetStringValue(),
		Description:  description.GetStringValue(),
		Role:         user.Role(role.GetNumberValue()),
		Salt:         salt,
		PasswordHash: argon2.IDKey([]byte(password.GetStringValue()), salt, 1, 64*1024, 4, 32),
	}

	return writeData(db, "users", username.GetStringValue(), u)
}

func deleteUser(db *bolt.DB, username string) error {
	return deleteData(db, "users", username)
}

func updateUser(db *bolt.DB, userData *structpb.Struct) error {
	username, ok := userData.Fields["username"]
	if !ok {
		return fmt.Errorf("username not provided")
	}

	var u user.User
	if err := readData(db, "users", username.GetStringValue(), &u); err != nil {
		return err
	}

	password, ok := userData.Fields["password"]
	if ok {
		u.PasswordHash = argon2.IDKey([]byte(password.GetStringValue()), u.Salt, 1, 64*1024, 4, 32)
	}

	description, ok := userData.Fields["description"]
	if ok {
		u.Description = description.GetStringValue()
	}

	role, ok := userData.Fields["role"]
	if ok {
		u.Role = user.Role(role.GetNumberValue())
	}

	return writeData(db, "users", username.GetStringValue(), u)
}

func getUser(db *bolt.DB, username string) (*user.User, error) {
	var u user.User
	err := readData(db, "users", username, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func getUsers(db *bolt.DB) ([]*user.User, error) {
	usernames, err := getKeys(db, "users")
	if err != nil {
		return nil, err
	}

	users := make([]*user.User, 0, len(usernames))
	for _, username := range usernames {
		u, err := getUser(db, username)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func checkPassword(db *bolt.DB, username, password string) (bool, error) {
	u, err := getUser(db, username)
	if err != nil {
		return false, err
	}

	hash := argon2.IDKey([]byte(password), u.Salt, 1, 64*1024, 4, 32)

	return string(hash) == string(u.PasswordHash), nil
}

func checkRole(db *bolt.DB, username string, role user.Role) (bool, error) {
	u, err := getUser(db, username)
	if err != nil {
		return false, err
	}

	return u.Role <= role, nil
}
