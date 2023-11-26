// SPDX-License-Identifier: BSD-3-Clause

// Package user defines the user management functionality in the system.
package user

// Constants for user management operations.
const (
	// UserCreate represents the operation of creating a new user.
	UserCreate = "create_user"
	// UserDelete represents the operation of deleting an existing user.
	UserDelete = "delete_user"
	// UserUpdate represents the operation of updating an existing user's details.
	UserUpdate = "update_user"
	// UserGet represents the operation of retrieving the list of users.
	UserGet = "get_users"
	// UserCheckPassword represents the operation of verifying a user's password.
	UserCheckPassword = "check_password"
	// UserCheckRole represents the operation of checking a user's role.
	UserCheckRole = "check_role"
)

// Role represents the role assigned to a user.
type Role uint

// Constants for user roles.
const (
	// RoleDebug represents a user with debugging privileges.
	RoleDebug Role = iota
	// RoleAdmin represents a user with administrative privileges.
	RoleAdmin
	// RoleUser represents a regular user.
	RoleUser
)

// User represents a user in the system.
type User struct {
	// Username is the unique identifier of the user.
	Username string
	// Description provides additional information about the user.
	Description string
	// Role is the role assigned to the user.
	Role Role
	// PasswordHash is the hashed representation of the user's password.
	PasswordHash []byte
	// Salt is the salt used in the password hashing process.
	Salt []byte
}
