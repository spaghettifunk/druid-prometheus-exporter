package export

import "github.com/prometheus/client_golang/prometheus"

// HealthHistoricalExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type HealthHistoricalExporter struct {
	SegmentMax           *prometheus.GaugeVec `description:"maximum byte limit available for segments"`
	SegmentUsed          *prometheus.GaugeVec `description:"bytes used for served segments"`
	SegmentUsedPercent   *prometheus.GaugeVec `description:"percentage of space used by served segments"`
	SegmentCount         *prometheus.GaugeVec `description:"number of served segments"`
	SegmentPendingDelete *prometheus.GaugeVec `description:"on-disk size in bytes of segments that are waiting to be cleared out"`
}

// NewHealthHistoricalExporter returns a new Jetty exporter object
func NewHealthHistoricalExporter() *HealthHistoricalExporter {
	hh := &HealthHistoricalExporter{
		SegmentMax: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "segment_max",
			Help:      "maximum byte limit available for segments",
		}, []string{}),
		SegmentUsed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "segment_used",
			Help:      "bytes used for served segments",
		}, []string{"dataSource", "tier", "priority"}),
		SegmentUsedPercent: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "segment_used_percent",
			Help:      "percentage of space used by served segments",
		}, []string{"dataSource", "tier", "priority"}),
		SegmentCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "segment_count",
			Help:      "number of served segments",
		}, []string{"dataSource", "tier", "priority"}),
		SegmentPendingDelete: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "segment_pending_delete",
			Help:      "on-disk size in bytes of segments that are waiting to be cleared out",
		}, []string{}),
	}

	// register all prometheus metrics
	prometheus.MustRegister(hh.SegmentMax)
	prometheus.MustRegister(hh.SegmentUsed)
	prometheus.MustRegister(hh.SegmentUsedPercent)
	prometheus.MustRegister(hh.SegmentCount)
	prometheus.MustRegister(hh.SegmentPendingDelete)

	return hh
}

// SetSegmentMax .
func (hh *HealthHistoricalExporter) SetSegmentMax(val float64) {
	hh.SegmentMax.WithLabelValues().Add(val)
}

// SetSegmentUsed .
func (hh *HealthHistoricalExporter) SetSegmentUsed(labels map[string]string, val float64) {
	hh.SegmentUsed.With(labels).Add(val)
}

// SetSegmentUsedPercent .
func (hh *HealthHistoricalExporter) SetSegmentUsedPercent(labels map[string]string, val float64) {
	hh.SegmentUsed.With(labels).Add(val)
}

// SetSegmentCount .
func (hh *HealthHistoricalExporter) SetSegmentCount(labels map[string]string, val float64) {
	hh.SegmentUsed.With(labels).Add(val)
}

// SetSegmentPendingDelete .
func (hh *HealthHistoricalExporter) SetSegmentPendingDelete(val float64) {
	hh.SegmentMax.WithLabelValues().Add(val)
}
