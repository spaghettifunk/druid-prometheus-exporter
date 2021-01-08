package export

import "github.com/prometheus/client_golang/prometheus"

// QueryHistoricalExporter contains all the Prometheus metrics that are possible to gather from the historicals
type QueryHistoricalExporter struct {
	QueryTime                *prometheus.HistogramVec `description:"milliseconds taken to complete a query"`
	QuerySegmentTime         *prometheus.HistogramVec `description:"milliseconds taken to query individual segment. Includes time to page in the segment from disk"`
	QueryWaitTime            *prometheus.HistogramVec `description:"milliseconds spent waiting for a segment to be scanned"`
	QuerySegmentAndCacheTime *prometheus.HistogramVec `description:"milliseconds taken to query individual segment or hit the cache (if it is enabled on the Historical process)"`
	QueryCPUTime             *prometheus.HistogramVec `description:"Microseconds of CPU time taken to complete a query"`
	QueryCount               *prometheus.GaugeVec     `description:"number of total queries"`
	QuerySuccessCount        *prometheus.GaugeVec     `description:"number of queries successfully processed"`
	QueryFailedCount         *prometheus.GaugeVec     `description:"number of failed queries"`
	QueryInterruptedCount    *prometheus.GaugeVec     `description:"number of queries interrupted due to cancellation or timeout"`
	SegmentScanPending       *prometheus.GaugeVec     `description:"number of segments in queue waiting to be scanned"`
}

// NewQueryHistoricalExporter returns a new historical exporter object
func NewQueryHistoricalExporter() *QueryHistoricalExporter {
	qh := &QueryHistoricalExporter{
		QueryTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_time",
			Help:      "milliseconds taken to complete a query",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"dataSource", "type", "hasFilters", "duration", "remoteAddress", "id", "numMetrics", "numComplexMetrics", "numDimensions", "threshold", "dimension"}),
		QuerySegmentTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_segment_time",
			Help:      "milliseconds taken to query individual segment. Includes time to page in the segment from disk",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"id", "status", "segment"}),
		QueryWaitTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_wait_time",
			Help:      "milliseconds spent waiting for a segment to be scanned",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"id", "segment"}),
		QuerySegmentAndCacheTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_segmentandcache_time",
			Help:      "milliseconds taken to query individual segment or hit the cache (if it is enabled on the Historical process)",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"id", "segment"}),
		QueryCPUTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_cpu_time",
			Help:      "Microseconds of CPU time taken to complete a query",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"dataSource", "type", "hasFilters", "duration", "remoteAddress", "id", "numMetrics", "numComplexMetrics", "numDimensions", "threshold", "dimension"}),
		QueryCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_count",
			Help:      "number of total queries",
		}, []string{}),
		QuerySuccessCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_success_count",
			Help:      "number of queries successfully processed",
		}, []string{}),
		QueryFailedCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{}),
		QueryInterruptedCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_interrupted_count",
			Help:      "number of queries interrupted due to cancellation or timeout",
		}, []string{}),
		SegmentScanPending: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "segment_scan_pending",
			Help:      "number of segments in queue waiting to be scanned",
		}, []string{}),
	}

	// register all prometheus metrics
	prometheus.MustRegister(qh.QueryTime)
	prometheus.MustRegister(qh.QuerySegmentTime)
	prometheus.MustRegister(qh.QueryWaitTime)
	prometheus.MustRegister(qh.QuerySegmentAndCacheTime)
	prometheus.MustRegister(qh.QueryCPUTime)
	prometheus.MustRegister(qh.QueryCount)
	prometheus.MustRegister(qh.QuerySuccessCount)
	prometheus.MustRegister(qh.QueryFailedCount)
	prometheus.MustRegister(qh.QueryInterruptedCount)
	prometheus.MustRegister(qh.SegmentScanPending)

	return qh
}

// SetQueryTime .
func (qh *QueryHistoricalExporter) SetQueryTime(labels map[string]string, val float64) {
	qh.QueryTime.With(labels).Observe(val)
}

// SetQuerySegmentTime .
func (qh *QueryHistoricalExporter) SetQuerySegmentTime(labels map[string]string, val float64) {
	qh.QuerySegmentTime.With(labels).Observe(val)
}

// SetQueryWaitTime .
func (qh *QueryHistoricalExporter) SetQueryWaitTime(labels map[string]string, val float64) {
	qh.QueryWaitTime.With(labels).Observe(val)
}

// SetQuerySegmentAndCacheTime .
func (qh *QueryHistoricalExporter) SetQuerySegmentAndCacheTime(labels map[string]string, val float64) {
	qh.QuerySegmentAndCacheTime.With(labels).Observe(val)
}

// SetQueryCPUTime .
func (qh *QueryHistoricalExporter) SetQueryCPUTime(labels map[string]string, val float64) {
	qh.QueryCPUTime.With(labels).Observe(val)
}

// SetQueryCount .
func (qh *QueryHistoricalExporter) SetQueryCount(val float64) {
	qh.QueryCount.WithLabelValues().Add(val)
}

// SetQuerySuccessCount .
func (qh *QueryHistoricalExporter) SetQuerySuccessCount(val float64) {
	qh.QuerySuccessCount.WithLabelValues().Add(val)
}

// SetQueryFailedCount .
func (qh *QueryHistoricalExporter) SetQueryFailedCount(val float64) {
	qh.QueryFailedCount.WithLabelValues().Add(val)
}

// SetQueryInterruptedCount .
func (qh *QueryHistoricalExporter) SetQueryInterruptedCount(val float64) {
	qh.QueryInterruptedCount.WithLabelValues().Add(val)
}

// SetSegmentScanPending .
func (qh *QueryHistoricalExporter) SetSegmentScanPending(val float64) {
	qh.SegmentScanPending.WithLabelValues().Add(val)
}
