package export

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	healthURL      = "/status/health"
	segmentDataURL = "/druid/coordinator/v1/datasources?simple"
	tasksURL       = "/druid/indexer/v1/tasks"
	supervisorURL  = "/druid/indexer/v1/supervisor?full"
)

// CommmonExporter holds all the Prometheus metrics that maps the metrics
type CommmonExporter struct {
	// HealthStatus returns the health checks status
	HealthStatus *prometheus.Desc
}

// NewCommonExporter returns a new collector object
func NewCommonExporter() *CommmonExporter {
	return &CommmonExporter{
		HealthStatus: prometheus.NewDesc("druid_health_status",
			"Health of Druid, 1 is healthy 0 is not",
			nil, prometheus.Labels{
				"druid": "health",
			},
		),
	}
}

// Describe will associate the value for druid exporter
func (dc *CommmonExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- dc.HealthStatus
}

// Collect returns the prometheus metrics from Druid's generic
func (dc *CommmonExporter) Collect(ch chan<- prometheus.Metric) {
	// ch <- prometheus.MustNewConstMetric(dc.HealthStatus, prometheus.CounterValue, dc.GetHealthMetrics())

	// for _, data := range dc.GetData(tasksURL) {
	// 	ch <- prometheus.MustNewConstMetric(dc.Tasks, prometheus.GaugeValue, float64(1), fmt.Sprintf("%v", data["dataSource"]), fmt.Sprintf("%v", data["groupId"]), fmt.Sprintf("%v", data["status"]), fmt.Sprintf("%v", data["createdTime"]))
	// }
}
