// SPDX-License-Identifier: BSD-3-Clause

package supervisord

import (
	"github.com/u-bmc/operator/pkg/cgroup"
	"github.com/u-bmc/operator/service"
)

func RunServiceConfined(svc service.Service) error {
	cg, err := cgroup.New(svc.Name())
	if err != nil {
		return err
	}
	// Technically we should check any error here but we also
	// need this do be run deferred.
	defer cg.Teardown() //nolint: errcheck

	return cg.Run()
}
