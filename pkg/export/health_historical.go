package export

import "github.com/prometheus/client_golang/prometheus"

// HealthHistoricalExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type HealthHistoricalExporter struct {
}

// NewHealthHistoricalExporter returns a new Jetty exporter object
func NewHealthHistoricalExporter() *HealthHistoricalExporter {
	qj := &HealthHistoricalExporter{}
	prometheus.MustRegister(qj)

	return qj
}

// Describe will associate the value for druid exporter
func (hh *HealthHistoricalExporter) Describe(ch chan<- *prometheus.Desc) {

}

// Collect returns the prometheus metrics from Druid's Jetty
func (hh *HealthHistoricalExporter) Collect(ch chan<- prometheus.Metric) {

}

// FormatMetrics .
func (hh *HealthHistoricalExporter) FormatMetrics() {
}
