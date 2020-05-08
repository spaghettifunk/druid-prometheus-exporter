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
		val, _ := f.Value.(float64)
		rf.HealthJVMExporter.SetPoolCommited(float64(val))
		break
	case "jvm/pool/init":
		val, _ := f.Value.(float64)
		rf.HealthJVMExporter.SetPoolInit(float64(val))
		break
	case "jvm/pool/max":
		val, _ := f.Value.(float64)
		rf.HealthJVMExporter.SetMemMax(float64(val))
		break
	case "jvm/pool/used":
		val, _ := f.Value.(float64)
		rf.HealthJVMExporter.SetPoolUsed(float64(val))
		break
	case "jvm/bufferpool/count":
		val, _ := f.Value.(float64)
		rf.HealthJVMExporter.SetBufferpoolCount(float64(val))
		break
	case "jvm/bufferpool/used":
		val, _ := f.Value.(float64)
		rf.HealthJVMExporter.SetBufferpoolUsed(float64(val))
		break
	case "jvm/bufferpool/capacity":
		val, _ := f.Value.(float64)
		rf.HealthJVMExporter.SetBufferpoolCapacity(float64(val))
		break
	case "jvm/mem/init":
		val, _ := f.Value.(float64)
		rf.HealthJVMExporter.SetMemInit(float64(val))
		break
	case "jvm/mem/max":
		val, _ := f.Value.(float64)
		rf.HealthJVMExporter.SetMemMax(float64(val))
		break
	case "jvm/mem/used":
		val, _ := f.Value.(float64)
		rf.HealthJVMExporter.SetMemUsed(float64(val))
		break
	case "jvm/mem/committed":
		val, _ := f.Value.(float64)
		rf.HealthJVMExporter.SetMemCommitted(float64(val))
		break
	case "jvm/gc/count":
		val, _ := f.Value.(float64)
		rf.HealthJVMExporter.SetGCCount(float64(val))
		break
	case "jvm/gc/cpu":
		val, _ := f.Value.(float64)
		rf.HealthJVMExporter.SetGCCPU(float64(val))
		break
	default:
		break
	}
}
