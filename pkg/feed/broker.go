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
		bf.QueryBrokerExporter.SetQueryTime(float64(val))
		break
	case "query/bytes":
		val, _ := f.Value.(float64)
		bf.QueryBrokerExporter.SetQueryBytes(float64(val))
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
		break
	case "sys/swap/max":
		break
	case "sys/swap/pageIn":
		break
	case "sys/swap/pageOut":
		break
	case "sys/disk/write/count":
		break
	case "sys/disk/read/count":
		break
	case "sys/disk/write/size":
		break
	case "sys/disk/read/size":
		break
	case "sys/net/read/size":
		break
	case "sys/mem/max":
		break
	case "sys/mem/used":
		break
	case "sys/fs/max":
		break
	case "sys/fs/used":
		break
	case "sys/net/write/size":
		break
	case "sys/storage/used":
		break
	case "sys/cpu":
		val, _ := f.Value.(float64)
		bf.SysExporter.SetCPU(float64(val))
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
		val, _ := f.Value.(float64)
		bf.HealthJVMExporter.SetGCCPU(float64(val))
		break
	case "jetty/numOpenConnections":
		bf.QueryJettyExporter.FormatMetrics()
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
		bf.QueryCacheExporter.FormatMetrics()
		break
	default:
		break
	}
}
