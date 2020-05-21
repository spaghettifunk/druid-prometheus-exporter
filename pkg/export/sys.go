package export

import "github.com/prometheus/client_golang/prometheus"

// SysExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type SysExporter struct {
	SwapFree       *prometheus.GaugeVec `description:"free swap"`
	SwapMax        *prometheus.GaugeVec `description:"max swap"`
	SwapPageIn     *prometheus.GaugeVec `description:"paged in swap"`
	SwapPageOut    *prometheus.GaugeVec `description:"paged out swap"`
	DiskWriteCount *prometheus.GaugeVec `description:"writes to disk"`
	DiskReadCount  *prometheus.GaugeVec `description:"reads from disk"`
	DiskWriteSize  *prometheus.GaugeVec `description:"bytes written to disk. It can be used to determine how much paging is occurring with regards to segments"`
	DiskReadSize   *prometheus.GaugeVec `description:"bytes read from disk. It can be used to determine how much paging is occurring with regards to segments"`
	NetWriteSize   *prometheus.GaugeVec `description:"bytes written to the network"`
	NetReadSize    *prometheus.GaugeVec `description:"bytes read from the network"`
	FSUsed         *prometheus.GaugeVec `description:"filesystem bytes used"`
	FSMax          *prometheus.GaugeVec `description:"filesystesm bytes max"`
	MemUsed        *prometheus.GaugeVec `description:"memory used"`
	MemMax         *prometheus.GaugeVec `description:"memory max"`
	StorageUsed    *prometheus.GaugeVec `description:"disk space used"`
	CPU            *prometheus.GaugeVec `description:"cpu used"`
}

// NewSysExporter returns a new Jetty exporter object
func NewSysExporter() *SysExporter {
	se := &SysExporter{
		SwapFree: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{}),
		SwapMax: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{}),
		SwapPageIn: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{}),
		SwapPageOut: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{}),
		DiskWriteCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{"fsDevName", "fsDirName", "fsTypeName", "fsSysTypeName", "fsOptions"}),
		DiskReadCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{"fsDevName", "fsDirName", "fsTypeName", "fsSysTypeName", "fsOptions"}),
		DiskWriteSize: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{"fsDevName", "fsDirName", "fsTypeName", "fsSysTypeName", "fsOptions"}),
		DiskReadSize: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{"fsDevName", "fsDirName", "fsTypeName", "fsSysTypeName", "fsOptions"}),
		NetWriteSize: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{"netName", "netAddress", "netHwaddr"}),
		NetReadSize: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{"netName", "netAddress", "netHwaddr"}),
		FSUsed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{"fsDevName", "fsDirName", "fsTypeName", "fsSysTypeName", "fsOptions"}),
		FSMax: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{"fsDevName", "fsDirName", "fsTypeName", "fsSysTypeName", "fsOptions"}),
		MemUsed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{}),
		MemMax: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{}),
		StorageUsed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{"fsDirName"}),
		CPU: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{"cpuName", "cpuTime"}),
	}

	// register all the prometheus metrics
	prometheus.MustRegister(se.SwapFree)
	prometheus.MustRegister(se.SwapMax)
	prometheus.MustRegister(se.SwapPageIn)
	prometheus.MustRegister(se.SwapPageOut)
	prometheus.MustRegister(se.DiskWriteCount)
	prometheus.MustRegister(se.DiskReadCount)
	prometheus.MustRegister(se.DiskWriteSize)
	prometheus.MustRegister(se.DiskReadSize)
	prometheus.MustRegister(se.NetWriteSize)
	prometheus.MustRegister(se.NetReadSize)
	prometheus.MustRegister(se.FSUsed)
	prometheus.MustRegister(se.FSMax)
	prometheus.MustRegister(se.MemUsed)
	prometheus.MustRegister(se.MemMax)
	prometheus.MustRegister(se.StorageUsed)
	prometheus.MustRegister(se.CPU)

	return se
}

// SetSwapFree .
func (se *SysExporter) SetSwapFree(val float64) {
	se.SwapFree.WithLabelValues().Add(val)
}

// SetSwapMax .
func (se *SysExporter) SetSwapMax(val float64) {
	se.SwapMax.WithLabelValues().Add(val)
}

// SetSwapPageIn .
func (se *SysExporter) SetSwapPageIn(val float64) {
	se.SwapPageIn.WithLabelValues().Add(val)
}

// SetSwapPageOut .
func (se *SysExporter) SetSwapPageOut(val float64) {
	se.SwapPageOut.WithLabelValues().Add(val)
}

// SetDiskWriteCount .
func (se *SysExporter) SetDiskWriteCount(labels map[string]string, val float64) {
	se.DiskWriteCount.With(labels).Add(val)
}

// SetDiskReadCount .
func (se *SysExporter) SetDiskReadCount(labels map[string]string, val float64) {
	se.DiskReadCount.With(labels).Add(val)
}

// SetDiskWriteSize .
func (se *SysExporter) SetDiskWriteSize(labels map[string]string, val float64) {
	se.DiskWriteSize.With(labels).Add(val)
}

// SetDiskReadSize .
func (se *SysExporter) SetDiskReadSize(labels map[string]string, val float64) {
	se.DiskReadSize.With(labels).Add(val)
}

// SetNetWriteSize .
func (se *SysExporter) SetNetWriteSize(labels map[string]string, val float64) {
	se.NetWriteSize.With(labels).Add(val)
}

// SetNetReadSize .
func (se *SysExporter) SetNetReadSize(labels map[string]string, val float64) {
	se.NetReadSize.With(labels).Add(val)
}

// SetFSUsed .
func (se *SysExporter) SetFSUsed(labels map[string]string, val float64) {
	se.FSUsed.With(labels).Add(val)
}

// SetFSMax .
func (se *SysExporter) SetFSMax(labels map[string]string, val float64) {
	se.FSMax.With(labels).Add(val)
}

// SetMemUsed .
func (se *SysExporter) SetMemUsed(val float64) {
	se.MemUsed.WithLabelValues().Add(val)
}

// SetMemMax .
func (se *SysExporter) SetMemMax(val float64) {
	se.MemMax.WithLabelValues().Add(val)
}

// SetStorageUsed .
func (se *SysExporter) SetStorageUsed(fsDirName string, val float64) {
	se.StorageUsed.With(prometheus.Labels{"fsDirName": fsDirName}).Add(val)
}

// SetCPU .
func (se *SysExporter) SetCPU(labels map[string]string, val float64) {
	se.CPU.With(labels).Add(val)
}
