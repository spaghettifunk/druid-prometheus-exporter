package export

import "github.com/prometheus/client_golang/prometheus"

// HealthJVMExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type HealthJVMExporter struct {
}

// NewHealthJVMExporter returns a new Jetty exporter object
func NewHealthJVMExporter() *HealthJVMExporter {
	qj := &HealthJVMExporter{}
	prometheus.MustRegister(qj)

	return qj
}

// Describe will associate the value for druid exporter
func (bc *HealthJVMExporter) Describe(ch chan<- *prometheus.Desc) {

}

// Collect returns the prometheus metrics from Druid's Jetty
func (bc *HealthJVMExporter) Collect(ch chan<- prometheus.Metric) {

}
