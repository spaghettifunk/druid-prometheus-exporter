package export

import "github.com/prometheus/client_golang/prometheus"

// QueryHistoricalExporter contains all the Prometheus metrics that are possible to gather from the historicals
type QueryHistoricalExporter struct {
}

// NewQueryHistoricalExporter returns a new historical exporter object
func NewQueryHistoricalExporter() *QueryHistoricalExporter {
	qh := &QueryHistoricalExporter{}
	prometheus.MustRegister(qh)

	return qh
}

// Describe will associate the value for druid exporter
func (qh *QueryHistoricalExporter) Describe(ch chan<- *prometheus.Desc) {

}

// Collect returns the prometheus metrics from Druid's historicals
func (qh *QueryHistoricalExporter) Collect(ch chan<- prometheus.Metric) {

}

// FormatMetrics .
func (qh *QueryHistoricalExporter) FormatMetrics() {
}
