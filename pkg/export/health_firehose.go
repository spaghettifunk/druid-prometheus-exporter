package export

import "github.com/prometheus/client_golang/prometheus"

// HealthFirehoseExporter contains all the Prometheus metrics that are possible to gather from the Jetty service
type HealthFirehoseExporter struct {
	IngestEventsBuffered *prometheus.GaugeVec `description:"number of events queued in the EventReceiverFirehose's buffer"`
	IngestBytesReceived  *prometheus.GaugeVec `description:"number of bytes received by the EventReceiverFirehose"`
}

// NewHealthFirehoseExporter returns a new Jetty exporter object
func NewHealthFirehoseExporter() *HealthFirehoseExporter {
	hf := &HealthFirehoseExporter{
		IngestEventsBuffered: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "firehose",
			Name:      "events_buffered",
			Help:      "number of events queued in the EventReceiverFirehose's buffer",
		}, []string{"serviceName", "dataSource", "taskId", "taskType", "bufferCapacity"}),
		IngestBytesReceived: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "firehose",
			Name:      "bytes_received",
			Help:      "number of bytes received by the EventReceiverFirehose",
		}, []string{"serviceName", "dataSource", "taskId", "taskType"}),
	}

	// register all the prometheus metrics
	prometheus.MustRegister(hf.IngestEventsBuffered)
	prometheus.MustRegister(hf.IngestBytesReceived)

	return hf
}

// SetIngestEventsBuffered .
func (hf *HealthFirehoseExporter) SetIngestEventsBuffered(labels map[string]string, val float64) {
	hf.IngestEventsBuffered.With(labels).Add(val)
}

// SetIngestBytesReceived .
func (hf *HealthFirehoseExporter) SetIngestBytesReceived(labels map[string]string, val float64) {
	hf.IngestBytesReceived.With(labels).Add(val)
}
