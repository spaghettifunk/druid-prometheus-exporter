package export

import "github.com/prometheus/client_golang/prometheus"

// HealthJVMExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type HealthJVMExporter struct {
	PoolCommitted      prometheus.Counter `description:"committed pool"`
	PoolInit           prometheus.Counter `description:"initial pool"`
	PoolMax            prometheus.Counter `description:"max pool"`
	PoolUsed           prometheus.Counter `description:"pool used"`
	BufferpoolCount    prometheus.Counter `description:"bufferpool count"`
	BufferpoolUsed     prometheus.Counter `description:"bufferpool used"`
	BufferpoolCapacity prometheus.Counter `description:"bufferpool capacity"`
	MemInit            prometheus.Counter `description:"initial memory"`
	MemMax             prometheus.Counter `description:"max memory"`
	MemUsed            prometheus.Counter `description:"used memory"`
	MemCommitted       prometheus.Counter `description:"committed memory"`
	GCCount            prometheus.Counter `description:"garbage collection count"`
	GCCPU              prometheus.Counter `description:"count of CPU time in Nanoseconds spent on garbage collection. Note: jvm/gc/cpu represents the total time over multiple GC cycles; divide by jvm/gc/count to get the mean GC time per cycle"`
}

// NewHealthJVMExporter returns a new Jetty exporter object
func NewHealthJVMExporter() *HealthJVMExporter {
	qj := &HealthJVMExporter{
		PoolCommitted: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "pool_committed",
			Help:      "committed pool",
			ConstLabels: prometheus.Labels{
				"jvm": "pool-committed",
			},
		}),
		PoolInit: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "pool_init",
			Help:      "initial pool",
			ConstLabels: prometheus.Labels{
				"jvm": "pool-init",
			},
		}),
		PoolMax: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "max_pool",
			Help:      "max pool",
			ConstLabels: prometheus.Labels{
				"jvm": "max-pool",
			},
		}),
		PoolUsed: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "pool_used",
			Help:      "pool used",
			ConstLabels: prometheus.Labels{
				"jvm": "pool-used",
			},
		}),
		BufferpoolCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "bufferpool_count",
			Help:      "bufferpool count",
			ConstLabels: prometheus.Labels{
				"jvm": "bufferpool-count",
			},
		}),
		BufferpoolUsed: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "bufferpool_used",
			Help:      "bufferpool used",
			ConstLabels: prometheus.Labels{
				"jvm": "bufferpool-used",
			},
		}),
		BufferpoolCapacity: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "bufferpool_capacity",
			Help:      "bufferpool capacity",
			ConstLabels: prometheus.Labels{
				"jvm": "bufferpool-used",
			},
		}),
		MemInit: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "mem_init",
			Help:      "initial memory",
			ConstLabels: prometheus.Labels{
				"jvm": "mem-init",
			},
		}),
		MemMax: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "mem_max",
			Help:      "max memory",
			ConstLabels: prometheus.Labels{
				"jvm": "mem-max",
			},
		}),
		MemUsed: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "mem_used",
			Help:      "used memory",
			ConstLabels: prometheus.Labels{
				"jvm": "mem-used",
			},
		}),
		MemCommitted: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "mem_committed",
			Help:      "committed memory",
			ConstLabels: prometheus.Labels{
				"jvm": "mem-committed",
			},
		}),
		GCCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "gc_count",
			Help:      "garbage collection count",
			ConstLabels: prometheus.Labels{
				"jvm": "gc-count",
			},
		}),
		GCCPU: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "gc_cpu",
			Help:      "count of CPU time in Nanoseconds spent on garbage collection. Note: jvm/gc/cpu represents the total time over multiple GC cycles; divide by jvm/gc/count to get the mean GC time per cycle",
			ConstLabels: prometheus.Labels{
				"jvm": "gc-cpu",
			},
		}),
	}

	// register all the prometheus metrics
	prometheus.MustRegister(qj.PoolCommitted)
	prometheus.MustRegister(qj.PoolInit)
	prometheus.MustRegister(qj.PoolMax)
	prometheus.MustRegister(qj.PoolUsed)
	prometheus.MustRegister(qj.BufferpoolCount)
	prometheus.MustRegister(qj.BufferpoolUsed)
	prometheus.MustRegister(qj.BufferpoolCapacity)
	prometheus.MustRegister(qj.MemInit)
	prometheus.MustRegister(qj.MemMax)
	prometheus.MustRegister(qj.MemUsed)
	prometheus.MustRegister(qj.MemCommitted)
	prometheus.MustRegister(qj.GCCount)
	prometheus.MustRegister(qj.GCCPU)

	return qj
}

// SetPoolCommited .
func (hj *HealthJVMExporter) SetPoolCommited(val float64) {
}

// SetPoolInit .
func (hj *HealthJVMExporter) SetPoolInit(val float64) {
}

// SetPoolMax .
func (hj *HealthJVMExporter) SetPoolMax(val float64) {
}

// SetPoolUsed .
func (hj *HealthJVMExporter) SetPoolUsed(val float64) {
}

// SetBufferpoolCount .
func (hj *HealthJVMExporter) SetBufferpoolCount(val float64) {
}

// SetBufferpoolUsed .
func (hj *HealthJVMExporter) SetBufferpoolUsed(val float64) {
}

// SetBufferpoolCapacity .
func (hj *HealthJVMExporter) SetBufferpoolCapacity(val float64) {
}

// SetMemInit .
func (hj *HealthJVMExporter) SetMemInit(val float64) {
}

// SetMemMax .
func (hj *HealthJVMExporter) SetMemMax(val float64) {
}

// SetMemUsed .
func (hj *HealthJVMExporter) SetMemUsed(val float64) {
}

// SetMemCommitted .
func (hj *HealthJVMExporter) SetMemCommitted(val float64) {
}

// SetGCCount .
func (hj *HealthJVMExporter) SetGCCount(val float64) {
}

// SetGCCPU .
func (hj *HealthJVMExporter) SetGCCPU(val float64) {
}
