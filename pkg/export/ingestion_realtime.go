package export

// IngestionRealtimeExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type IngestionRealtimeExporter struct {
}

// NewIngestionRealtimeExporter returns a new Jetty exporter object
func NewIngestionRealtimeExporter() *IngestionRealtimeExporter {
	qj := &IngestionRealtimeExporter{}

	return qj
}
