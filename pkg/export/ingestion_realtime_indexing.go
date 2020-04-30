package export

import "github.com/prometheus/client_golang/prometheus"

// IngestionRealtimeIndexingExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type IngestionRealtimeIndexingExporter struct {
}

// NewIngestionRealtimeIndexingExporter returns a new Jetty exporter object
func NewIngestionRealtimeIndexingExporter() *IngestionRealtimeIndexingExporter {
	qj := &IngestionRealtimeIndexingExporter{}
	prometheus.MustRegister(qj)

	return qj
}

// Describe will associate the value for druid exporter
func (bc *IngestionRealtimeIndexingExporter) Describe(ch chan<- *prometheus.Desc) {

}

// Collect returns the prometheus metrics from Druid's Jetty
func (bc *IngestionRealtimeIndexingExporter) Collect(ch chan<- prometheus.Metric) {

}
