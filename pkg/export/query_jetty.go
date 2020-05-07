package export

import "github.com/prometheus/client_golang/prometheus"

// QueryJettyExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type QueryJettyExporter struct {
	NumOpenConnections prometheus.Counter `description:"number of open jetty connections"`
}

// NewQueryJettyExporter returns a new Jetty exporter object
func NewQueryJettyExporter() *QueryJettyExporter {
	qj := &QueryJettyExporter{
		NumOpenConnections: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "jetty",
			Name:      "num_open_connections",
			Help:      "number of open jetty connections",
			ConstLabels: prometheus.Labels{
				"sys": "num-open-connections",
			},
		}),
	}

	prometheus.MustRegister(qj.NumOpenConnections)

	return qj
}

// SetNumOpenConnections .
func (qj *QueryJettyExporter) SetNumOpenConnections(val float64) {

}
