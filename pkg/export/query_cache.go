package export

import "github.com/prometheus/client_golang/prometheus"

// QueryCacheExporter contains all the Prometheus metrics that are possible to gather from the tranquillity service
type QueryCacheExporter struct {
	DeltaNumEntries   *prometheus.GaugeVec `description:"number of cache entries (since last emission)"`
	DeltaSizeBytes    *prometheus.GaugeVec `description:"size in bytes of cache entries (since last emission)"`
	DeltaHits         *prometheus.GaugeVec `description:"number of cache hits (since last emission)"`
	DeltaMisses       *prometheus.GaugeVec `description:"number of cache misses (since last emission)"`
	DeltaEvictions    *prometheus.GaugeVec `description:"number of cache evictions (since last emission)"`
	DeltaHitRate      prometheus.Counter   `description:"cache hit rate (since last emission)"`
	DeltaAverageBytes prometheus.Counter   `description:"average cache entry byte size (since last emission)"`
	DeltaTimeouts     *prometheus.GaugeVec `description:"number of cache timeouts (since last emission)"`
	DeltaErrors       *prometheus.GaugeVec `description:"number of cache errors (since last emission)"`
	DeltaPutOK        prometheus.Counter   `description:"number of new cache entries successfully cached (since last emission)"`
	DeltaPutError     prometheus.Counter   `description:"number of new cache entries that could not be cached due to errors (since last emission)"`
	DeltaPutOversized prometheus.Counter   `description:"number of potential new cache entries that were skipped due to being too large (based on druid.{broker,historical,realtime}.cache.maxEntrySize properties) (since last emission)"`
	TotalNumEntries   *prometheus.GaugeVec `description:"number of cache entries"`
	TotalSizeBytes    *prometheus.GaugeVec `description:"size in bytes of cache entries"`
	TotalHits         *prometheus.GaugeVec `description:"number of cache hits"`
	TotalMisses       *prometheus.GaugeVec `description:"number of cache misses"`
	TotalEvictions    *prometheus.GaugeVec `description:"number of cache evictions"`
	TotalHitRate      prometheus.Counter   `description:"cache hit rate"`
	TotalAverageBytes prometheus.Counter   `description:"average cache entry byte size"`
	TotalTimeouts     *prometheus.GaugeVec `description:"number of cache timeouts"`
	TotalErrors       *prometheus.GaugeVec `description:"number of cache errors"`
	TotalPutOK        prometheus.Counter   `description:"number of new cache entries successfully cached"`
	TotalPutError     prometheus.Counter   `description:"number of new cache entries that could not be cached due to errors"`
	TotalPutOversized prometheus.Counter   `description:"number of potential new cache entries that were skipped due to being too large (based on druid.{broker,historical,realtime}.cache.maxEntrySize properties)"`
}

// NewQueryCacheExporter returns a new Cache exporter object
func NewQueryCacheExporter() *QueryCacheExporter {
	qc := &QueryCacheExporter{
		DeltaAverageBytes: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "delta_average_bytes",
			Help:      "average cache entry byte size (since last emission)",
			ConstLabels: prometheus.Labels{
				"cache": "delta-average-bytes",
			},
		}),
		DeltaErrors: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "delta_errors",
			Help:      "number of cache errors (since last emission)",
		}, []string{}),
		DeltaEvictions: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "delta_evictions",
			Help:      "number of cache evictions (since last emission)",
		}, []string{}),
		DeltaHitRate: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "delta_hitrate",
			Help:      "cache hit rate (since last emission)",
			ConstLabels: prometheus.Labels{
				"cache": "delta-hitrate",
			},
		}),
		DeltaHits: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "delta_hits",
			Help:      "number of cache hits (since last emission)",
		}, []string{}),
		DeltaMisses: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "delta_misses",
			Help:      "number of cache misses (since last emission)",
		}, []string{}),
		DeltaNumEntries: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "delta_num_entries",
			Help:      "number of cache entries (since last emission)",
		}, []string{}),
		DeltaPutError: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "delta_put_error",
			Help:      "number of new cache entries that could not be cached due to errors (since last emission)",
			ConstLabels: prometheus.Labels{
				"cache": "delta-put-error",
			},
		}),
		DeltaPutOK: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "delta_put_ok",
			Help:      "number of new cache entries successfully cached (since last emission)",
			ConstLabels: prometheus.Labels{
				"cache": "delta-put-ok",
			},
		}),
		DeltaPutOversized: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "delta_put_oversized",
			Help:      "number of potential new cache entries that were skipped due to being too large (based on druid.{broker,historical,realtime}.cache.maxEntrySize properties)  (since last emission)",
			ConstLabels: prometheus.Labels{
				"cache": "delta-put-oversized",
			},
		}),
		DeltaSizeBytes: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "delta_size_bytes",
			Help:      "size in bytes of cache entries (since last emission)",
		}, []string{}),
		DeltaTimeouts: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "delta_timeouts",
			Help:      "number of cache timeouts (since last emission)",
		}, []string{}),
		TotalAverageBytes: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "total_average_bytes",
			Help:      "average cache entry byte size",
			ConstLabels: prometheus.Labels{
				"cache": "total-average-bytes",
			},
		}),
		TotalErrors: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "total_errors",
			Help:      "number of cache errors",
		}, []string{}),
		TotalEvictions: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "total_evictions",
			Help:      "number of cache evictions",
		}, []string{}),
		TotalHitRate: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "total_hitrate",
			Help:      "cache hit rate",
			ConstLabels: prometheus.Labels{
				"cache": "total-hitrate",
			},
		}),
		TotalHits: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "total_hits",
			Help:      "number of cache hits",
		}, []string{}),
		TotalMisses: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "total_misses",
			Help:      "number of cache misses",
		}, []string{}),
		TotalNumEntries: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "total_num_entries",
			Help:      "number of cache entries",
		}, []string{}),
		TotalPutError: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "total_put_error",
			Help:      "number of new cache entries that could not be cached due to errors",
			ConstLabels: prometheus.Labels{
				"cache": "total-put-error",
			},
		}),
		TotalPutOK: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "total_put_ok",
			Help:      "number of new cache entries successfully cached",
			ConstLabels: prometheus.Labels{
				"cache": "total-put-ok",
			},
		}),
		TotalSizeBytes: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "total_size_bytes",
			Help:      "size in bytes of cache entries",
		}, []string{}),
		TotalTimeouts: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "total_timeouts",
			Help:      "number of cache timeouts",
		}, []string{}),
		TotalPutOversized: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "cache",
			Name:      "total_put_oversized",
			Help:      "number of potential new cache entries that were skipped due to being too large (based on druid.{broker,historical,realtime}.cache.maxEntrySize properties)",
			ConstLabels: prometheus.Labels{
				"cache": "total-put-oversized",
			},
		}),
	}

	// register all the prometheus metrics
	prometheus.MustRegister(qc.DeltaNumEntries)
	prometheus.MustRegister(qc.DeltaSizeBytes)
	prometheus.MustRegister(qc.DeltaHits)
	prometheus.MustRegister(qc.DeltaMisses)
	prometheus.MustRegister(qc.DeltaEvictions)
	prometheus.MustRegister(qc.DeltaHitRate)
	prometheus.MustRegister(qc.DeltaAverageBytes)
	prometheus.MustRegister(qc.DeltaTimeouts)
	prometheus.MustRegister(qc.DeltaErrors)
	prometheus.MustRegister(qc.DeltaPutOK)
	prometheus.MustRegister(qc.DeltaPutError)
	prometheus.MustRegister(qc.DeltaPutOversized)
	prometheus.MustRegister(qc.TotalNumEntries)
	prometheus.MustRegister(qc.TotalSizeBytes)
	prometheus.MustRegister(qc.TotalHits)
	prometheus.MustRegister(qc.TotalMisses)
	prometheus.MustRegister(qc.TotalEvictions)
	prometheus.MustRegister(qc.TotalHitRate)
	prometheus.MustRegister(qc.TotalAverageBytes)
	prometheus.MustRegister(qc.TotalTimeouts)
	prometheus.MustRegister(qc.TotalErrors)
	prometheus.MustRegister(qc.TotalPutOK)
	prometheus.MustRegister(qc.TotalPutError)
	prometheus.MustRegister(qc.TotalPutOversized)

	return qc
}

// SetDeltaNumEntries .
func (qc *QueryCacheExporter) SetDeltaNumEntries(val float64) {
	qc.DeltaNumEntries.WithLabelValues().Add(val)
}

// SetDeltaSizeBytes .
func (qc *QueryCacheExporter) SetDeltaSizeBytes(val float64) {
	qc.DeltaSizeBytes.WithLabelValues().Add(val)
}

// SetDeltaHits .
func (qc *QueryCacheExporter) SetDeltaHits(val float64) {
	qc.DeltaHits.WithLabelValues().Add(val)
}

// SetDeltaMisses .
func (qc *QueryCacheExporter) SetDeltaMisses(val float64) {
	qc.DeltaMisses.WithLabelValues().Add(val)
}

// SetDeltaEvictions .
func (qc *QueryCacheExporter) SetDeltaEvictions(val float64) {
	qc.DeltaEvictions.WithLabelValues().Add(val)
}

// SetDeltaHitRate .
func (qc *QueryCacheExporter) SetDeltaHitRate(val float64) {
	qc.DeltaHitRate.Add(val)
}

// SetDeltaAverageBytes .
func (qc *QueryCacheExporter) SetDeltaAverageBytes(val float64) {
	qc.DeltaAverageBytes.Add(val)
}

// SetDeltaTimeouts .
func (qc *QueryCacheExporter) SetDeltaTimeouts(val float64) {
	qc.DeltaTimeouts.WithLabelValues().Add(val)
}

// SetDeltaErrors .
func (qc *QueryCacheExporter) SetDeltaErrors(val float64) {
	qc.DeltaErrors.WithLabelValues().Add(val)
}

// SetDeltaPutOK .
func (qc *QueryCacheExporter) SetDeltaPutOK(val float64) {
	qc.DeltaPutOK.Add(val)
}

// SetDeltaPutError .
func (qc *QueryCacheExporter) SetDeltaPutError(val float64) {
	qc.DeltaPutError.Add(val)
}

// SetDeltaPutOversized .
func (qc *QueryCacheExporter) SetDeltaPutOversized(val float64) {
	qc.DeltaPutOversized.Add(val)
}

// SetTotalNumEntries .
func (qc *QueryCacheExporter) SetTotalNumEntries(val float64) {
	qc.TotalNumEntries.WithLabelValues().Add(val)
}

// SetTotalSizeBytes .
func (qc *QueryCacheExporter) SetTotalSizeBytes(val float64) {
	qc.TotalSizeBytes.WithLabelValues().Add(val)
}

// SetTotalHits .
func (qc *QueryCacheExporter) SetTotalHits(val float64) {
	qc.TotalHits.WithLabelValues().Add(val)
}

// SetTotalMisses .
func (qc *QueryCacheExporter) SetTotalMisses(val float64) {
	qc.TotalMisses.WithLabelValues().Add(val)
}

// SetTotalEvictions .
func (qc *QueryCacheExporter) SetTotalEvictions(val float64) {
	qc.TotalEvictions.WithLabelValues().Add(val)
}

// SetTotalHitRate .
func (qc *QueryCacheExporter) SetTotalHitRate(val float64) {
	qc.TotalHitRate.Add(val)
}

// SetTotalAverageBytes .
func (qc *QueryCacheExporter) SetTotalAverageBytes(val float64) {
	qc.TotalAverageBytes.Add(val)
}

// SetTotalTimeouts .
func (qc *QueryCacheExporter) SetTotalTimeouts(val float64) {
	qc.TotalTimeouts.WithLabelValues().Add(val)
}

// SetTotalErrors .
func (qc *QueryCacheExporter) SetTotalErrors(val float64) {
	qc.TotalErrors.WithLabelValues().Add(val)
}

// SetTotalPutOK .
func (qc *QueryCacheExporter) SetTotalPutOK(val float64) {
	qc.TotalPutOK.Add(val)
}

// SetTotalPutError .
func (qc *QueryCacheExporter) SetTotalPutError(val float64) {
	qc.TotalPutError.Add(val)
}

// SetTotalPutOversized .
func (qc *QueryCacheExporter) SetTotalPutOversized(val float64) {
	qc.TotalPutOversized.Add(val)
}
