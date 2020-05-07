package export

// IngestionRealtimeIndexingExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type IngestionRealtimeIndexingExporter struct {
}

// NewIngestionRealtimeIndexingExporter returns a new Jetty exporter object
func NewIngestionRealtimeIndexingExporter() *IngestionRealtimeIndexingExporter {
	qj := &IngestionRealtimeIndexingExporter{}

	return qj
}
