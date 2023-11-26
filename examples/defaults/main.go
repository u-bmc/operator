// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"github.com/u-bmc/operator"
	"github.com/u-bmc/operator/pkg/log"
)

func main() {
	if err := operator.Launch(log.NewDefaultLogger(), operator.NewDefaultServiceMap()); err != nil {
		panic(err)
	}
}
