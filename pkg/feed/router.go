package feed

import "github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"

// Router initiate the metrics for the router feeds
type Router struct {
	HealthJVMExporter *export.HealthJVMExporter
}

// Evaluate decides which prometheus export to call
func (rf *Router) Evaluate(f Feed) {
	switch f.Metric {
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
		rf.HealthJVMExporter.SetGCCPU(float64(val))
		break
	default:
		break
	}
}
