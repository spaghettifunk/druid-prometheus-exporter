package export

import "github.com/prometheus/client_golang/prometheus"

// IngestionKafkaExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type IngestionKafkaExporter struct {
	Lag    *prometheus.GaugeVec `description:"total lag between the offsets consumed by the Kafka indexing tasks and latest offsets in Kafka brokers across all partitions. Minimum emission period for this metric is a minute"`
	MaxLag *prometheus.GaugeVec `description:"max lag between the offsets consumed by the Kafka indexing tasks and latest offsets in Kafka brokers across all partitions. Minimum emission period for this metric is a minute"`
	AVGLag *prometheus.GaugeVec `description:"average lag between the offsets consumed by the Kafka indexing tasks and latest offsets in Kafka brokers across all partitions. Minimum emission period for this metric is a minute"`
}

// NewIngestionKafkaExporter returns a new Jetty exporter object
func NewIngestionKafkaExporter() *IngestionKafkaExporter {
	ik := &IngestionKafkaExporter{
		Lag: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_kafka_lag",
			Help:      "total lag between the offsets consumed by the Kafka indexing tasks and latest offsets in Kafka brokers across all partitions. Minimum emission period for this metric is a minute",
		}, []string{"dataSource"}),
		MaxLag: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_kafka_max_lag",
			Help:      "max lag between the offsets consumed by the Kafka indexing tasks and latest offsets in Kafka brokers across all partitions. Minimum emission period for this metric is a minute",
		}, []string{"dataSource"}),
		AVGLag: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "realtime",
			Name:      "ingest_kafka_avg_lag",
			Help:      "average lag between the offsets consumed by the Kafka indexing tasks and latest offsets in Kafka brokers across all partitions. Minimum emission period for this metric is a minute",
		}, []string{"dataSource"}),
	}

	// register all the prometheus metrics
	prometheus.MustRegister(ik.Lag)
	prometheus.MustRegister(ik.MaxLag)
	prometheus.MustRegister(ik.AVGLag)

	return ik
}

// SetLag .
func (ik *IngestionKafkaExporter) SetLag(source string, val float64) {
	ik.Lag.With(prometheus.Labels{"dataSource": source}).Add(val)
}

// SetMaxLag .
func (ik *IngestionKafkaExporter) SetMaxLag(source string, val float64) {
	ik.MaxLag.With(prometheus.Labels{"dataSource": source}).Add(val)
}

// SetAVGLag .
func (ik *IngestionKafkaExporter) SetAVGLag(source string, val float64) {
	ik.AVGLag.With(prometheus.Labels{"dataSource": source}).Add(val)
}
