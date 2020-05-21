package feed

import (
	"github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"
)

// Broker initiate the metrics for the broker feeds
type Broker struct {
	QueryBrokerExporter *export.QueryBrokerExporter
	QueryJettyExporter  *export.QueryJettyExporter
	QueryCacheExporter  *export.QueryCacheExporter
	SQLExporter         *export.SQLExporter
	SysExporter         *export.SysExporter
	HealthJVMExporter   *export.HealthJVMExporter
}

// Evaluate decides which prometheus export to call
func (bf *Broker) Evaluate(f Feed) {
	switch f.Metric {
	case "query/time":
		val, _ := f.Value.(float64)
		bf.QueryBrokerExporter.SetQueryTime(f.DataSource, float64(val))
		break
	case "query/bytes":
		val, _ := f.Value.(float64)
		bf.QueryBrokerExporter.SetQueryBytes(f.DataSource, float64(val))
		break
	case "query/node/time":
		val, _ := f.Value.(float64)
		bf.QueryBrokerExporter.SetQueryNodeTime(float64(val))
		break
	case "query/node/bytes":
		val, _ := f.Value.(float64)
		bf.QueryBrokerExporter.SetQueryNodeBytes(float64(val))
		break
	case "query/node/ttfb":
		val, _ := f.Value.(float64)
		bf.QueryBrokerExporter.SetQueryNodetTtfb(float64(val))
		break
	case "query/node/backpressure":
		val, _ := f.Value.(float64)
		bf.QueryBrokerExporter.SetQueryNodeBackpressure(float64(val))
		break
	case "query/count":
		val, _ := f.Value.(float64)
		bf.QueryBrokerExporter.SetQueryCount(float64(val))
		break
	case "query/success/count":
		val, _ := f.Value.(float64)
		bf.QueryBrokerExporter.SetQuerySuccessCount(float64(val))
		break
	case "query/failed/count":
		val, _ := f.Value.(float64)
		bf.QueryBrokerExporter.SetQueryFailedCount(float64(val))
		break
	case "query/interrupted/count":
		val, _ := f.Value.(float64)
		bf.QueryBrokerExporter.SetQueryInterruptedCount(float64(val))
		break
	case "sqlQuery/time":
		val, _ := f.Value.(float64)
		bf.SQLExporter.SetSQLQueryTime(float64(val))
		break
	case "sqlQuery/bytes":
		val, _ := f.Value.(float64)
		bf.SQLExporter.SetSQLQueryBytes(float64(val))
		break
	case "sys/swap/free":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetSwapFree(float64(val))
		break
	case "sys/swap/max":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetSwapMax(float64(val))
		break
	case "sys/swap/pageIn":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetSwapPageIn(float64(val))
		break
	case "sys/swap/pageOut":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetSwapPageOut(float64(val))
		break
	case "sys/disk/write/count":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetDiskWriteCount(float64(val))
		break
	case "sys/disk/read/count":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetDiskReadCount(float64(val))
		break
	case "sys/disk/write/size":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetDiskWriteSize(float64(val))
		break
	case "sys/disk/read/size":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetDiskReadSize(float64(val))
		break
	case "sys/net/read/size":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetNetReadSize(float64(val))
		break
	case "sys/mem/max":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetMemMax(float64(val))
		break
	case "sys/mem/used":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetMemUsed(float64(val))
		break
	case "sys/fs/max":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetFSMax(float64(val))
		break
	case "sys/fs/used":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetFSUsed(float64(val))
		break
	case "sys/net/write/size":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetNetWriteSize(float64(val))
		break
	case "sys/storage/used":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetStorageUsed(float64(val))
		break
	case "sys/cpu":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetCPU(float64(val))
		break
	case "jvm/pool/committed":
		val, _ := f.Value.(float64)
		bf.HealthJVMExporter.SetPoolCommited(float64(val))
		break
	case "jvm/pool/init":
		val, _ := f.Value.(float64)
		bf.HealthJVMExporter.SetPoolInit(float64(val))
		break
	case "jvm/pool/max":
		val, _ := f.Value.(float64)
		bf.HealthJVMExporter.SetPoolMax(float64(val))
		break
	case "jvm/pool/used":
		val, _ := f.Value.(float64)
		bf.HealthJVMExporter.SetPoolUsed(float64(val))
		break
	case "jvm/bufferpool/count":
		val, _ := f.Value.(float64)
		bf.HealthJVMExporter.SetBufferpoolCount(float64(val))
		break
	case "jvm/bufferpool/used":
		val, _ := f.Value.(float64)
		bf.HealthJVMExporter.SetBufferpoolUsed(float64(val))
		break
	case "jvm/bufferpool/capacity":
		val, _ := f.Value.(float64)
		bf.HealthJVMExporter.SetBufferpoolCapacity(float64(val))
		break
	case "jvm/mem/init":
		val, _ := f.Value.(float64)
		bf.HealthJVMExporter.SetMemInit(float64(val))
		break
	case "jvm/mem/max":
		val, _ := f.Value.(float64)
		bf.HealthJVMExporter.SetMemMax(float64(val))
		break
	case "jvm/mem/used":
		val, _ := f.Value.(float64)
		bf.HealthJVMExporter.SetMemUsed(float64(val))
		break
	case "jvm/mem/committed":
		val, _ := f.Value.(float64)
		bf.HealthJVMExporter.SetMemCommitted(float64(val))
		break
	case "jvm/gc/count":
		val, _ := f.Value.(float64)
		bf.HealthJVMExporter.SetGCCount(float64(val))
		break
	case "jvm/gc/cpu":
		val, _ := f.Value.(float64)
		bf.HealthJVMExporter.SetGCCPU(float64(val))
		break
	case "jetty/numOpenConnections":
		val, _ := f.Value.(float64)
		bf.QueryJettyExporter.SetNumOpenConnections(float64(val))
		break
	case "query/cache/delta/numEntries":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetDeltaNumEntries(float64(val))
		break
	case "query/cache/delta/sizeBytes":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetDeltaSizeBytes(float64(val))
		break
	case "query/cache/delta/hits":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetDeltaHits(float64(val))
		break
	case "query/cache/delta/misses":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetDeltaMisses(float64(val))
		break
	case "query/cache/delta/evictions":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetDeltaEvictions(float64(val))
		break
	case "query/cache/delta/hitRate":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetDeltaHitRate(float64(val))
		break
	case "query/cache/delta/averageBytes":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetDeltaAverageBytes(float64(val))
		break
	case "query/cache/delta/timeouts":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetDeltaTimeouts(float64(val))
		break
	case "query/cache/delta/errors":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetDeltaErrors(float64(val))
		break
	case "query/cache/delta/put/ok":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetDeltaPutOK(float64(val))
		break
	case "query/cache/delta/put/error":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetDeltaPutError(float64(val))
		break
	case "query/cache/delta/put/oversized":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetDeltaPutOversized(float64(val))
		break
	case "query/cache/total/numEntries":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetTotalNumEntries(float64(val))
		break
	case "query/cache/total/sizeBytes":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetTotalSizeBytes(float64(val))
		break
	case "query/cache/total/hits":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetTotalHits(float64(val))
		break
	case "query/cache/total/misses":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetTotalMisses(float64(val))
		break
	case "query/cache/total/evictions":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetTotalEvictions(float64(val))
		break
	case "query/cache/total/hitRate":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetTotalHitRate(float64(val))
		break
	case "query/cache/total/averageBytes":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetTotalAverageBytes(float64(val))
		break
	case "query/cache/total/timeouts":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetTotalTimeouts(float64(val))
		break
	case "query/cache/total/errors":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetTotalErrors(float64(val))
		break
	case "query/cache/total/put/ok":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetTotalPutOK(float64(val))
		break
	case "query/cache/total/put/error":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetTotalPutError(float64(val))
		break
	case "query/cache/total/put/oversized":
		val, _ := f.Value.(float64)
		bf.QueryCacheExporter.SetTotalPutOversized(float64(val))
		break
	default:
		break
	}
}
