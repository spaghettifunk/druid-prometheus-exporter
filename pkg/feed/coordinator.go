package feed

import "github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"

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
		cf.CoordinationExporter.SetTierTotalCapacity(float64(val))
		break
	case "tier/replication/factor":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetTierReplicationFactor(float64(val))
		break
	case "tier/historical/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetTierHistoricalCount(float64(val))
		break
	case "segment/underReplicated/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentUnderReplicatedCount(float64(val))
		break
	case "segment/unavailable/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentUnavailableCount(float64(val))
		break
	case "segment/overShadowed/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentOvershadowedCount(float64(val))
		break
	case "segment/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentCount(float64(val))
		break
	case "segment/size":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentSize(float64(val))
		break
	case "segment/dropQueue/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentDropQueueCount(float64(val))
		break
	case "segment/loadQueue/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentLoadQueueCount(float64(val))
		break
	case "segment/loadQueue/failed":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentLoadQueueFailed(float64(val))
		break
	case "segment/loadQueue/size":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentLoadQueueSize(float64(val))
		break
	case "segment/cost/normalized":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentCostNormalized(float64(val))
		break
	case "segment/cost/normalization":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentCostNormalization(float64(val))
		break
	case "segment/cost/raw":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentCostRaw(float64(val))
		break
	case "segment/unneeded/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentUnneededCount(float64(val))
		break
	case "segment/deleted/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentDeletedCount(float64(val))
		break
	case "segment/dropped/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentDroppedCount(float64(val))
		break
	case "segment/moved/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentMovedCount(float64(val))
		break
	case "segment/assigned/count":
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentAssignedCount(float64(val))
		break
	case "jetty/numOpenConnections":
		val, _ := f.Value.(float64)
		cf.QueryJettyExporter.SetNumOpenConnections(float64(val))
		break
	case "jvm/pool/committed":
		val, _ := f.Value.(float64)
		cf.HealthJVMExporter.SetPoolCommited(float64(val))
		break
	case "jvm/pool/init":
		val, _ := f.Value.(float64)
		cf.HealthJVMExporter.SetPoolInit(float64(val))
		break
	case "jvm/pool/max":
		val, _ := f.Value.(float64)
		cf.HealthJVMExporter.SetPoolMax(float64(val))
		break
	case "jvm/pool/used":
		val, _ := f.Value.(float64)
		cf.HealthJVMExporter.SetPoolUsed(float64(val))
		break
	case "jvm/bufferpool/count":
		val, _ := f.Value.(float64)
		cf.HealthJVMExporter.SetBufferpoolCount(float64(val))
		break
	case "jvm/bufferpool/used":
		val, _ := f.Value.(float64)
		cf.HealthJVMExporter.SetBufferpoolUsed(float64(val))
		break
	case "jvm/bufferpool/capacity":
		val, _ := f.Value.(float64)
		cf.HealthJVMExporter.SetBufferpoolCapacity(float64(val))
		break
	case "jvm/mem/init":
		val, _ := f.Value.(float64)
		cf.HealthJVMExporter.SetMemInit(float64(val))
		break
	case "jvm/mem/max":
		val, _ := f.Value.(float64)
		cf.HealthJVMExporter.SetMemMax(float64(val))
		break
	case "jvm/mem/used":
		val, _ := f.Value.(float64)
		cf.HealthJVMExporter.SetMemUsed(float64(val))
		break
	case "jvm/mem/committed":
		val, _ := f.Value.(float64)
		cf.HealthJVMExporter.SetMemCommitted(float64(val))
		break
	case "jvm/gc/count":
		val, _ := f.Value.(float64)
		cf.HealthJVMExporter.SetGCCount(float64(val))
		break
	case "jvm/gc/cpu":
		val, _ := f.Value.(float64)
		cf.HealthJVMExporter.SetGCCPU(float64(val))
		break
	default:
		break
	}
}
