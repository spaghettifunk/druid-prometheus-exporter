package export

import "github.com/prometheus/client_golang/prometheus"

// IngestionRealtimeIndexingExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type IngestionRealtimeIndexingExporter struct {
	TaskRunTime       *prometheus.HistogramVec `description:"milliseconds taken to run a task"`
	TaskActionLogTime *prometheus.HistogramVec `description:"milliseconds taken to log a task action to the audit log"`
	TaskActionRunTime *prometheus.HistogramVec `description:"milliseconds taken to execute a task action"`
	SegmentAddedBytes *prometheus.SummaryVec   `description:"size in bytes of new segments created"`
	SegmentMovedBytes *prometheus.SummaryVec   `description:"size in bytes of segments moved/archived via the Move Task"`
	SegmentNukedBytes *prometheus.SummaryVec   `description:"size in bytes of segments deleted via the Kill Task"`
	TaskSuccessCount  *prometheus.GaugeVec     `description:"number of successful tasks per emission period. This metric is only available if the TaskCountStatsMonitor module is included"`
	TaskFailedCount   *prometheus.GaugeVec     `description:"number of failed tasks per emission period. This metric is only available if the TaskCountStatsMonitor module is included"`
	TaskRunningCount  *prometheus.GaugeVec     `description:"number of current running tasks. This metric is only available if the TaskCountStatsMonitor module is included"`
	TaskPendingCount  *prometheus.GaugeVec     `description:"number of current pending tasks. This metric is only available if the TaskCountStatsMonitor module is included"`
	TaskWaitingCount  *prometheus.GaugeVec     `description:"number of current waiting tasks. This metric is only available if the TaskCountStatsMonitor module is included"`
}

// NewIngestionRealtimeIndexingExporter returns a new Jetty exporter object
func NewIngestionRealtimeIndexingExporter() *IngestionRealtimeIndexingExporter {
	re := &IngestionRealtimeIndexingExporter{
		TaskRunTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "realtime_indexing",
			Name:      "task_runtime",
			Help:      "milliseconds taken to run a task",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"dataSource", "taskId", "taskType", "taskStatus"}),
		TaskActionLogTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "realtime_indexing",
			Name:      "task_action_logtime",
			Help:      "milliseconds taken to log a task action to the audit log",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"dataSource", "taskId", "taskType"}),
		TaskActionRunTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "realtime_indexing",
			Name:      "task_action_runtime",
			Help:      "milliseconds taken to execute a task action",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"dataSource", "taskId", "taskType"}),
		SegmentAddedBytes: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Namespace:  "druid",
			Subsystem:  "realtime_indexing",
			Name:       "segment_added_bytes",
			Help:       "size in bytes of new segments created",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, []string{"dataSource", "taskId", "taskType", "interval"}),
		SegmentMovedBytes: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Namespace:  "druid",
			Subsystem:  "realtime_indexing",
			Name:       "segment_moved_bytes",
			Help:       "size in bytes of segments moved/archived via the Move Task",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, []string{"dataSource", "taskId", "taskType", "interval"}),
		SegmentNukedBytes: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Namespace:  "druid",
			Subsystem:  "realtime_indexing",
			Name:       "segment_nuked_bytes",
			Help:       "size in bytes of segments deleted via the Kill Task",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, []string{"dataSource", "taskId", "taskType", "interval"}),
		TaskSuccessCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime_indexing",
			Name:      "task_success_count",
			Help:      "number of successful tasks per emission period. This metric is only available if the TaskCountStatsMonitor module is included",
		}, []string{"dataSource"}),
		TaskFailedCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime_indexing",
			Name:      "task_failed_count",
			Help:      "number of failed tasks per emission period. This metric is only available if the TaskCountStatsMonitor module is included",
		}, []string{"dataSource"}),
		TaskRunningCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime_indexing",
			Name:      "task_running_count",
			Help:      "number of current running tasks. This metric is only available if the TaskCountStatsMonitor module is included",
		}, []string{"dataSource"}),
		TaskPendingCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime_indexing",
			Name:      "task_pending_count",
			Help:      "number of current pending tasks. This metric is only available if the TaskCountStatsMonitor module is included",
		}, []string{"dataSource"}),
		TaskWaitingCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime_indexing",
			Name:      "task_waiting_count",
			Help:      "number of current waiting tasks. This metric is only available if the TaskCountStatsMonitor module is included",
		}, []string{"dataSource"}),
	}

	// register all the prometheus metrics
	prometheus.MustRegister(re.SegmentAddedBytes)
	prometheus.MustRegister(re.SegmentMovedBytes)
	prometheus.MustRegister(re.SegmentNukedBytes)
	prometheus.MustRegister(re.TaskActionLogTime)
	prometheus.MustRegister(re.TaskActionRunTime)
	prometheus.MustRegister(re.TaskFailedCount)
	prometheus.MustRegister(re.TaskPendingCount)
	prometheus.MustRegister(re.TaskWaitingCount)
	prometheus.MustRegister(re.TaskSuccessCount)
	prometheus.MustRegister(re.TaskRunningCount)
	prometheus.MustRegister(re.TaskRunTime)

	return re
}

// SetTaskRunTime .
func (re *IngestionRealtimeIndexingExporter) SetTaskRunTime(labels map[string]string, val float64) {
	re.TaskRunTime.With(labels).Observe(val)
}

// SetTaskActionLogTime .
func (re *IngestionRealtimeIndexingExporter) SetTaskActionLogTime(labels map[string]string, val float64) {
	re.TaskActionLogTime.With(labels).Observe(val)
}

// SetTaskActionRunTime .
func (re *IngestionRealtimeIndexingExporter) SetTaskActionRunTime(labels map[string]string, val float64) {
	re.TaskActionRunTime.With(labels).Observe(val)
}

// SetSegmentAddedBytes .
func (re *IngestionRealtimeIndexingExporter) SetSegmentAddedBytes(labels map[string]string, val float64) {
	re.SegmentAddedBytes.With(labels).Observe(val)
}

// SetSegmentMovedBytes .
func (re *IngestionRealtimeIndexingExporter) SetSegmentMovedBytes(labels map[string]string, val float64) {
	re.SegmentMovedBytes.With(labels).Observe(val)
}

// SetSegmentNukedBytes .
func (re *IngestionRealtimeIndexingExporter) SetSegmentNukedBytes(labels map[string]string, val float64) {
	re.SegmentNukedBytes.With(labels).Observe(val)
}

// SetTaskSuccessCount .
func (re *IngestionRealtimeIndexingExporter) SetTaskSuccessCount(source string, val float64) {
	re.TaskSuccessCount.With(prometheus.Labels{"dataSource": source}).Add(val)
}

// SetTaskFailedCount .
func (re *IngestionRealtimeIndexingExporter) SetTaskFailedCount(source string, val float64) {
	re.TaskFailedCount.With(prometheus.Labels{"dataSource": source}).Add(val)
}

// SetTaskRunningCount .
func (re *IngestionRealtimeIndexingExporter) SetTaskRunningCount(source string, val float64) {
	re.TaskRunningCount.With(prometheus.Labels{"dataSource": source}).Add(val)
}

// SetTaskPendingCount .
func (re *IngestionRealtimeIndexingExporter) SetTaskPendingCount(source string, val float64) {
	re.TaskPendingCount.With(prometheus.Labels{"dataSource": source}).Add(val)
}

// SetTaskWaitingCount .
func (re *IngestionRealtimeIndexingExporter) SetTaskWaitingCount(source string, val float64) {
	re.TaskWaitingCount.With(prometheus.Labels{"dataSource": source}).Add(val)
}
