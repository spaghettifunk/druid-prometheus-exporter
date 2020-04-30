package export

import "github.com/prometheus/client_golang/prometheus"

// SysExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type SysExporter struct {
}

// NewSysExporter returns a new Jetty exporter object
func NewSysExporter() *SysExporter {
	qj := &SysExporter{}
	prometheus.MustRegister(qj)

	return qj
}

// Describe will associate the value for druid exporter
func (bc *SysExporter) Describe(ch chan<- *prometheus.Desc) {

}

// Collect returns the prometheus metrics from Druid's Jetty
func (bc *SysExporter) Collect(ch chan<- prometheus.Metric) {

}
