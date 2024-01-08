// SPDX-License-Identifier: BSD-3-Clause

package time

import (
	"encoding/base64"
	"time"

	"github.com/cloudflare/roughtime/client"
	"github.com/cloudflare/roughtime/config"
)

// RoughtimeServer represents a Roughtime server configuration.
type RoughtimeServer struct {
	// Protocol is the network protocol to use (e.g., "udp", "tcp").
	Protocol string
	// Address is the network address of the Roughtime server.
	Address string
	// PublicKey is the base64-encoded public key of the Roughtime server.
	PublicKey string
	// PublicKeyType is the type of the public key (e.g., "ed25519").
	PublicKeyType string
}

// GetTimeFromRoughtime queries the given Roughtime servers for the current time.
// It returns the mean time reported by the servers, or the local time if no servers could be queried successfully.
func GetTimeFromRoughtime(servers []RoughtimeServer) (time.Time, error) {
	// Initialize a slice to store the times reported by the servers.
	times := make([]time.Time, 0, len(servers))
	var err error

	// Loop over all the servers.
	for _, server := range servers {
		// Decode the server's public key from base64.
		pk, err := base64.StdEncoding.DecodeString(server.PublicKey)
		// If there's an error, skip to the next server.
		if err != nil {
			continue
		}

		// Create a config.Server object for the current server.
		srv := &config.Server{
			Name:          server.Address,
			PublicKeyType: server.PublicKeyType,
			PublicKey:     pk,
			Addresses: []config.ServerAddress{
				{Protocol: server.Protocol, Address: server.Address},
			},
		}

		// Query the current server for the time.
		res, err := client.Get(srv, 3, 5*time.Second, nil)
		// If there's an error, skip to the next server.
		if err != nil {
			continue
		}

		// Extract the time from the server's response.
		t, _ := res.Now()
		// Add the time to the slice of times.
		times = append(times, t)
	}

	// If no servers could be queried successfully, return the local time.
	if len(times) == 0 {
		return time.Now(), err
	}

	// Initialize a variable to store the total offset.
	var total time.Duration

	// Loop over all the times and add their offsets from the local time to the total.
	for _, t := range times {
		total += time.Until(t)
	}

	// Calculate the mean offset.
	meanOffset := total / time.Duration(len(times))

	// Return the current time adjusted by the mean offset.
	return time.Now().Add(meanOffset), nil
}
