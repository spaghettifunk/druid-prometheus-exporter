package export

// IngestionKafkaExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type IngestionKafkaExporter struct {
}

// NewIngestionKafkaExporter returns a new Jetty exporter object
func NewIngestionKafkaExporter() *IngestionKafkaExporter {
	qj := &IngestionKafkaExporter{}

	return qj
}
