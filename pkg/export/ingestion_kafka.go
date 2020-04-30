package export

import "github.com/prometheus/client_golang/prometheus"

// IngestionKafkaExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type IngestionKafkaExporter struct {
}

// NewIngestionKafkaExporter returns a new Jetty exporter object
func NewIngestionKafkaExporter() *IngestionKafkaExporter {
	qj := &IngestionKafkaExporter{}
	prometheus.MustRegister(qj)

	return qj
}

// Describe will associate the value for druid exporter
func (bc *IngestionKafkaExporter) Describe(ch chan<- *prometheus.Desc) {

}

// Collect returns the prometheus metrics from Druid's Jetty
func (bc *IngestionKafkaExporter) Collect(ch chan<- prometheus.Metric) {

}
