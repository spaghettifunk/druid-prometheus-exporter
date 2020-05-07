package export

import "github.com/prometheus/client_golang/prometheus"

// QueryJettyExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type QueryJettyExporter struct {
}

// NewQueryJettyExporter returns a new Jetty exporter object
func NewQueryJettyExporter() *QueryJettyExporter {
	qj := &QueryJettyExporter{}
	prometheus.MustRegister(qj)

	return qj
}

// Describe will associate the value for druid exporter
func (jq *QueryJettyExporter) Describe(ch chan<- *prometheus.Desc) {

}

// Collect returns the prometheus metrics from Druid's Jetty
func (jq *QueryJettyExporter) Collect(ch chan<- prometheus.Metric) {

}

// FormatMetrics .
func (jq *QueryJettyExporter) FormatMetrics() {
}
