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
		break
	case "query/bytes":
		break
	case "query/node/time":
		break
	case "query/node/bytes":
		break
	case
		"query/node/ttfb":
		break
	case "query/node/backpressure":
		break
	case "query/count":
		break
	case
		"query/success/count":
		break
	case "query/failed/count":
		break
	case "query/interrupted/count":
		break
	case
		"segment/scan/pending":
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
	case
		"query/cache/delta/misses":
		break
	case "query/cache/delta/evictions":
		break
	case "query/cache/delta/hitRate":
		break
	case
		"query/cache/delta/averageBytes":
		break
	case "query/cache/delta/timeouts":
		break
	case "query/cache/delta/errors":
		break
	case
		"query/cache/delta/put/ok":
		break
	case "query/cache/delta/put/error":
		break
	case "query/cache/delta/put/oversized":
		break
	case
		"query/cache/total/numEntries":
		break
	case "query/cache/total/sizeBytes":
		break
	case "query/cache/total/hits":
		break
	case
		"query/cache/total/misses":
		break
	case "query/cache/total/evictions":
		break
	case "query/cache/total/hitRate":
		break
	case
		"query/cache/total/averageBytes":
		break
	case "query/cache/total/timeouts":
		break
	case "query/cache/total/errors":
		break
	case
		"query/cache/total/put/ok":
		break
	case "query/cache/total/put/error":
		break
	case "query/cache/total/put/oversized":
		val, _ := f.Value.(float64)
		hf.QueryCacheExporter.SetCacheTotalPutOversized(float64(val))
		break
	case "jvm/pool/committed":
		break
	case "jvm/pool/init":
		break
	case "jvm/pool/max":
		break
	case "jvm/pool/used":
		break
	case "jvm/bufferpool/count":
		break
	case "jvm/bufferpool/used":
		break
	case "jvm/bufferpool/capacity":
		break
	case "jvm/mem/init":
		break
	case "jvm/mem/max":
		break
	case "jvm/mem/used":
		break
	case "jvm/mem/committed":
		break
	case "jvm/gc/count":
		break
	case "jvm/gc/cpu":
		break
		val, _ := f.Value.(float64)
		hf.HealthJVMExporter.SetGCCPU(float64(val))
		break
	case "jetty/numOpenConnections":
		val, _ := f.Value.(float64)
		hf.QueryJettyExporter.SetNumOpenConnections(float64(val))
		break
	case "sys/swap/free":
	case "sys/swap/max":
		break
	case "sys/swap/pageIn":
		break
	case "sys/swap/pageOut":
		break
	case
		"sys/disk/write/count":
		break
	case "sys/disk/read/count":
		break
	case "sys/disk/write/size":
		break
	case
		"sys/disk/read/size":
		break
	case "sys/net/read/size	":
		break
	case "sys/mem/max":
		break
	case "sys/mem/used":
		break
	case
		"sys/fs/max":
		break
	case "sys/fs/used":
		break
	case "sys/net/write/size":
		break
	case "sys/net/read/size":
		break
	case
		"sys/storage/used":
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
