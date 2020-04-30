package export

import "github.com/prometheus/client_golang/prometheus"

// IngestionRealtimeExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type IngestionRealtimeExporter struct {
}

// NewIngestionRealtimeExporter returns a new Jetty exporter object
func NewIngestionRealtimeExporter() *IngestionRealtimeExporter {
	qj := &IngestionRealtimeExporter{}
	prometheus.MustRegister(qj)

	return qj
}

// Describe will associate the value for druid exporter
func (bc *IngestionRealtimeExporter) Describe(ch chan<- *prometheus.Desc) {

}

// Collect returns the prometheus metrics from Druid's Jetty
func (bc *IngestionRealtimeExporter) Collect(ch chan<- prometheus.Metric) {

}
