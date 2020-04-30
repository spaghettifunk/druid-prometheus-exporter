package export

import "github.com/prometheus/client_golang/prometheus"

// QueryBrokerExporter contains all the Prometheus metrics that are possible to gather from the brokers
type QueryBrokerExporter struct {
}

// NewQueryBrokerExporter returns a new broker exporter object
func NewQueryBrokerExporter() *QueryBrokerExporter {
	qb := &QueryBrokerExporter{}
	prometheus.MustRegister(qb)

	return qb
}

// Describe will associate the value for druid exporter
func (bc *QueryBrokerExporter) Describe(ch chan<- *prometheus.Desc) {

}

// Collect returns the prometheus metrics from Druid's Brokers
func (bc *QueryBrokerExporter) Collect(ch chan<- prometheus.Metric) {

}
