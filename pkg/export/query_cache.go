package export

import "github.com/prometheus/client_golang/prometheus"

// QueryCacheExporter contains all the Prometheus metrics that are possible to gather from the tranquillity service
type QueryCacheExporter struct {
	CacheTotalPutOversized prometheus.Counter `description:"number of potential new cache entries that were skipped due to being too large (based on druid.{broker,historical,realtime}.cache.maxEntrySize properties)"`
}

// NewQueryCacheExporter returns a new Cache exporter object
func NewQueryCacheExporter() *QueryCacheExporter {
	qc := &QueryCacheExporter{
		CacheTotalPutOversized: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "total_put_oversized",
			Help:      "number of potential new cache entries that were skipped due to being too large (based on druid.{broker,historical,realtime}.cache.maxEntrySize properties)",
			ConstLabels: prometheus.Labels{
				"cache": "total-put-oversized",
			},
		}),
	}

	return qc
}

// SetCacheTotalPutOversized .
func (qc *QueryCacheExporter) SetCacheTotalPutOversized(val float64) {

}
