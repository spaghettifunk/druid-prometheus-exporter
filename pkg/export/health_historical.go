package export

import "github.com/prometheus/client_golang/prometheus"

// HealthHistoricalExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type HealthHistoricalExporter struct {
	SegmentMax prometheus.Counter `description:"maximum byte limit available for segments"`
}

// NewHealthHistoricalExporter returns a new Jetty exporter object
func NewHealthHistoricalExporter() *HealthHistoricalExporter {
	hh := &HealthHistoricalExporter{
		SegmentMax: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "segment_max",
			Help:      "maximum byte limit available for segments",
			ConstLabels: prometheus.Labels{
				"historical": "segment-max",
			},
		}),
	}

	// register all prometheus metrics
	prometheus.MustRegister(hh.SegmentMax)

	return hh
}

// SetSegmentMax .
func (hh *HealthHistoricalExporter) SetSegmentMax(val float64) {

}
