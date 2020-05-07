package export

import "github.com/prometheus/client_golang/prometheus"

// QueryHistoricalExporter contains all the Prometheus metrics that are possible to gather from the historicals
type QueryHistoricalExporter struct {
	SegmentScanPending prometheus.Counter `description:"number of segments in queue waiting to be scanned"`
}

// NewQueryHistoricalExporter returns a new historical exporter object
func NewQueryHistoricalExporter() *QueryHistoricalExporter {
	qh := &QueryHistoricalExporter{
		SegmentScanPending: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "segment_scan_pending",
			Help:      "number of segments in queue waiting to be scanned",
			ConstLabels: prometheus.Labels{
				"historical": "segment-scan-pending",
			},
		}),
	}

	// register all prometheus metrics
	prometheus.MustRegister(qh.SegmentScanPending)

	return qh
}

// SetSegmentScanPending .
func (qh *QueryHistoricalExporter) SetSegmentScanPending(val float64) {

}
