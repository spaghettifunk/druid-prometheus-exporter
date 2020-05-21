package feed

import (
	"strings"

	"github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"
)

// Coordinator initiate the metrics for the coordinator feeds
type Coordinator struct {
	CoordinationExporter *export.CoordinationExporter
	QueryJettyExporter   *export.QueryJettyExporter
	HealthJVMExporter    *export.HealthJVMExporter
}

// Evaluate decides which prometheus export to call
func (cf *Coordinator) Evaluate(f Feed) {
	switch f.Metric {
	case "tier/total/capacity":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetTierTotalCapacity(f.Tier, float64(val))
		break
	case "tier/replication/factor":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetTierReplicationFactor(f.Tier, float64(val))
		break
	case "tier/historical/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetTierHistoricalCount(f.Tier, float64(val))
		break
	case "segment/underReplicated/count":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"tier":       f.Tier,
			"dataSource": f.DataSource,
		}
		cf.CoordinationExporter.SetSegmentUnderReplicatedCount(labels, float64(val))
		break
	case "segment/unavailable/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentUnavailableCount(f.DataSource, float64(val))
		break
	case "segment/overShadowed/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentOvershadowedCount(float64(val))
		break
	case "segment/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentCount(f.DataSource, float64(val))
		break
	case "segment/size":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentSize(f.DataSource, float64(val))
		break
	case "segment/dropQueue/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentDropQueueCount(f.Server, float64(val))
		break
	case "segment/loadQueue/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentLoadQueueCount(f.Server, float64(val))
		break
	case "segment/loadQueue/failed":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentLoadQueueFailed(f.Server, float64(val))
		break
	case "segment/loadQueue/size":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentLoadQueueSize(f.Server, float64(val))
		break
	case "segment/cost/normalized":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentCostNormalized(f.Tier, float64(val))
		break
	case "segment/cost/normalization":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentCostNormalization(f.Tier, float64(val))
		break
	case "segment/cost/raw":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentCostRaw(f.Tier, float64(val))
		break
	case "segment/unneeded/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentUnneededCount(f.Tier, float64(val))
		break
	case "segment/deleted/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentDeletedCount(f.Tier, float64(val))
		break
	case "segment/dropped/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentDroppedCount(f.Tier, float64(val))
		break
	case "segment/moved/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentMovedCount(f.Tier, float64(val))
		break
	case "segment/assigned/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentAssignedCount(f.Tier, float64(val))
		break
	case "jetty/numOpenConnections":
		val, _ := f.Value.(float64)
		cf.QueryJettyExporter.SetNumOpenConnections(float64(val))
		break
	case "jvm/pool/committed":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"poolKind": f.PoolKind,
			"poolName": f.PoolName,
		}
		cf.HealthJVMExporter.SetPoolCommited(labels, float64(val))
		break
	case "jvm/pool/init":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"poolKind": f.PoolKind,
			"poolName": f.PoolName,
		}
		cf.HealthJVMExporter.SetPoolInit(labels, float64(val))
		break
	case "jvm/pool/max":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"poolKind": f.PoolKind,
			"poolName": f.PoolName,
		}
		cf.HealthJVMExporter.SetPoolMax(labels, float64(val))
		break
	case "jvm/pool/used":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"poolKind": f.PoolKind,
			"poolName": f.PoolName,
		}
		cf.HealthJVMExporter.SetPoolUsed(labels, float64(val))
		break
	case "jvm/bufferpool/count":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"bufferPoolName": f.BufferPoolName,
		}
		cf.HealthJVMExporter.SetBufferpoolCount(labels, float64(val))
		break
	case "jvm/bufferpool/used":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"bufferPoolName": f.BufferPoolName,
		}
		cf.HealthJVMExporter.SetBufferpoolUsed(labels, float64(val))
		break
	case "jvm/bufferpool/capacity":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"bufferPoolName": f.BufferPoolName,
		}
		cf.HealthJVMExporter.SetBufferpoolCapacity(labels, float64(val))
		break
	case "jvm/mem/init":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"memKind": f.MemKind,
		}
		cf.HealthJVMExporter.SetMemInit(labels, float64(val))
		break
	case "jvm/mem/max":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"memKind": f.MemKind,
		}
		cf.HealthJVMExporter.SetMemMax(labels, float64(val))
		break
	case "jvm/mem/used":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"memKind": f.MemKind,
		}
		cf.HealthJVMExporter.SetMemUsed(labels, float64(val))
		break
	case "jvm/mem/committed":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"memKind": f.MemKind,
		}
		cf.HealthJVMExporter.SetMemCommitted(labels, float64(val))
		break
	case "jvm/gc/count":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"gcName": strings.Join(f.GCName, ","),
			"gcGen":  strings.Join(f.GCGen, ","),
		}
		cf.HealthJVMExporter.SetGCCount(labels, float64(val))
		break
	case "jvm/gc/cpu":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"gcName": strings.Join(f.GCName, ","),
			"gcGen":  strings.Join(f.GCGen, ","),
		}
		cf.HealthJVMExporter.SetGCCPU(labels, float64(val))
		break
	default:
		break
	}
}
