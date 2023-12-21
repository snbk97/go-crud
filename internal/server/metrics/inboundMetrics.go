package server

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var InboundMetricsHistogram = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_inbound_latency_seconds",
		Help:    "Number of http requests",
		Buckets: []float64{0.002, 0.003, 0.005, 0.009, 0.03, 0.05, 0.09, 0.1, 0.2, 0.3, 0.4},
	},
	[]string{"app", "method", "endpoint", "status"},
)
