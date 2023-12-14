// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"context"

	"github.com/u-bmc/operator"
)

func main() {
	if err := operator.Launch(context.Background()); err != nil {
		panic(err)
	}
}
