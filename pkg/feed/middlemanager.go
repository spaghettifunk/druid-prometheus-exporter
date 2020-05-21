package feed

import (
	"strings"

	"github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"
)

// MiddleManager initiate the metrics for the middle manager feeds
type MiddleManager struct {
	HealthJVMExporter *export.HealthJVMExporter
}

// Evaluate decides which prometheus export to call
func (mmf *MiddleManager) Evaluate(f Feed) {
	switch f.Metric {
	case "jvm/pool/committed":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"poolKind": f.PoolKind,
			"poolName": f.PoolName,
		}
		mmf.HealthJVMExporter.SetPoolCommited(labels, float64(val))
		break
	case "jvm/pool/init":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"poolKind": f.PoolKind,
			"poolName": f.PoolName,
		}
		mmf.HealthJVMExporter.SetPoolInit(labels, float64(val))
		break
	case "jvm/pool/max":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"poolKind": f.PoolKind,
			"poolName": f.PoolName,
		}
		mmf.HealthJVMExporter.SetPoolMax(labels, float64(val))
		break
	case "jvm/pool/used":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"poolKind": f.PoolKind,
			"poolName": f.PoolName,
		}
		mmf.HealthJVMExporter.SetPoolUsed(labels, float64(val))
		break
	case "jvm/bufferpool/count":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"bufferPoolName": f.BufferPoolName,
		}
		mmf.HealthJVMExporter.SetBufferpoolCount(labels, float64(val))
		break
	case "jvm/bufferpool/used":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"bufferPoolName": f.BufferPoolName,
		}
		mmf.HealthJVMExporter.SetBufferpoolUsed(labels, float64(val))
		break
	case "jvm/bufferpool/capacity":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"bufferPoolName": f.BufferPoolName,
		}
		mmf.HealthJVMExporter.SetBufferpoolCapacity(labels, float64(val))
		break
	case "jvm/mem/init":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"memKind": f.MemKind,
		}
		mmf.HealthJVMExporter.SetMemInit(labels, float64(val))
		break
	case "jvm/mem/max":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"memKind": f.MemKind,
		}
		mmf.HealthJVMExporter.SetMemMax(labels, float64(val))
		break
	case "jvm/mem/used":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"memKind": f.MemKind,
		}
		mmf.HealthJVMExporter.SetMemUsed(labels, float64(val))
		break
	case "jvm/mem/committed":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"memKind": f.MemKind,
		}
		mmf.HealthJVMExporter.SetMemCommitted(labels, float64(val))
		break
	case "jvm/gc/count":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"gcName": strings.Join(f.GCName, ","),
			"gcGen":  strings.Join(f.GCGen, ","),
		}
		mmf.HealthJVMExporter.SetGCCount(labels, float64(val))
		break
	case "jvm/gc/cpu":
		val, _ := f.Value.(float64)
		labels := map[string]string{
			"gcName": strings.Join(f.GCName, ","),
			"gcGen":  strings.Join(f.GCGen, ","),
		}
		mmf.HealthJVMExporter.SetGCCPU(labels, float64(val))
		break
	default:
		break
	}
}
