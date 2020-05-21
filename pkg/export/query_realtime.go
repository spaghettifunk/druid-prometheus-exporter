package export

// QueryRealtimeExporter contains all the Prometheus metrics that are possible to gather from the tranquillity service
type QueryRealtimeExporter struct {
}

// NewQueryRealtimeExporter returns a new Realtime exporter object
func NewQueryRealtimeExporter() *QueryRealtimeExporter {
	qr := &QueryRealtimeExporter{}

	// register all the prometheus metrics

	return qr
}
