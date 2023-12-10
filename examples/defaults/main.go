// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"github.com/u-bmc/operator"
	"github.com/u-bmc/operator/pkg/log"
	"github.com/u-bmc/operator/service/supervisord"
)

func main() {
	if err := operator.Launch(log.NewDefaultLogger(), supervisord.New(), operator.NewDefaultServiceMap()); err != nil {
		panic(err)
	}
}
