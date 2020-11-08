// Package prometheus contains the prometheus init to be able to setup gauge metrics
package prometheus

import "github.com/prometheus/client_golang/prometheus"

var (
	// LinkUp is containing the GaugeVec struct
	LinkUp = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "sample",
			Subsystem: "external_url",
			Name:      "up",
			Help:      "is url up or down",
		},
		// We will want to monitor the worker ID that processed the
		// job, and the type of job that was processed
		[]string{"url"},
	)
	// ResponseMs is containing a GaugeVec struct
	ResponseMs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "sample",
			Subsystem: "external_url",
			Name:      "response_ms",
			Help:      "response time of the url",
		},
		// We will want to monitor the worker ID that processed the
		// job, and the type of job that was processed
		[]string{"url"},
	)
)

// Init is doing basic prometheus initialization
func init() {

	// register with the prometheus collector
	prometheus.MustRegister(LinkUp)
	prometheus.MustRegister(ResponseMs)

}
