// SPDX-License-Identifier: BSD-3-Clause

package log

import (
	stdlog "log"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
)

const (
	InfoLevel = iota
	DebugLevel
	TraceLevel
)

func NewDefaultStdLogger() logr.Logger {
	return stdr.New(stdlog.New(os.Stdout, "", stdlog.LstdFlags))
}

func NewDefaultLogger() logr.Logger {
	zl := zerolog.
		New(zerolog.NewConsoleWriter()).
		With().
		Timestamp().
		Logger()

	return zerologr.New(&zl)
}
