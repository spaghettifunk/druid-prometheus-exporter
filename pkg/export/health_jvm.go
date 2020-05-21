package export

import "github.com/prometheus/client_golang/prometheus"

// HealthJVMExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type HealthJVMExporter struct {
	PoolCommitted      *prometheus.GaugeVec `description:"committed pool"`
	PoolInit           *prometheus.GaugeVec `description:"initial pool"`
	PoolMax            *prometheus.GaugeVec `description:"max pool"`
	PoolUsed           *prometheus.GaugeVec `description:"pool used"`
	BufferpoolCount    *prometheus.GaugeVec `description:"bufferpool count"`
	BufferpoolUsed     *prometheus.GaugeVec `description:"bufferpool used"`
	BufferpoolCapacity *prometheus.GaugeVec `description:"bufferpool capacity"`
	MemInit            *prometheus.GaugeVec `description:"initial memory"`
	MemMax             *prometheus.GaugeVec `description:"max memory"`
	MemUsed            *prometheus.GaugeVec `description:"used memory"`
	MemCommitted       *prometheus.GaugeVec `description:"committed memory"`
	GCCount            *prometheus.GaugeVec `description:"garbage collection count"`
	GCCPU              *prometheus.GaugeVec `description:"count of CPU time in Nanoseconds spent on garbage collection. Note: jvm/gc/cpu represents the total time over multiple GC cycles; divide by jvm/gc/count to get the mean GC time per cycle"`
}

// NewHealthJVMExporter returns a new Jetty exporter object
func NewHealthJVMExporter() *HealthJVMExporter {
	hj := &HealthJVMExporter{
		PoolCommitted: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "pool_committed",
			Help:      "committed pool",
		}, []string{"poolKind", "poolName"}),
		PoolInit: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "pool_init",
			Help:      "initial pool",
		}, []string{"poolKind", "poolName"}),
		PoolMax: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "max_pool",
			Help:      "max pool",
		}, []string{"poolKind", "poolName"}),
		PoolUsed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "pool_used",
			Help:      "pool used",
		}, []string{"poolKind", "poolName"}),
		BufferpoolCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "bufferpool_count",
			Help:      "bufferpool count",
		}, []string{"bufferPoolName"}),
		BufferpoolUsed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "bufferpool_used",
			Help:      "bufferpool used",
		}, []string{"bufferPoolName"}),
		BufferpoolCapacity: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "bufferpool_capacity",
			Help:      "bufferpool capacity",
		}, []string{"bufferPoolName"}),
		MemInit: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "mem_init",
			Help:      "initial memory",
		}, []string{"memKind"}),
		MemMax: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "mem_max",
			Help:      "max memory",
		}, []string{"memKind"}),
		MemUsed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "mem_used",
			Help:      "used memory",
		}, []string{"memKind"}),
		MemCommitted: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "mem_committed",
			Help:      "committed memory",
		}, []string{"memKind"}),
		GCCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "gc_count",
			Help:      "garbage collection count",
		}, []string{"gcName", "gcGen"}),
		GCCPU: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "jvm",
			Name:      "gc_cpu",
			Help:      "count of CPU time in Nanoseconds spent on garbage collection. Note: jvm/gc/cpu represents the total time over multiple GC cycles; divide by jvm/gc/count to get the mean GC time per cycle",
		}, []string{"gcName", "gcGen"}),
	}

	// register all the prometheus metrics
	prometheus.MustRegister(hj.PoolCommitted)
	prometheus.MustRegister(hj.PoolInit)
	prometheus.MustRegister(hj.PoolMax)
	prometheus.MustRegister(hj.PoolUsed)
	prometheus.MustRegister(hj.BufferpoolCount)
	prometheus.MustRegister(hj.BufferpoolUsed)
	prometheus.MustRegister(hj.BufferpoolCapacity)
	prometheus.MustRegister(hj.MemInit)
	prometheus.MustRegister(hj.MemMax)
	prometheus.MustRegister(hj.MemUsed)
	prometheus.MustRegister(hj.MemCommitted)
	prometheus.MustRegister(hj.GCCount)
	prometheus.MustRegister(hj.GCCPU)

	return hj
}

// SetPoolCommited .
func (hj *HealthJVMExporter) SetPoolCommited(labels map[string]string, val float64) {
	hj.PoolCommitted.With(labels).Add(val)
}

// SetPoolInit .
func (hj *HealthJVMExporter) SetPoolInit(labels map[string]string, val float64) {
	hj.PoolInit.With(labels).Add(val)
}

// SetPoolMax .
func (hj *HealthJVMExporter) SetPoolMax(labels map[string]string, val float64) {
	hj.PoolMax.With(labels).Add(val)
}

// SetPoolUsed .
func (hj *HealthJVMExporter) SetPoolUsed(labels map[string]string, val float64) {
	hj.PoolUsed.With(labels).Add(val)
}

// SetBufferpoolCount .
func (hj *HealthJVMExporter) SetBufferpoolCount(labels map[string]string, val float64) {
	hj.BufferpoolCount.With(labels).Add(val)
}

// SetBufferpoolUsed .
func (hj *HealthJVMExporter) SetBufferpoolUsed(labels map[string]string, val float64) {
	hj.BufferpoolUsed.With(labels).Add(val)
}

// SetBufferpoolCapacity .
func (hj *HealthJVMExporter) SetBufferpoolCapacity(labels map[string]string, val float64) {
	hj.BufferpoolCapacity.With(labels).Add(val)
}

// SetMemInit .
func (hj *HealthJVMExporter) SetMemInit(labels map[string]string, val float64) {
	hj.MemInit.With(labels).Add(val)
}

// SetMemMax .
func (hj *HealthJVMExporter) SetMemMax(labels map[string]string, val float64) {
	hj.MemMax.With(labels).Add(val)
}

// SetMemUsed .
func (hj *HealthJVMExporter) SetMemUsed(labels map[string]string, val float64) {
	hj.MemUsed.With(labels).Add(val)
}

// SetMemCommitted .
func (hj *HealthJVMExporter) SetMemCommitted(labels map[string]string, val float64) {
	hj.MemCommitted.With(labels).Add(val)
}

// SetGCCount .
func (hj *HealthJVMExporter) SetGCCount(labels map[string]string, val float64) {
	hj.GCCount.With(labels).Add(val)
}

// SetGCCPU .
func (hj *HealthJVMExporter) SetGCCPU(labels map[string]string, val float64) {
	hj.GCCPU.With(labels).Add(val)
}
