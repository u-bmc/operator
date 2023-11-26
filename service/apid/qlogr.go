// SPDX-License-Identifier: BSD-3-Clause

package apid

import "github.com/go-logr/logr"

type qlogr struct {
	l logr.Logger
}

func (q qlogr) Write(p []byte) (int, error) {
	q.l.Info(string(p))

	return len(p), nil
}

func (q qlogr) Close() error {
	return nil
}
