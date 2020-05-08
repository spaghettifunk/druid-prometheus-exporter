package feed

import "github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"

// Historical initiate the metrics for the historical feeds
type Historical struct {
	QueryHistoricalExporter  *export.QueryHistoricalExporter
	SQLExporter              *export.SQLExporter
	QueryCacheExporter       *export.QueryCacheExporter
	HealthJVMExporter        *export.HealthJVMExporter
	QueryJettyExporter       *export.QueryJettyExporter
	SysExporter              *export.SysExporter
	HealthHistoricalExporter *export.HealthHistoricalExporter
}

// Evaluate decides whicfh prometheus export to call
func (hf *Historical) Evaluate(f Feed) {
	switch f.Metric {
	case "query/time":
		val, _ := f.Value.(float64)
		hf.QueryHistoricalExporter.SetQueryTime(float64(val))
		break
	case "query/segment/time":
		val, _ := f.Value.(float64)
		hf.QueryHistoricalExporter.SetQuerySegmentTime(float64(val))
		break
	case "query/wait/time":
		val, _ := f.Value.(float64)
		hf.QueryHistoricalExporter.SetQueryWaitTime(float64(val))
		break
	case "query/segmentAndCache/time":
		val, _ := f.Value.(float64)
		hf.QueryHistoricalExporter.SetQuerySegmentAndCacheTime(float64(val))
		break
	case "query/cpu/time":
		val, _ := f.Value.(float64)
		hf.QueryHistoricalExporter.SetQueryCPUTime(float64(val))
		break
	case "query/count":
		val, _ := f.Value.(float64)
		hf.QueryHistoricalExporter.SetQueryCount(float64(val))
		break
	case "query/success/count":
		val, _ := f.Value.(float64)
		hf.QueryHistoricalExporter.SetQuerySuccessCount(float64(val))
		break
	case "query/failed/count":
		val, _ := f.Value.(float64)
		hf.QueryHistoricalExporter.SetQueryFailedCount(float64(val))
		break
	case "query/interrupted/count":
		val, _ := f.Value.(float64)
		hf.QueryHistoricalExporter.SetQueryInterruptedCount(float64(val))
		break
	case "segment/scan/pending":
		val, _ := f.Value.(float64)
		hf.QueryHistoricalExporter.SetSegmentScanPending(float64(val))
		break
	case "sqlQuery/time":
		val, _ := f.Value.(float64)
		hf.SQLExporter.SetSQLQueryTime(float64(val))
		break
	case "sqlQuery/bytes":
		val, _ := f.Value.(float64)
		hf.SQLExporter.SetSQLQueryBytes(float64(val))
		break
	case "query/cache/delta/numEntries":
		break
	case "query/cache/delta/sizeBytes":
		break
	case "query/cache/delta/hits":
		break
	case "query/cache/delta/misses":
		break
	case "query/cache/delta/evictions":
		break
	case "query/cache/delta/hitRate":
		break
	case "query/cache/delta/averageBytes":
		break
	case "query/cache/delta/timeouts":
		break
	case "query/cache/delta/errors":
		break
	case "query/cache/delta/put/ok":
		break
	case "query/cache/delta/put/error":
		break
	case "query/cache/delta/put/oversized":
		break
	case "query/cache/total/numEntries":
		break
	case "query/cache/total/sizeBytes":
		break
	case "query/cache/total/hits":
		break
	case "query/cache/total/misses":
		break
	case "query/cache/total/evictions":
		break
	case "query/cache/total/hitRate":
		break
	case "query/cache/total/averageBytes":
		break
	case "query/cache/total/timeouts":
		break
	case "query/cache/total/errors":
		break
	case "query/cache/total/put/ok":
		break
	case "query/cache/total/put/error":
		break
	case "query/cache/total/put/oversized":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetTotalPutOversized(float64(val))
		break
	case "jvm/pool/committed":
		val, _ := f.Value.(float64)
		hf.HealthJVMExporter.SetPoolCommited(float64(val))
		break
	case "jvm/pool/init":
		val, _ := f.Value.(float64)
		hf.HealthJVMExporter.SetPoolInit(float64(val))
		break
	case "jvm/pool/max":
		val, _ := f.Value.(float64)
		hf.HealthJVMExporter.SetPoolMax(float64(val))
		break
	case "jvm/pool/used":
		val, _ := f.Value.(float64)
		hf.HealthJVMExporter.SetPoolUsed(float64(val))
		break
	case "jvm/bufferpool/count":
		val, _ := f.Value.(float64)
		hf.HealthJVMExporter.SetBufferpoolCount(float64(val))
		break
	case "jvm/bufferpool/used":
		val, _ := f.Value.(float64)
		hf.HealthJVMExporter.SetBufferpoolUsed(float64(val))
		break
	case "jvm/bufferpool/capacity":
		val, _ := f.Value.(float64)
		hf.HealthJVMExporter.SetBufferpoolCapacity(float64(val))
		break
	case "jvm/mem/init":
		val, _ := f.Value.(float64)
		hf.HealthJVMExporter.SetMemInit(float64(val))
		break
	case "jvm/mem/max":
		val, _ := f.Value.(float64)
		hf.HealthJVMExporter.SetMemMax(float64(val))
		break
	case "jvm/mem/used":
		val, _ := f.Value.(float64)
		hf.HealthJVMExporter.SetMemUsed(float64(val))
		break
	case "jvm/mem/committed":
		val, _ := f.Value.(float64)
		hf.HealthJVMExporter.SetMemCommitted(float64(val))
		break
	case "jvm/gc/count":
		val, _ := f.Value.(float64)
		hf.HealthJVMExporter.SetGCCount(float64(val))
		break
	case "jvm/gc/cpu":
		val, _ := f.Value.(float64)
		hf.HealthJVMExporter.SetGCCPU(float64(val))
		break
	case "jetty/numOpenConnections":
		val, _ := f.Value.(float64)
		hf.QueryJettyExporter.SetNumOpenConnections(float64(val))
		break
	case "sys/swap/free":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetSwapFree(float64(val))
		break
	case "sys/swap/max":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetSwapMax(float64(val))
		break
	case "sys/swap/pageIn":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetSwapPageIn(float64(val))
		break
	case "sys/swap/pageOut":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetSwapPageOut(float64(val))
		break
	case "sys/disk/write/count":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetDiskWriteCount(float64(val))
		break
	case "sys/disk/read/count":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetDiskReadCount(float64(val))
		break
	case "sys/disk/write/size":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetDiskWriteSize(float64(val))
		break
	case "sys/disk/read/size":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetDiskReadSize(float64(val))
		break
	case "sys/net/read/size	":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetNetReadSize(float64(val))
		break
	case "sys/mem/max":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetMemMax(float64(val))
		break
	case "sys/mem/used":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetMemUsed(float64(val))
		break
	case "sys/fs/max":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetFSMax(float64(val))
		break
	case "sys/fs/used":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetFSUsed(float64(val))
		break
	case "sys/net/write/size":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetNetWriteSize(float64(val))
		break
	case "sys/net/read/size":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetNetReadSize(float64(val))
		break
	case "sys/storage/used":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetStorageUsed(float64(val))
		break
	case "sys/cpu":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetCPU(float64(val))
		break
	case "segment/max":
		val, _ := f.Value.(float64)
		hf.HealthHistoricalExporter.SetSegmentMax(float64(val))
		break
	default:
		break
	}
}
