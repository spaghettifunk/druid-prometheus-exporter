package feed

import (
	"strings"

	"github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"
)

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
		labels := map[string]string{
			"dataSource":        f.DataSource,
			"remoteAddress":     f.RemoteAddess,
			"id":                f.ID,
			"numMetrics":        f.Metric,
			"type":              f.Type,
			"hasFilters":        f.HasFilters,
			"duration":          f.Duration,
			"numComplexMetrics": f.NumComplexMetrics,
			"numDimensions":     f.NumDimensions,
			"threshold":         f.Threshold,
			"dimension":         f.Dimension,
		}
		hf.QueryHistoricalExporter.SetQueryTime(labels, float64(val))
		break
	case "query/segment/time":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"id":      f.ID,
			"status":  f.Status,
			"segment": f.Segment,
		}
		hf.QueryHistoricalExporter.SetQuerySegmentTime(labels, float64(val))
		break
	case "query/wait/time":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"id":      f.ID,
			"segment": f.Segment,
		}
		hf.QueryHistoricalExporter.SetQueryWaitTime(labels, float64(val))
		break
	case "query/segmentAndCache/time":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"id":      f.ID,
			"segment": f.Segment,
		}
		hf.QueryHistoricalExporter.SetQuerySegmentAndCacheTime(labels, float64(val))
		break
	case "query/cpu/time":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"dataSource":        f.DataSource,
			"remoteAddress":     f.RemoteAddess,
			"id":                f.ID,
			"numMetrics":        f.Metric,
			"type":              f.Type,
			"hasFilters":        f.HasFilters,
			"duration":          f.Duration,
			"numComplexMetrics": f.NumComplexMetrics,
			"numDimensions":     f.NumDimensions,
			"threshold":         f.Threshold,
			"dimension":         f.Dimension,
		}
		hf.QueryHistoricalExporter.SetQueryCPUTime(labels, float64(val))
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
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetDeltaNumEntries(float64(val))
		break
	case "query/cache/delta/sizeBytes":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetDeltaSizeBytes(float64(val))
		break
	case "query/cache/delta/hits":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetDeltaHits(float64(val))
		break
	case "query/cache/delta/misses":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetDeltaMisses(float64(val))
		break
	case "query/cache/delta/evictions":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetDeltaEvictions(float64(val))
		break
	case "query/cache/delta/hitRate":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetDeltaHitRate(float64(val))
		break
	case "query/cache/delta/averageBytes":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetDeltaAverageBytes(float64(val))
		break
	case "query/cache/delta/timeouts":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetDeltaTimeouts(float64(val))
		break
	case "query/cache/delta/errors":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetDeltaErrors(float64(val))
		break
	case "query/cache/delta/put/ok":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetDeltaPutOK(float64(val))
		break
	case "query/cache/delta/put/error":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetDeltaPutError(float64(val))
		break
	case "query/cache/delta/put/oversized":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetDeltaPutOversized(float64(val))
		break
	case "query/cache/total/numEntries":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetTotalNumEntries(float64(val))
		break
	case "query/cache/total/sizeBytes":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetTotalSizeBytes(float64(val))
		break
	case "query/cache/total/hits":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetTotalHits(float64(val))
		break
	case "query/cache/total/misses":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetTotalMisses(float64(val))
		break
	case "query/cache/total/evictions":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetTotalEvictions(float64(val))
		break
	case "query/cache/total/hitRate":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetTotalHitRate(float64(val))
		break
	case "query/cache/total/averageBytes":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetTotalAverageBytes(float64(val))
		break
	case "query/cache/total/timeouts":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetTotalTimeouts(float64(val))
		break
	case "query/cache/total/errors":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetTotalErrors(float64(val))
		break
	case "query/cache/total/put/ok":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetTotalPutOK(float64(val))
		break
	case "query/cache/total/put/error":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetTotalPutError(float64(val))
		break
	case "query/cache/total/put/oversized":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetTotalPutOversized(float64(val))
		break
	case "jvm/pool/committed":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"poolKind": f.PoolKind,
			"poolName": f.PoolName,
		}
		hf.HealthJVMExporter.SetPoolCommited(labels, float64(val))
		break
	case "jvm/pool/init":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"poolKind": f.PoolKind,
			"poolName": f.PoolName,
		}
		hf.HealthJVMExporter.SetPoolInit(labels, float64(val))
		break
	case "jvm/pool/max":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"poolKind": f.PoolKind,
			"poolName": f.PoolName,
		}
		hf.HealthJVMExporter.SetPoolMax(labels, float64(val))
		break
	case "jvm/pool/used":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"poolKind": f.PoolKind,
			"poolName": f.PoolName,
		}
		hf.HealthJVMExporter.SetPoolUsed(labels, float64(val))
		break
	case "jvm/bufferpool/count":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"bufferPoolName": f.BufferPoolName,
		}
		hf.HealthJVMExporter.SetBufferpoolCount(labels, float64(val))
		break
	case "jvm/bufferpool/used":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"bufferPoolName": f.BufferPoolName,
		}
		hf.HealthJVMExporter.SetBufferpoolUsed(labels, float64(val))
		break
	case "jvm/bufferpool/capacity":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"bufferPoolName": f.BufferPoolName,
		}
		hf.HealthJVMExporter.SetBufferpoolCapacity(labels, float64(val))
		break
	case "jvm/mem/init":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"memKind": f.MemKind,
		}
		hf.HealthJVMExporter.SetMemInit(labels, float64(val))
		break
	case "jvm/mem/max":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"memKind": f.MemKind,
		}
		hf.HealthJVMExporter.SetMemMax(labels, float64(val))
		break
	case "jvm/mem/used":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"memKind": f.MemKind,
		}
		hf.HealthJVMExporter.SetMemUsed(labels, float64(val))
		break
	case "jvm/mem/committed":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"memKind": f.MemKind,
		}
		hf.HealthJVMExporter.SetMemCommitted(labels, float64(val))
		break
	case "jvm/gc/count":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"gcName": strings.Join(f.GCName, ","),
			"gcGen":  strings.Join(f.GCGen, ","),
		}
		hf.HealthJVMExporter.SetGCCount(labels, float64(val))
		break
	case "jvm/gc/cpu":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"gcName": strings.Join(f.GCName, ","),
			"gcGen":  strings.Join(f.GCGen, ","),
		}
		hf.HealthJVMExporter.SetGCCPU(labels, float64(val))
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
		labels := map[string]string{
			"fsDevName":     f.FSDevName,
			"fsDirName":     f.FSDirName,
			"fsTypeName":    f.FSTypeName,
			"fsSysTypeName": f.FSSysTypeName,
			"fsOptions":     strings.Join(f.FSOptions, ","),
		}
		hf.SysExporter.SetDiskWriteCount(labels, float64(val))
		break
	case "sys/disk/read/count":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"fsDevName":     f.FSDevName,
			"fsDirName":     f.FSDirName,
			"fsTypeName":    f.FSTypeName,
			"fsSysTypeName": f.FSSysTypeName,
			"fsOptions":     strings.Join(f.FSOptions, ","),
		}
		hf.SysExporter.SetDiskReadCount(labels, float64(val))
		break
	case "sys/disk/write/size":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"fsDevName":     f.FSDevName,
			"fsDirName":     f.FSDirName,
			"fsTypeName":    f.FSTypeName,
			"fsSysTypeName": f.FSSysTypeName,
			"fsOptions":     strings.Join(f.FSOptions, ","),
		}
		hf.SysExporter.SetDiskWriteSize(labels, float64(val))
		break
	case "sys/disk/read/size":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"fsDevName":     f.FSDevName,
			"fsDirName":     f.FSDirName,
			"fsTypeName":    f.FSTypeName,
			"fsSysTypeName": f.FSSysTypeName,
			"fsOptions":     strings.Join(f.FSOptions, ","),
		}
		hf.SysExporter.SetDiskReadSize(labels, float64(val))
		break
	case "sys/net/read/size":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"netName":    f.NetName,
			"netAddress": f.NetAddress,
			"netHwaddr":  f.NetHwAddr,
		}
		hf.SysExporter.SetNetReadSize(labels, float64(val))
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
		labels := map[string]string{
			"fsDevName":     f.FSDevName,
			"fsDirName":     f.FSDirName,
			"fsTypeName":    f.FSTypeName,
			"fsSysTypeName": f.FSSysTypeName,
			"fsOptions":     strings.Join(f.FSOptions, ","),
		}
		hf.SysExporter.SetFSMax(labels, float64(val))
		break
	case "sys/fs/used":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"fsDevName":     f.FSDevName,
			"fsDirName":     f.FSDirName,
			"fsTypeName":    f.FSTypeName,
			"fsSysTypeName": f.FSSysTypeName,
			"fsOptions":     strings.Join(f.FSOptions, ","),
		}
		hf.SysExporter.SetFSUsed(labels, float64(val))
		break
	case "sys/net/write/size":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"netName":    f.NetName,
			"netAddress": f.NetAddress,
			"netHwaddr":  f.NetHwAddr,
		}
		hf.SysExporter.SetNetWriteSize(labels, float64(val))
		break
	case "sys/storage/used":
		val, _ := f.Value.(float64)
		hf.SysExporter.SetStorageUsed(f.FSDirName, float64(val))
		break
	case "sys/cpu":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"cpuName": f.CPUName,
			"cpuTime": f.CPUTime,
		}
		hf.SysExporter.SetCPU(labels, float64(val))
		break
	case "segment/max":
		val, _ := f.Value.(float64)
		hf.HealthHistoricalExporter.SetSegmentMax(float64(val))
		break
	default:
		break
	}
}
