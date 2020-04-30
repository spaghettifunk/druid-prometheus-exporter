package export

import "github.com/prometheus/client_golang/prometheus"

// HealthFirehoseExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type HealthFirehoseExporter struct {
}

// NewHealthFirehoseExporter returns a new Jetty exporter object
func NewHealthFirehoseExporter() *HealthFirehoseExporter {
	qj := &HealthFirehoseExporter{}
	prometheus.MustRegister(qj)

	return qj
}

// Describe will associate the value for druid exporter
func (bc *HealthFirehoseExporter) Describe(ch chan<- *prometheus.Desc) {

}

// Collect returns the prometheus metrics from Druid's Jetty
func (bc *HealthFirehoseExporter) Collect(ch chan<- prometheus.Metric) {

}
