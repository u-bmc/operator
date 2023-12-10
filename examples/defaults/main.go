// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"github.com/u-bmc/operator"
	"github.com/u-bmc/operator/pkg/log"
	"github.com/u-bmc/operator/service"
	"github.com/u-bmc/operator/service/supervisord"
)

func main() {
	serviceMap := operator.NewDefaultServiceMap()
	serviceSlice := make([]service.Service, 0, len(serviceMap))
	for _, svc := range serviceMap {
		serviceSlice = append(serviceSlice, svc)
	}

	if err := operator.Launch(log.NewDefaultLogger(), supervisord.New(supervisord.WithServices(serviceSlice...)), serviceMap); err != nil {
		panic(err)
	}
}
