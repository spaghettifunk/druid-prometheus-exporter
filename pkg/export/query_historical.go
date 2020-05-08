package export

import "github.com/prometheus/client_golang/prometheus"

// QueryHistoricalExporter contains all the Prometheus metrics that are possible to gather from the historicals
type QueryHistoricalExporter struct {
	QueryTime                prometheus.Counter `description:"milliseconds taken to complete a query"`
	QuerySegmentTime         prometheus.Counter `description:"milliseconds taken to query individual segment. Includes time to page in the segment from disk"`
	QueryWaitTime            prometheus.Counter `description:"milliseconds spent waiting for a segment to be scanned"`
	QuerySegmentAndCacheTime prometheus.Counter `description:"milliseconds taken to query individual segment or hit the cache (if it is enabled on the Historical process)"`
	QueryCPUTime             prometheus.Counter `description:"Microseconds of CPU time taken to complete a query"`
	QueryCount               prometheus.Counter `description:"number of total queries"`
	QuerySuccessCount        prometheus.Counter `description:"number of queries successfully processed"`
	QueryFailedCount         prometheus.Counter `description:"number of failed queries"`
	QueryInterruptedCount    prometheus.Counter `description:"number of queries interrupted due to cancellation or timeout"`
	SegmentScanPending       prometheus.Counter `description:"number of segments in queue waiting to be scanned"`
}

// NewQueryHistoricalExporter returns a new historical exporter object
func NewQueryHistoricalExporter() *QueryHistoricalExporter {
	qh := &QueryHistoricalExporter{
		QueryTime: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_time",
			Help:      "milliseconds taken to complete a query",
			ConstLabels: prometheus.Labels{
				"historical": "query-time",
			},
		}),
		QuerySegmentTime: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_segment_time",
			Help:      "milliseconds taken to query individual segment. Includes time to page in the segment from disk",
			ConstLabels: prometheus.Labels{
				"historical": "query-segment-time",
			},
		}),
		QueryWaitTime: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_wait_time",
			Help:      "milliseconds spent waiting for a segment to be scanned",
			ConstLabels: prometheus.Labels{
				"historical": "query-wait-time",
			},
		}),
		QuerySegmentAndCacheTime: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_segmentandcache_time",
			Help:      "milliseconds taken to query individual segment or hit the cache (if it is enabled on the Historical process)",
			ConstLabels: prometheus.Labels{
				"historical": "query-segmentandcache-time",
			},
		}),
		QueryCPUTime: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_cpu_time",
			Help:      "Microseconds of CPU time taken to complete a query",
			ConstLabels: prometheus.Labels{
				"historical": "query-cpu-time",
			},
		}),
		QueryCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_count",
			Help:      "number of total queries",
			ConstLabels: prometheus.Labels{
				"historical": "query-count",
			},
		}),
		QuerySuccessCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_success_count",
			Help:      "number of queries successfully processed",
			ConstLabels: prometheus.Labels{
				"historical": "query-success-count",
			},
		}),
		QueryFailedCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
			ConstLabels: prometheus.Labels{
				"historical": "query-failed-count",
			},
		}),
		QueryInterruptedCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "historical",
			Name:      "query_interrupted_count",
			Help:      "number of queries interrupted due to cancellation or timeout",
			ConstLabels: prometheus.Labels{
				"historical": "query-interrupted-count",
			},
		}),
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
func (qh *QueryHistoricalExporter) SetQueryTime(val float64) {

}

// SetQuerySegmentTime .
func (qh *QueryHistoricalExporter) SetQuerySegmentTime(val float64) {

}

// SetQueryWaitTime .
func (qh *QueryHistoricalExporter) SetQueryWaitTime(val float64) {

}

// SetQuerySegmentAndCacheTime .
func (qh *QueryHistoricalExporter) SetQuerySegmentAndCacheTime(val float64) {

}

// SetQueryCPUTime .
func (qh *QueryHistoricalExporter) SetQueryCPUTime(val float64) {

}

// SetQueryCount .
func (qh *QueryHistoricalExporter) SetQueryCount(val float64) {

}

// SetQuerySuccessCount .
func (qh *QueryHistoricalExporter) SetQuerySuccessCount(val float64) {

}

// SetQueryFailedCount .
func (qh *QueryHistoricalExporter) SetQueryFailedCount(val float64) {

}

// SetQueryInterruptedCount .
func (qh *QueryHistoricalExporter) SetQueryInterruptedCount(val float64) {

}

// SetSegmentScanPending .
func (qh *QueryHistoricalExporter) SetSegmentScanPending(val float64) {

}
