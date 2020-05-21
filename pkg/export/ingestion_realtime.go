package export

import "github.com/prometheus/client_golang/prometheus"

// IngestionRealtimeExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type IngestionRealtimeExporter struct {
	EventsThrownAway     *prometheus.GaugeVec     `description:"number of events rejected because they are outside the windowPeriod"`
	EventsUnparsable     *prometheus.GaugeVec     `description:"number of events rejected because the events are unparseable"`
	EventsDuplicate      *prometheus.GaugeVec     `description:"number of events rejected because the events are duplicated"`
	EventsProcessed      *prometheus.GaugeVec     `description:"number of events successfully processed per emission period"`
	EventsMessageGap     *prometheus.GaugeVec     `description:"time gap between the data time in event and current system time"`
	RowsOutput           *prometheus.GaugeVec     `description:"number of Druid rows persisted"`
	PersistsCount        *prometheus.GaugeVec     `description:"number of times persist occurred"`
	PersistsTime         *prometheus.HistogramVec `description:"milliseconds spent doing intermediate persist"`
	PersistsCPU          *prometheus.HistogramVec `description:"cpu time in Nanoseconds spent on doing intermediate persist"`
	PersistsBackPressure *prometheus.HistogramVec `description:"milliseconds spent creating persist tasks and blocking waiting for them to finish"`
	PersistsFailed       *prometheus.GaugeVec     `description:"number of persists that failed"`
	HandOffFailed        *prometheus.GaugeVec     `description:"number of handoffs that failed"`
	HandOffCount         *prometheus.GaugeVec     `description:"number of handoffs that happene"`
	MergeTime            *prometheus.HistogramVec `description:"milliseconds spent merging intermediate segments"`
	MergeCPU             *prometheus.HistogramVec `description:"cpu time in Nanoseconds spent on merging intermediate segments"`
	SinkCount            *prometheus.GaugeVec     `description:"number of sinks not handoffed"`
}

// NewIngestionRealtimeExporter returns a new Jetty exporter object
func NewIngestionRealtimeExporter() *IngestionRealtimeExporter {
	re := &IngestionRealtimeExporter{
		EventsThrownAway: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_events_thrown_away_count",
			Help:      "number of events rejected because they are outside the windowPeriod",
		}, []string{"dataSource", "taskId"}),
		EventsUnparsable: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_events_unparsable_count",
			Help:      "number of events rejected because the events are unparseable",
		}, []string{"dataSource", "taskId"}),
		EventsDuplicate: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_events_duplicate_count",
			Help:      "number of events rejected because the events are duplicated",
		}, []string{"dataSource", "taskId"}),
		EventsProcessed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_events_processed_count",
			Help:      "number of events successfully processed per emission period",
		}, []string{"dataSource", "taskId"}),
		EventsMessageGap: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_events_messagegap",
			Help:      "time gap between the data time in event and current system time",
		}, []string{"dataSource", "taskId"}),
		RowsOutput: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_rows_output_count",
			Help:      "number of Druid rows persisted",
		}, []string{"dataSource", "taskId"}),
		PersistsCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_persists_count",
			Help:      "number of times persist occurred",
		}, []string{"dataSource", "taskId"}),
		PersistsTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_persists_time",
			Help:      "milliseconds spent doing intermediate persist",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"dataSource"}),
		PersistsCPU: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_persists_cpu",
			Help:      "cpu time in Nanoseconds spent on doing intermediate persist",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"dataSource"}),
		PersistsBackPressure: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_persists_backpressure",
			Help:      "milliseconds spent creating persist tasks and blocking waiting for them to finish",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"dataSource"}),
		PersistsFailed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_persists_failed_count",
			Help:      "number of times persist failed",
		}, []string{"dataSource", "taskId"}),
		HandOffFailed: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_handoff_failed_count",
			Help:      "Number of times handoff failed",
		}, []string{"dataSource", "taskId"}),
		HandOffCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_handoff_count",
			Help:      "number of times handoff count",
		}, []string{"dataSource", "taskId"}),
		MergeTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_merge_time",
			Help:      "milliseconds spent merging intermediate segments",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"dataSource"}),
		MergeCPU: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_merge_cpu",
			Help:      "cpu time in Nanoseconds spent on merging intermediate segments",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"dataSource"}),
		SinkCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_sink_count",
			Help:      "number of sinks not handoffed",
		}, []string{"dataSource", "taskId"}),
	}

	// register all the prometheus metrics
	prometheus.MustRegister(re.EventsThrownAway)
	prometheus.MustRegister(re.EventsUnparsable)
	prometheus.MustRegister(re.EventsDuplicate)
	prometheus.MustRegister(re.EventsProcessed)
	prometheus.MustRegister(re.EventsMessageGap)
	prometheus.MustRegister(re.RowsOutput)
	prometheus.MustRegister(re.PersistsCount)
	prometheus.MustRegister(re.PersistsCPU)
	prometheus.MustRegister(re.PersistsTime)
	prometheus.MustRegister(re.PersistsBackPressure)
	prometheus.MustRegister(re.PersistsFailed)
	prometheus.MustRegister(re.HandOffFailed)
	prometheus.MustRegister(re.HandOffCount)
	prometheus.MustRegister(re.MergeTime)
	prometheus.MustRegister(re.MergeCPU)
	prometheus.MustRegister(re.SinkCount)

	return re
}

// SetEventsThrownAway .
func (re *IngestionRealtimeExporter) SetEventsThrownAway(labels map[string]string, val float64) {
	re.EventsThrownAway.With(labels).Add(val)
}

// SetEventsUnparsable .
func (re *IngestionRealtimeExporter) SetEventsUnparsable(labels map[string]string, val float64) {
	re.EventsUnparsable.With(labels).Add(val)
}

// SetEventsDuplicate .
func (re *IngestionRealtimeExporter) SetEventsDuplicate(labels map[string]string, val float64) {
	re.EventsDuplicate.With(labels).Add(val)
}

// SetEventsProcessed .
func (re *IngestionRealtimeExporter) SetEventsProcessed(labels map[string]string, val float64) {
	re.EventsProcessed.With(labels).Add(val)
}

// SetEventsMessageGap .
func (re *IngestionRealtimeExporter) SetEventsMessageGap(labels map[string]string, val float64) {
	re.EventsMessageGap.With(labels).Add(val)
}

// SetRowsOutput .
func (re *IngestionRealtimeExporter) SetRowsOutput(labels map[string]string, val float64) {
	re.RowsOutput.With(labels).Add(val)
}

// SetPersistsCount .
func (re *IngestionRealtimeExporter) SetPersistsCount(labels map[string]string, val float64) {
	re.PersistsCount.With(labels).Add(val)
}

// SetPersistsTime .
func (re *IngestionRealtimeExporter) SetPersistsTime(source string, val float64) {
	re.PersistsTime.With(prometheus.Labels{"dataSource": source}).Observe(val)
}

// SetPersistsCPU .
func (re *IngestionRealtimeExporter) SetPersistsCPU(source string, val float64) {
	re.PersistsCPU.With(prometheus.Labels{"dataSource": source}).Observe(val)
}

// SetPersistsBackPressure .
func (re *IngestionRealtimeExporter) SetPersistsBackPressure(source string, val float64) {
	re.PersistsBackPressure.With(prometheus.Labels{"dataSource": source}).Observe(val)
}

// SetPersistsFailed .
func (re *IngestionRealtimeExporter) SetPersistsFailed(labels map[string]string, val float64) {
	re.PersistsFailed.With(labels).Add(val)
}

// SetHandOffFailed .
func (re *IngestionRealtimeExporter) SetHandOffFailed(labels map[string]string, val float64) {
	re.HandOffFailed.With(labels).Add(val)
}

// SetHandOffCount .
func (re *IngestionRealtimeExporter) SetHandOffCount(labels map[string]string, val float64) {
	re.HandOffCount.With(labels).Add(val)
}

// SetMergeTime .
func (re *IngestionRealtimeExporter) SetMergeTime(source string, val float64) {
	re.MergeTime.With(prometheus.Labels{"dataSource": source}).Observe(val)
}

// SetMergeCPU .
func (re *IngestionRealtimeExporter) SetMergeCPU(source string, val float64) {
	re.MergeCPU.With(prometheus.Labels{"dataSource": source}).Observe(val)
}

// SetSinkCount .
func (re *IngestionRealtimeExporter) SetSinkCount(labels map[string]string, val float64) {
	re.SinkCount.With(labels).Add(val)
}
