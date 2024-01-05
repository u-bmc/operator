// SPDX-License-Identifier: BSD-3-Clause

package telemetry

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func NATSHeaderToOTLP(header nats.Header) propagation.HeaderCarrier {
	if header == nil {
		return nil
	}

	totalLength := 0
	for _, value := range header {
		totalLength += len(value)
	}

	sharedValues := make([]string, totalLength)
	output := make(propagation.HeaderCarrier, len(header))

	for key, value := range header {
		if value == nil {
			// Preserve nil values. ReverseProxy distinguishes
			// between nil and zero-length header values.
			output[key] = nil
			continue
		}

		n := copy(sharedValues, value)
		output[key] = sharedValues[:n:n]
		sharedValues = sharedValues[n:]
	}

	return output
}

func OTLPHeaderToNATS(header propagation.HeaderCarrier) nats.Header {
	if header == nil {
		return nil
	}

	totalLength := 0
	for _, value := range header {
		totalLength += len(value)
	}

	sharedValues := make([]string, totalLength)
	output := make(nats.Header, len(header))

	for key, value := range header {
		if value == nil {
			// Preserve nil values. ReverseProxy distinguishes
			// between nil and zero-length header values.
			output[key] = nil
			continue
		}

		n := copy(sharedValues, value)
		output[key] = sharedValues[:n:n]
		sharedValues = sharedValues[n:]
	}

	return output
}

func HeaderFromContext(ctx context.Context) nats.Header {
	prop := otel.GetTextMapPropagator()
	headers := make(propagation.HeaderCarrier)
	prop.Inject(ctx, headers)

	return OTLPHeaderToNATS(headers)
}

func TracedHandler(ctx context.Context, handler func(ctx context.Context, req micro.Request)) micro.Handler {
	return micro.HandlerFunc(func(req micro.Request) {
		prop := otel.GetTextMapPropagator()
		headers := NATSHeaderToOTLP(nats.Header(req.Headers()))
		handler(prop.Extract(ctx, headers), req)
	})
}

func TracedMsg(ctx context.Context, subject string, data []byte) *nats.Msg {
	prop := otel.GetTextMapPropagator()
	headers := make(propagation.HeaderCarrier)
	prop.Inject(ctx, headers)

	return &nats.Msg{
		Header:  OTLPHeaderToNATS(headers),
		Subject: subject,
		Data:    data,
	}
}
