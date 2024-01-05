// SPDX-License-Identifier: BSD-3-Clause

package operator

import (
	"context"

	"cirello.io/oversight"
	"github.com/go-logr/logr"
	"github.com/u-bmc/operator/pkg/log"
	"github.com/u-bmc/operator/pkg/telemetry"
	"github.com/u-bmc/operator/pkg/version"
	"github.com/u-bmc/operator/service"
	"github.com/u-bmc/operator/service/apid"
	"github.com/u-bmc/operator/service/hardwared"
	"github.com/u-bmc/operator/service/ipcd"
	"github.com/u-bmc/operator/service/kvmd"
	"github.com/u-bmc/operator/service/netd"
	"github.com/u-bmc/operator/service/registryd"
	"github.com/u-bmc/operator/service/telemetryd"
	"github.com/u-bmc/operator/service/updated"
)

func Launch(ctx context.Context, opts ...Option) error {
	c := config{
		log:  log.NewDefaultLogger(),
		svcs: NewDefaultServiceMap(),
	}

	for _, opt := range opts {
		opt.apply(&c)
	}

	// Note: Tracing currently noop until configurable
	telemetry.SetupOtelSDK(ctx, telemetry.NewResource("u-bmc", version.SemVer), c.log)

	c.log.Info("Starting u-bmc operator", "version", version.Version())

	p := make([]oversight.ChildProcessSpecification, len(c.svcs))
	for i, svc := range c.svcs {
		p[i] = oversight.ChildProcessSpecification{
			Name:    svc.Name(),
			Start:   svc.Run,
			Restart: oversight.Permanent(),
		}
	}

	supervise := oversight.New(
		oversight.WithSpecification(-1, 0, oversight.OneForOne()),
		oversight.WithLogger(log.NewOversightLogger(c.log)),
		oversight.Process(p...),
	)

	return supervise.Start(ctx)
}

type config struct {
	log  logr.Logger
	svcs []service.Service
}

type Option interface {
	apply(*config)
}

type logOption struct {
	log logr.Logger
}

func (o *logOption) apply(c *config) {
	c.log = o.log
}

func WithLogger(l logr.Logger) Option {
	return &logOption{
		log: l,
	}
}

type serviceMapOption struct {
	svcs []service.Service
}

func (o *serviceMapOption) apply(c *config) {
	c.svcs = o.svcs
}

func WithServiceMap(svcs []service.Service) Option {
	return &serviceMapOption{
		svcs: svcs,
	}
}

func NewDefaultServiceMap() []service.Service {
	return []service.Service{
		ipcd.New(),
		registryd.New(),
		netd.New(),
		apid.New(),
		hardwared.New(),
		kvmd.New(),
		telemetryd.New(),
		updated.New(),
	}
}
