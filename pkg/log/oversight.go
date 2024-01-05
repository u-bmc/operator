// SPDX-License-Identifier: BSD-3-Clause

package log

import (
	"fmt"

	"cirello.io/oversight"
	"github.com/go-logr/logr"
)

type OversightLogger struct {
	l logr.Logger
}

func (l *OversightLogger) Printf(format string, v ...interface{}) {
	l.l.V(InfoLevel).Info(fmt.Sprintf(format, v...))
}

func (l *OversightLogger) Println(v ...interface{}) {
	l.l.V(InfoLevel).Info(fmt.Sprint(v...))
}

func NewOversightLogger(l logr.Logger) oversight.Logger {
	return &OversightLogger{
		l: l,
	}
}
