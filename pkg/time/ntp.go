// SPDX-License-Identifier: BSD-3-Clause

package time

import (
	"time"

	"github.com/beevik/ntp"
)

func GetTimeFromNTP(servers []string) (time.Time, error) {
	// Initialize an empty slice to store the offsets
	offsets := make([]time.Duration, 0, len(servers))
	var err error

	// Loop over all the servers
	for _, server := range servers {
		// Query the current server for the time
		resp, err := ntp.Query(server)
		// If there's an error, skip to the next server
		if err != nil {
			continue
		}

		// Validate the response from the server
		err = resp.Validate()
		// If there's an error, skip to the next server
		if err != nil {
			continue
		}

		// If there's no error, append the offset to the offsets slice
		offsets = append(offsets, resp.ClockOffset)
	}

	// If no offsets were collected, return an error
	if len(offsets) == 0 {
		return time.Now(), err
	}

	// Initialize a variable to store the total offset
	var total time.Duration
	// Loop over all the offsets and add them to the total
	for _, offset := range offsets {
		total += offset
	}

	// Calculate the mean offset
	meanOffset := total / time.Duration(len(offsets))
	// Return the current time adjusted by the mean offset
	return time.Now().Add(meanOffset), nil
}
