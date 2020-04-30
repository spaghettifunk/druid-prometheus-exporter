package export

import "github.com/prometheus/client_golang/prometheus"

// QueryRealtimeExporter contains all the Prometheus metrics that are possible to gather from the tranquillity service
type QueryRealtimeExporter struct {
}

// NewQueryRealtimeExporter returns a new Realtime exporter object
func NewQueryRealtimeExporter() *QueryRealtimeExporter {
	qr := &QueryRealtimeExporter{}
	prometheus.MustRegister(qr)

	return qr
}

// Describe will associate the value for druid exporter
func (bc *QueryRealtimeExporter) Describe(ch chan<- *prometheus.Desc) {

}

// Collect returns the prometheus metrics from Druid's Tranquillity
func (bc *QueryRealtimeExporter) Collect(ch chan<- prometheus.Metric) {

}
