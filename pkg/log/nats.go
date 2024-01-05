// SPDX-License-Identifier: BSD-3-Clause

package log

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/nats-io/nats-server/v2/server"
)

type NATSLogger struct {
	l logr.Logger
}

func (l *NATSLogger) Fatalf(format string, v ...interface{}) {
	l.l.Error(fmt.Errorf(format, v...), "NATS fatal error")
}

func (l *NATSLogger) Errorf(format string, v ...interface{}) {
	l.l.Error(fmt.Errorf(format, v...), "NATS error")
}

func (l *NATSLogger) Warnf(format string, v ...interface{}) {
	l.l.Error(fmt.Errorf(format, v...), "NATS warning")
}

func (l *NATSLogger) Noticef(format string, v ...interface{}) {
	l.l.V(InfoLevel).Info(fmt.Sprintf(format, v...))
}

func (l *NATSLogger) Debugf(format string, v ...interface{}) {
	l.l.V(DebugLevel).Info(fmt.Sprintf(format, v...))
}

func (l *NATSLogger) Tracef(format string, v ...interface{}) {
	l.l.V(TraceLevel).Info(fmt.Sprintf(format, v...))
}

func NewNATSLogger(l logr.Logger) server.Logger {
	return &NATSLogger{
		l: l,
	}
}
