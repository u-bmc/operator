// SPDX-License-Identifier: BSD-3-Clause

package hwmon

import (
	"os"
	"path/filepath"
	"strings"
)

type SensorEntry struct {
	Type  string
	Kind  string
	Value string
	Name  string
	Path  string
}

const SysfsBasePath = "/sys/class/hwmon"

func GetSensors() ([]SensorEntry, error) {
	entries, err := os.ReadDir(SysfsBasePath)
	if err != nil {
		return nil, err
	}

	var Sensors []SensorEntry
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		files, err := os.ReadDir(filepath.Join(SysfsBasePath, entry.Name()))
		if err != nil {
			return nil, err
		}

		name, err := os.ReadFile(filepath.Join(SysfsBasePath, entry.Name(), "name"))
		if err != nil {
			return nil, err
		}

		for _, f := range files {
			if f.IsDir() || !strings.Contains(f.Name(), "_") {
				continue
			}

			value, err := os.ReadFile(filepath.Join(SysfsBasePath, entry.Name(), f.Name()))
			if err != nil {
				return nil, err
			}

			components := strings.SplitN(f.Name(), "_", 2)

			Sensors = append(Sensors, SensorEntry{
				Type:  components[0],
				Kind:  components[1],
				Value: string(value),
				Name:  string(name),
				Path:  filepath.Join(SysfsBasePath, entry.Name(), f.Name()),
			})
		}
	}

	return Sensors, nil
}
