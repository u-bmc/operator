// SPDX-License-Identifier: BSD-3-Clause

package ipc

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/u-bmc/operator/api/gen/ipc/v1alpha1/ipcv1alpha1connect"
	"golang.org/x/net/http2"
)

func NewDefaultClient() ipcv1alpha1connect.IPCServiceClient {
	return NewClient("localhost", "10984")
}

func NewClient(addr, port string) ipcv1alpha1connect.IPCServiceClient {
	return ipcv1alpha1connect.NewIPCServiceClient(
		&http.Client{
			Transport: &http2.Transport{
				AllowHTTP: true,
				DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
					return net.Dial(network, addr)
				},
			},
			Timeout: 10 * time.Second,
		},
		fmt.Sprintf("http://%s:%s", addr, port),
	)
}
