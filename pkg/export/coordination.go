package export

import "github.com/prometheus/client_golang/prometheus"

// CoordinationExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type CoordinationExporter struct {
}

// NewCoordinationExporter returns a new Jetty exporter object
func NewCoordinationExporter() *CoordinationExporter {
	qj := &CoordinationExporter{}
	prometheus.MustRegister(qj)

	return qj
}

// Describe will associate the value for druid exporter
func (ce *CoordinationExporter) Describe(ch chan<- *prometheus.Desc) {

}

// Collect returns the prometheus metrics from Druid's Jetty
func (ce *CoordinationExporter) Collect(ch chan<- prometheus.Metric) {

}

// FormatMetrics .
func (ce *CoordinationExporter) FormatMetrics() {
}
