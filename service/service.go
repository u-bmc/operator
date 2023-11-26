// SPDX-License-Identifier: BSD-3-Clause

package service

import "github.com/google/uuid"

// Service is an interface for long running processes or daemons.
// A service will be restarted if it returns an error.
// If a service returns nil, it is regarded to be done aka a oneshot service.
// Both name and UUID should be unique per machine.
// The UUID should be unique across machines as well as the name can be duplicate if multiple machines are connected.
type Service interface {
	// UUID returns the unique identifier of the service.
	// This should be unique per machine and across machines.
	UUID() uuid.UUID

	// Name returns the unique name of the service.
	// This should be unique per machine.
	Name() string

	// Run starts the service with the provided context.
	// It returns an error if the service needs to be restarted.
	Run() error
}

// NewServiceMap creates a map of services with their names as keys.
// It takes a slice of pointers to Service objects as input and returns a map.
func NewServiceMap(svcs ...Service) map[string]Service {
	// Initialize a map with size equal to the number of services
	m := make(map[string]Service, len(svcs))

	// Iterate over the services and add them to the map
	for _, svc := range svcs {
		// Use the service name as the key
		m[svc.Name()] = svc
	}

	// Return the populated map
	return m
}
