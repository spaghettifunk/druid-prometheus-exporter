package export

import "github.com/prometheus/client_golang/prometheus"

// SQLExporter contains all the Prometheus metrics that are possible to gather from the SQL service
type SQLExporter struct {
}

// NewSQLExporter returns a new SQL exporter object
func NewSQLExporter() *SQLExporter {
	qj := &SQLExporter{}
	prometheus.MustRegister(qj)

	return qj
}

// Describe will associate the value for druid exporter
func (bc *SQLExporter) Describe(ch chan<- *prometheus.Desc) {

}

// Collect returns the prometheus metrics from Druid's SQL
func (bc *SQLExporter) Collect(ch chan<- prometheus.Metric) {

}
