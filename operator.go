// SPDX-License-Identifier: BSD-3-Clause

package operator

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/go-logr/logr"
	"github.com/u-bmc/operator/service"
	"github.com/u-bmc/operator/service/apid"
	"github.com/u-bmc/operator/service/hardwared"
	"github.com/u-bmc/operator/service/ipcd"
	"github.com/u-bmc/operator/service/netd"
	"github.com/u-bmc/operator/service/registryd"
	"go.opentelemetry.io/otel"
	"golang.org/x/sys/unix"
)

func Launch(log logr.Logger, supervisor service.Service, svcs map[string]service.Service) error {
	// If we have not been launched as the 'operator' that means we are being either
	// called from the cgroup shim or supervisord. In either case, we need to launch
	// the service that we are being called as. Supervisord will create a symlink
	// inside the /run directory for us to be called by.
	// if !strings.Contains(os.Args[0], "operator") {
	// 	shim := flag.Bool("shim", false, "Run the cgroup shim")
	// 	flag.Parse()

	// 	// When we request to be run via the shim we relaunch the same process
	// 	// inside a cgroup. As we need to call execve() to do this, we need to
	// 	// we need this flag do distinguish between being called by the shim or
	// 	// calling the shim from supervisord to relaunch the process.
	// 	if *shim {
	// 		return cgroup.ExecShim()
	// 	}

	// 	if svc, ok := svcs[os.Args[0]]; ok {
	// 		return svc.Run()
	// 	}
	// }

	// Should this be run here?
	otel.SetLogger(log)

	// Start Supervisor instance guarded by a Recoverer
	for i := 0; i < 10; i++ {
		err := func() error {
			defer func() {
				if r := recover(); r != nil {
					stack := debug.Stack()
					log.Error(fmt.Errorf("%v", r), "Panic occurred in supervisord and was recovered", "stack", string(stack))
				}
			}()

			return supervisor.Run()
		}()

		log.Error(err, "Fatal issue during supervisord runtime", "retries", i)

		// Wait a second before restarting supervisord
		time.Sleep(time.Second)
	}

	if os.Getpid() == 1 {
		if err := unix.Reboot(unix.LINUX_REBOOT_CMD_RESTART); err != nil {
			log.Error(err, "Failed to restart system")
		}
	}

	return fmt.Errorf("fatal operator error")
}

func NewDefaultServiceMap() map[string]service.Service {
	return service.NewServiceMap(
		ipcd.New(),
		registryd.New(),
		netd.New(),
		apid.New(),
		hardwared.New(),
		// TODO: kvmd, telemetryd and updated
	)
}
