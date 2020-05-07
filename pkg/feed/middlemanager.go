package feed

import "github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"

// MiddleManager initiate the metrics for the middle manager feeds
type MiddleManager struct {
	HealthJVMExporter *export.HealthJVMExporter
}

// Evaluate decides which prometheus export to call
func (mmf *MiddleManager) Evaluate(f Feed) {
	switch f.Metric {
	case "jvm/pool/committed":
		val, _ := f.Value.(float64)
		mmf.HealthJVMExporter.SetPoolCommited(float64(val))
		break
	case "jvm/pool/init":
		val, _ := f.Value.(float64)
		mmf.HealthJVMExporter.SetPoolInit(float64(val))
		break
	case "jvm/pool/max":
		val, _ := f.Value.(float64)
		mmf.HealthJVMExporter.SetMemMax(float64(val))
		break
	case "jvm/pool/used":
		val, _ := f.Value.(float64)
		mmf.HealthJVMExporter.SetPoolUsed(float64(val))
		break
	case "jvm/bufferpool/count":
		val, _ := f.Value.(float64)
		mmf.HealthJVMExporter.SetBufferpoolCount(float64(val))
		break
	case "jvm/bufferpool/used":
		val, _ := f.Value.(float64)
		mmf.HealthJVMExporter.SetBufferpoolUsed(float64(val))
		break
	case "jvm/bufferpool/capacity":
		val, _ := f.Value.(float64)
		mmf.HealthJVMExporter.SetBufferpoolCapacity(float64(val))
		break
	case "jvm/mem/init":
		val, _ := f.Value.(float64)
		mmf.HealthJVMExporter.SetMemInit(float64(val))
		break
	case "jvm/mem/max":
		val, _ := f.Value.(float64)
		mmf.HealthJVMExporter.SetMemMax(float64(val))
		break
	case "jvm/mem/used":
		val, _ := f.Value.(float64)
		mmf.HealthJVMExporter.SetMemUsed(float64(val))
		break
	case "jvm/mem/committed":
		val, _ := f.Value.(float64)
		mmf.HealthJVMExporter.SetMemCommitted(float64(val))
		break
	case "jvm/gc/count":
		val, _ := f.Value.(float64)
		mmf.HealthJVMExporter.SetGCCount(float64(val))
		break
	case "jvm/gc/cpu":
		val, _ := f.Value.(float64)
		mmf.HealthJVMExporter.SetGCCPU(float64(val))
		break
	default:
		break
	}
}
