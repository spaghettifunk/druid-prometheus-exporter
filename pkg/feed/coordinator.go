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
		break
	case "tier/replication/factor":
		break
	case "tier/historical/count":
		break
	case "segment/underReplicated/count":
		break
	case "segment/unavailable/count":
		break
	case "segment/overShadowed/count":
		break
	case "segment/count":
		break
	case "segment/size":
		break
	case "segment/dropQueue/count":
		break
	case "segment/loadQueue/count":
		break
	case "segment/loadQueue/failed":
		break
	case "segment/loadQueue/size":
		break
	case "segment/cost/normalized":
		break
	case "segment/cost/normalization":
		break
	case "segment/cost/raw":
		break
	case "segment/unneeded/count":
		break
	case "segment/deleted/count":
		break
	case "segment/dropped/count":
		break
	case "segment/moved/count":
		break
	case "segment/assigned/count":
		break
		val, _ := f.Value.(float64)
		cf.CoordinationExporter.SetSegmentAssignedCount(float64(val))
		break
	case "jetty/numOpenConnections":
		val, _ := f.Value.(float64)
		cf.QueryJettyExporter.SetNumOpenConnections(float64(val))
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
		cf.HealthJVMExporter.SetGCCPU(float64(val))
		break
	default:
		break
	}
}
