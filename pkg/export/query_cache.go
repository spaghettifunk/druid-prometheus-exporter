package export

import "github.com/prometheus/client_golang/prometheus"

// QueryCacheExporter contains all the Prometheus metrics that are possible to gather from the tranquillity service
type QueryCacheExporter struct {
}

// NewQueryCacheExporter returns a new Cache exporter object
func NewQueryCacheExporter() *QueryCacheExporter {
	qc := &QueryCacheExporter{}
	prometheus.MustRegister(qc)

	return qc
}

// Describe will associate the value for druid exporter
func (bc *QueryCacheExporter) Describe(ch chan<- *prometheus.Desc) {

}

// Collect returns the prometheus metrics from Druid's Tranquillity
func (bc *QueryCacheExporter) Collect(ch chan<- prometheus.Metric) {

}
