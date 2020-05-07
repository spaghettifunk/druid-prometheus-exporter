package export

import "github.com/prometheus/client_golang/prometheus"

// SysExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type SysExporter struct {
	SwapFree       prometheus.Counter `description:"free swap"`
	SwapMax        prometheus.Counter `description:"max swap"`
	SwapPageIn     prometheus.Counter `description:"paged in swap"`
	SwapPageOut    prometheus.Counter `description:"paged out swap"`
	DiskWriteCount prometheus.Counter `description:"writes to disk"`
	DiskReadCount  prometheus.Counter `description:"reads from disk"`
	DiskWriteSize  prometheus.Counter `description:"bytes written to disk. It can be used to determine how much paging is occurring with regards to segments"`
	DiskReadSize   prometheus.Counter `description:"bytes read from disk. It can be used to determine how much paging is occurring with regards to segments"`
	NetWriteSize   prometheus.Counter `description:"bytes written to the network"`
	NetReadSize    prometheus.Counter `description:"bytes read from the network"`
	FSUsed         prometheus.Counter `description:"filesystem bytes used"`
	FSMax          prometheus.Counter `description:"filesystesm bytes max"`
	MemUsed        prometheus.Counter `description:"memory used"`
	MemMax         prometheus.Counter `description:"memory max"`
	StorageUsed    prometheus.Counter `description:"disk space used"`
	CPU            prometheus.Counter `description:"cpu used"`
}

// NewSysExporter returns a new Jetty exporter object
func NewSysExporter() *SysExporter {
	qj := &SysExporter{
		SwapFree: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "swap_free",
			Help:      "free swap",
			ConstLabels: prometheus.Labels{
				"sys": "swap-free",
			},
		}),
		SwapMax: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "swap_max",
			Help:      "max swap",
			ConstLabels: prometheus.Labels{
				"sys": "swap-max",
			},
		}),
		SwapPageIn: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "swap_pagein",
			Help:      "paged in swap",
			ConstLabels: prometheus.Labels{
				"sys": "swap-pagein",
			},
		}),
		SwapPageOut: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "swap_pageout",
			Help:      "paged out swap",
			ConstLabels: prometheus.Labels{
				"sys": "swap-pageout",
			},
		}),
		DiskWriteCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "disk_write_count",
			Help:      "writes to disk",
			ConstLabels: prometheus.Labels{
				"sys": "disk-write-count",
			},
		}),
		DiskReadCount: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "disk_read_count",
			Help:      "reads from disk",
			ConstLabels: prometheus.Labels{
				"sys": "disk-read-count",
			},
		}),
		DiskWriteSize: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "disk_write_size",
			Help:      "bytes written to disk. It can be used to determine how much paging is occurring with regards to segments",
			ConstLabels: prometheus.Labels{
				"sys": "disk-write-size",
			},
		}),
		DiskReadSize: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "disk_read_size",
			Help:      "bytes read from disk. It can be used to determine how much paging is occurring with regards to segments",
			ConstLabels: prometheus.Labels{
				"sys": "disk-read-size",
			},
		}),
		NetWriteSize: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "net_write_size",
			Help:      "bytes written to the network",
			ConstLabels: prometheus.Labels{
				"sys": "net-write-size",
			},
		}),
		NetReadSize: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "net_read_size",
			Help:      "bytes read from the network",
			ConstLabels: prometheus.Labels{
				"sys": "net-read-size",
			},
		}),
		FSUsed: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "fs_used",
			Help:      "filesystem bytes used",
			ConstLabels: prometheus.Labels{
				"sys": "fs-used",
			},
		}),
		FSMax: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "fs_max",
			Help:      "filesystesm bytes max",
			ConstLabels: prometheus.Labels{
				"sys": "fs-max",
			},
		}),
		MemUsed: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "mem_used",
			Help:      "memory used",
			ConstLabels: prometheus.Labels{
				"sys": "mem-used",
			},
		}),
		MemMax: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "mem_max",
			Help:      "memory max",
			ConstLabels: prometheus.Labels{
				"sys": "mem-max",
			},
		}),
		StorageUsed: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "storage_used",
			Help:      "disk space used",
			ConstLabels: prometheus.Labels{
				"sys": "storage-used",
			},
		}),
		CPU: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "druid",
			Subsystem: "sys",
			Name:      "cpu",
			Help:      "cpu used",
			ConstLabels: prometheus.Labels{
				"sys": "cpu",
			},
		}),
	}

	// register all the prometheus metrics
	prometheus.MustRegister(qj.SwapFree)
	prometheus.MustRegister(qj.SwapMax)
	prometheus.MustRegister(qj.SwapPageIn)
	prometheus.MustRegister(qj.SwapPageOut)
	prometheus.MustRegister(qj.DiskWriteCount)
	prometheus.MustRegister(qj.DiskReadCount)
	prometheus.MustRegister(qj.DiskWriteSize)
	prometheus.MustRegister(qj.DiskReadSize)
	prometheus.MustRegister(qj.NetWriteSize)
	prometheus.MustRegister(qj.NetReadSize)
	prometheus.MustRegister(qj.FSUsed)
	prometheus.MustRegister(qj.FSMax)
	prometheus.MustRegister(qj.MemUsed)
	prometheus.MustRegister(qj.MemMax)
	prometheus.MustRegister(qj.StorageUsed)
	prometheus.MustRegister(qj.CPU)

	return qj
}

// SetSwapFree .
func (se *SysExporter) SetSwapFree(val float64) {
}

// SetSwapMax .
func (se *SysExporter) SetSwapMax(val float64) {
}

// SetSwapPageIn .
func (se *SysExporter) SetSwapPageIn(val float64) {
}

// SetSwapPageOut .
func (se *SysExporter) SetSwapPageOut(val float64) {
}

// SetDiskWriteCount .
func (se *SysExporter) SetDiskWriteCount(val float64) {
}

// SetDiskReadCount .
func (se *SysExporter) SetDiskReadCount(val float64) {
}

// SetDiskWriteSize .
func (se *SysExporter) SetDiskWriteSize(val float64) {
}

// SetDiskReadSize .
func (se *SysExporter) SetDiskReadSize(val float64) {
}

// SetNetWriteSize .
func (se *SysExporter) SetNetWriteSize(val float64) {
}

// SetNetReadSize .
func (se *SysExporter) SetNetReadSize(val float64) {
}

// SetFSUsed .
func (se *SysExporter) SetFSUsed(val float64) {
}

// SetFSMax .
func (se *SysExporter) SetFSMax(val float64) {
}

// SetMemUsed .
func (se *SysExporter) SetMemUsed(val float64) {
}

// SetMemMax .
func (se *SysExporter) SetMemMax(val float64) {
}

// SetStorageUsed .
func (se *SysExporter) SetStorageUsed(val float64) {
}

// SetCPU .
func (se *SysExporter) SetCPU(val float64) {
}
