// SPDX-License-Identifier: BSD-3-Clause

package log

import (
	stdlog "log"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
)

func NewDefaultLogger() logr.Logger {
	return stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags))
}
