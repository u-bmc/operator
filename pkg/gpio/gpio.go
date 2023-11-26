// SPDX-License-Identifier: BSD-3-Clause

package gpio

import "github.com/warthog618/gpiod"

// ReadValue reads the value of a single GPIO line.
// It returns the value and any error encountered.
func ReadValue(gpio int) (int, error) {
	line, err := gpiod.RequestLine("gpiochip0", gpio, gpiod.AsInput)
	if err != nil {
		// Handle error if unable to request line
		return 0, err
	}
	defer line.Close()

	value, err := line.Value()

	// Return values separately so defer works as expected
	return value, err
}

// ReadValues reads the values of multiple GPIO lines.
// It returns a slice of values and any error encountered.
func ReadValues(gpios []int) ([]int, error) {
	lines, err := gpiod.RequestLines("gpiochip0", gpios, gpiod.AsInput)
	if err != nil {
		// Handle error if unable to request lines
		return nil, err
	}
	defer lines.Close()

	values := make([]int, len(gpios))
	if err := lines.Values(values); err != nil {
		// Handle error if unable to read values
		return nil, err
	}

	return values, nil
}

// WriteValue writes a value to a single GPIO line.
// It returns any error encountered.
func WriteValue(gpio int, value int) error {
	line, err := gpiod.RequestLine("gpiochip0", gpio, gpiod.AsOutput(value))
	if err != nil {
		// Handle error if unable to request line
		return err
	}
	defer line.Close()

	return nil
}

// WriteValues writes values to multiple GPIO lines.
// It returns any error encountered.
func WriteValues(gpios []int, values []int) error {
	lines, err := gpiod.RequestLines("gpiochip0", gpios, gpiod.AsOutput(values...))
	if err != nil {
		// Handle error if unable to request lines
		return err
	}
	defer lines.Close()

	return nil
}

// WatchLine watches a single GPIO line for changes.
// It returns a function to stop watching and any error encountered.
func WatchLine(gpio int, handler func(evt gpiod.LineEvent)) (func() error, error) {
	line, err := gpiod.RequestLine("gpiochip0", gpio, gpiod.WithBothEdges, gpiod.WithEventHandler(handler))
	if err != nil {
		// Handle error if unable to request line
		return nil, err
	}

	return line.Close, nil
}
