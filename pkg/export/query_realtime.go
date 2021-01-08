package export

import "github.com/prometheus/client_golang/prometheus"

// QueryRealtimeExporter contains all the Prometheus metrics that are possible to gather from the tranquillity service
type QueryRealtimeExporter struct {
	QueryTime             *prometheus.HistogramVec `description:"milliseconds taken to complete a query"`
	QueryWaitTime         *prometheus.HistogramVec `description:"milliseconds spent waiting for a segment to be scanned"`
	SegmentScanPending    *prometheus.GaugeVec     `description:"number of segments in queue waiting to be scanned"`
	QueryCount            *prometheus.GaugeVec     `description:"number of total queries"`
	QuerySuccessCount     *prometheus.GaugeVec     `description:"number of queries successfully processed"`
	QueryFailedCount      *prometheus.GaugeVec     `description:"number of failed queries"`
	QueryInterruptedCount *prometheus.GaugeVec     `description:"number of queries interrupted due to cancellation or timeout"`
}

// NewQueryRealtimeExporter returns a new Realtime exporter object
func NewQueryRealtimeExporter() *QueryRealtimeExporter {
	qr := &QueryRealtimeExporter{
		QueryTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "query_time",
			Help:      "milliseconds taken to complete a query",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"dataSource", "type", "hasFilters", "duration", "remoteAddress", "id", "numMetrics", "numComplexMetrics", "numDimensions", "threshold", "dimension"}),
		QueryWaitTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "query_wait_time",
			Help:      "milliseconds spent waiting for a segment to be scanned",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"id", "segment"}),
		SegmentScanPending: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "segment_scan_pending",
			Help:      "number of segments in queue waiting to be scanned",
		}, []string{}),
		QueryCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "query_count",
			Help:      "number of total queries",
		}, []string{}),
		QuerySuccessCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "query_success_count",
			Help:      "number of queries successfully processed",
		}, []string{}),
		QueryFailedCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{}),
		QueryInterruptedCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "query_interrupted_count",
			Help:      "number of queries interrupted due to cancellation or timeout",
		}, []string{}),
	}

	// register all the prometheus metrics
	prometheus.MustRegister(qr.QueryTime)
	prometheus.MustRegister(qr.QueryWaitTime)
	prometheus.MustRegister(qr.SegmentScanPending)
	prometheus.MustRegister(qr.QueryCount)
	prometheus.MustRegister(qr.QuerySuccessCount)
	prometheus.MustRegister(qr.QueryFailedCount)
	prometheus.MustRegister(qr.QueryInterruptedCount)

	return qr
}

// SetQueryTime .
func (qr *QueryRealtimeExporter) SetQueryTime(labels map[string]string, val float64) {
	qr.QueryTime.With(labels).Observe(val)
}

// SetQueryWaitTime .
func (qr *QueryRealtimeExporter) SetQueryWaitTime(labels map[string]string, val float64) {
	qr.QueryWaitTime.With(labels).Observe(val)
}

// SetQueryCount .
func (qr *QueryRealtimeExporter) SetQueryCount(val float64) {
	qr.QueryCount.WithLabelValues().Add(val)
}

// SetQuerySuccessCount .
func (qr *QueryRealtimeExporter) SetQuerySuccessCount(val float64) {
	qr.QuerySuccessCount.WithLabelValues().Add(val)
}

// SetQueryFailedCount .
func (qr *QueryRealtimeExporter) SetQueryFailedCount(val float64) {
	qr.QueryFailedCount.WithLabelValues().Add(val)
}

// SetQueryInterruptedCount .
func (qr *QueryRealtimeExporter) SetQueryInterruptedCount(val float64) {
	qr.QueryInterruptedCount.WithLabelValues().Add(val)
}

// SetSegmentScanPending .
func (qr *QueryRealtimeExporter) SetSegmentScanPending(val float64) {
	qr.SegmentScanPending.WithLabelValues().Add(val)
}
