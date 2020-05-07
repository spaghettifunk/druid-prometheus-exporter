package export

import "github.com/prometheus/client_golang/prometheus"

// CoordinationExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type CoordinationExporter struct {
	SegmentAssignedCount prometheus.Counter `description:"number of segments assigned to be loaded in the cluster"`
}

// NewCoordinationExporter returns a new Jetty exporter object
func NewCoordinationExporter() *CoordinationExporter {
	ce := &CoordinationExporter{
		SegmentAssignedCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "coordinator",
			Name:      "segment_assigned_count",
			Help:      "number of segments assigned to be loaded in the cluster",
			ConstLabels: prometheus.Labels{
				"coordinator": "segment-assigned-count",
			},
		}),
	}

	// register all the prometheus metrics
	prometheus.MustRegister(ce.SegmentAssignedCount)

	return ce
}

// SetSegmentAssignedCount .
func (ce *CoordinationExporter) SetSegmentAssignedCount(val float64) {

}
