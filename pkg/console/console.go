// SPDX-License-Identifier: BSD-3-Clause

// Package console provides functions to interact with the console via a serial port.
package console

import (
	"fmt"
	"io"

	"go.bug.st/serial"
)

// GetConsole opens a serial port at the specified path and baud rate.
// It returns an io.ReadWriteCloser for the port, or an error if the port cannot be opened.
func GetConsole(path string, baud int) (io.ReadWriteCloser, error) {
	// Get the list of available serial ports
	ports, err := serial.GetPortsList()
	if err != nil {
		// Return the error if we can't get the list
		return nil, err
	}

	// Check if the specified port is in the list of available ports
	found := false
	for _, port := range ports {
		if port == path {
			found = true
			break
		}
	}
	if !found {
		// Return an error if the specified port is not found
		return nil, fmt.Errorf("port not found")
	}

	// Open the specified port with the specified baud rate and settings
	port, err := serial.Open(path, &serial.Mode{
		BaudRate: baud,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	})
	if err != nil {
		// Return the error if we can't open the port
		return nil, err
	}

	// Return the opened port
	return port, nil
}
