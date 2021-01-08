package feed

import (
	"time"
)

// Feed represents the structure that will contain all the values of each feed
// without disticion of service or metric
type Feed struct {
	Timestamp         time.Time   `json:"timestamp"`
	Metric            string      `json:"metric"`
	Service           string      `json:"service"`
	Host              string      `json:"host"`
	Value             interface{} `json:"value"`
	Feed              string      `json:"feed"`
	DataSource        string      `json:"dataSource"`
	ID                string      `json:"id"`
	NativeQueryIDs    string      `json:"nativeQueryIds"`
	RemoteAddess      string      `json:"remoteAddress"`
	Success           string      `json:"success"`
	FSDevName         string      `json:"fsDevName"`
	FSDirName         string      `json:"fsDirName"`
	FSOptions         []string    `json:"fsOptions"`
	FSSysTypeName     string      `json:"fsSysTypeName"`
	FSTypeName        string      `json:"fsTypeName"`
	NetAddress        string      `json:"netAddress"`
	NetHwAddr         string      `json:"netHwaddr"`
	NetName           string      `json:"netName"`
	CPUName           string      `json:"cpuName"`
	CPUTime           string      `json:"cpuTime"`
	MemKind           string      `json:"memKind"`
	PoolKind          string      `json:"poolKind"`
	PoolName          string      `json:"poolName"`
	BufferPoolName    string      `json:"bufferpoolName"`
	GCGen             []string    `json:"gcGen"`
	GCName            []string    `json:"gcName"`
	GCGenSpaceName    string      `json:"gcGenSpaceName"`
	Tier              string      `json:"tier"`
	Server            string      `json:"server"`
	Type              string      `json:"type"`
	HasFilters        string      `json:"hasFilters"`
	Duration          string      `json:"duration"`
	NumComplexMetrics string      `json:"numComplexMetrics"`
	NumDimensions     string      `json:"numDimensions"`
	Threshold         string      `json:"threshold"`
	Dimension         string      `json:"dimension"`
	Status            string      `json:"status"`
	Segment           string      `json:"segment"`
}

// Service is the interface that will evaluate the feed based on the
// service name
type Service interface {
	// Evaluate decides which prometheus export to call
	Evaluate(Feed)
}
