// SPDX-License-Identifier: BSD-3-Clause

package ipc

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/peterbourgon/unixtransport"
	"github.com/u-bmc/operator/api/gen/ipc/v1alpha1/ipcv1alpha1connect"
	"golang.org/x/net/http2"
)

type Transport string

const (
	TCP  Transport = "http://"
	Unix Transport = "http+unix://"
)

func NewDefaultClient() ipcv1alpha1connect.IPCServiceClient {
	c, _ := NewClient(TCP, "localhost:10984")
	return c
}

func NewClient(t Transport, addr string) (ipcv1alpha1connect.IPCServiceClient, error) {
	var (
		client  *http.Client
		baseURL string
	)

	switch t {
	case TCP:
		client = &http.Client{
			Transport: &http2.Transport{
				AllowHTTP: true,
				DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
					return net.Dial(network, addr)
				},
				ReadIdleTimeout: 10 * time.Second,
			},
		}
		baseURL = fmt.Sprintf("%s%s", t, addr)
	case Unix:
		tp := &http.Transport{}
		unixtransport.Register(tp)
		client = &http.Client{
			Transport: tp,
		}
		baseURL = fmt.Sprintf("%s%s", t, addr)
	default:
		return nil, fmt.Errorf("unsupported transport: %s", t)
	}

	ipcClient := ipcv1alpha1connect.NewIPCServiceClient(
		client,
		baseURL,
	)

	return ipcClient, nil
}
