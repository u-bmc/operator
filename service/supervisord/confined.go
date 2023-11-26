// SPDX-License-Identifier: BSD-3-Clause

package supervisord

import "github.com/u-bmc/operator/pkg/cgroup"

func RunProcessConfined(name string, args ...string) error {
	cg, err := cgroup.New(name)
	if err != nil {
		return err
	}
	// Technically we should check any error here but we also
	// need this do be run deferred.
	defer cg.Teardown() //nolint: errcheck

	return cg.Run(args)
}
