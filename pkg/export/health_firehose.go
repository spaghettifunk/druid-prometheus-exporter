package export

// HealthFirehoseExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type HealthFirehoseExporter struct {
}

// NewHealthFirehoseExporter returns a new Jetty exporter object
func NewHealthFirehoseExporter() *HealthFirehoseExporter {
	qj := &HealthFirehoseExporter{}

	return qj
}
