// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

// gRPC Prometheus monitoring interceptors for client-side gRPC.

package grpc_prometheus

import (
	prom "github.com/prometheus/client_golang/prometheus"
)

var (
	// DefaultClientMetrics is the default instance of ClientMetrics. It is
	// intended to be used in conjunction the default Prometheus metrics
	// registry.
	DefaultClientMetrics = NewClientMetrics()

	// UnaryClientInterceptor is a gRPC client-side interceptor that provides Prometheus monitoring for Unary RPCs.
	UnaryClientInterceptor = DefaultClientMetrics.UnaryClientInterceptor()

	// StreamClientInterceptor is a gRPC client-side interceptor that provides Prometheus monitoring for Streaming RPCs.
	StreamClientInterceptor = DefaultClientMetrics.StreamClientInterceptor()
)

func init() {
	prom.MustRegister(DefaultClientMetrics.clientStartedCounter)
	prom.MustRegister(DefaultClientMetrics.clientHandledCounter)
	prom.MustRegister(DefaultClientMetrics.clientStreamMsgReceived)
	prom.MustRegister(DefaultClientMetrics.clientStreamMsgSent)
}

func WithClientOptions(opts ...Option) {
	DefaultClientMetrics.WithOptions(opts...)
}

// EnableClientHandlingTimeHistogram turns on recording of handling time of
// RPCs. Histogram metrics can be very expensive for Prometheus to retain and
// query. This function acts on the DefaultClientMetrics variable and the
// default Prometheus metrics registry.
func EnableClientHandlingTimeHistogram(opts ...HistogramOption) {
	DefaultClientMetrics.EnableClientHandlingTimeHistogram(opts...)
	prom.Register(DefaultClientMetrics.clientHandledHistogram)
}

// EnableClientStreamReceiveTimeHistogram turns on recording of
// single message receive time of streaming RPCs.
// This function acts on the DefaultClientMetrics variable and the
// default Prometheus metrics registry.
func EnableClientStreamReceiveTimeHistogram(opts ...HistogramOption) {
	DefaultClientMetrics.EnableClientStreamReceiveTimeHistogram(opts...)
	prom.Register(DefaultClientMetrics.clientStreamRecvHistogram)
}

// EnableClientStreamSendTimeHistogram turns on recording of
// single message send time of streaming RPCs.
// This function acts on the DefaultClientMetrics variable and the
// default Prometheus metrics registry.
func EnableClientStreamSendTimeHistogram(opts ...HistogramOption) {
	DefaultClientMetrics.EnableClientStreamSendTimeHistogram(opts...)
	prom.Register(DefaultClientMetrics.clientStreamSendHistogram)
}
