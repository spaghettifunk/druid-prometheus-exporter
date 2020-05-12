package export

import (
	"github.com/prometheus/client_golang/prometheus"
)

// QueryBrokerExporter contains all the Prometheus metrics that are possible to gather from the brokers
type QueryBrokerExporter struct {
	QueryTime             *prometheus.HistogramVec `description:"milliseconds taken to complete a query"`
	QueryBytes            *prometheus.HistogramVec `description:"number of bytes returned in query response"`
	QueryNodeTime         prometheus.Summary       `description:"milliseconds taken to query individual historical/realtime processes"`
	QueryNodeBytes        prometheus.Summary       `description:"number of bytes returned from querying individual historical/realtime processes"`
	QueryNodetTtfb        prometheus.Summary       `description:"time to first byte. Milliseconds elapsed until Broker starts receiving the response from individual historical/realtime processes"`
	QueryNodeBackpressure prometheus.Summary       `description:"milliseconds that the channel to this process has spent suspended due to backpressure"`
	QueryCount            *prometheus.GaugeVec     `description:"number of total queries"`
	QuerySuccessCount     *prometheus.GaugeVec     `description:"number of queries successfully processed"`
	QueryFailedCount      *prometheus.GaugeVec     `description:"number of failed queries"`
	QueryInterruptedCount *prometheus.GaugeVec     `description:"number of queries interrupted due to cancellation or timeout"`
	SQLQueryTime          prometheus.Summary       `description:"milliseconds taken to complete a SQL query"`
	SQLQueryBytes         prometheus.Summary       `description:"number of bytes returned in SQL query response"`
}

// NewQueryBrokerExporter returns a new broker exporter object
func NewQueryBrokerExporter() *QueryBrokerExporter {
	qb := &QueryBrokerExporter{
		QueryTime: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_time",
			Help:      "milliseconds taken to complete a query",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"dataSource"}),
		QueryBytes: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_bytes",
			Help:      "number of bytes returned in query response",
			Buckets:   []float64{10, 100, 500, 1000, 2000, 3000, 5000, 7000, 10000},
		}, []string{"dataSource"}),
		QueryNodeTime: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace:  "druid",
			Subsystem:  "broker",
			Name:       "query_node_time",
			Help:       "milliseconds taken to query individual historical/realtime processes",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}),
		QueryNodeBytes: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace:  "druid",
			Subsystem:  "broker",
			Name:       "query_node_bytes",
			Help:       "number of bytes returned from querying individual historical/realtime processes",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}),
		QueryNodetTtfb: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace:  "druid",
			Subsystem:  "broker",
			Name:       "query_node_ttfb",
			Help:       "time to first byte. Milliseconds elapsed until Broker starts receiving the response from individual historical/realtime processes",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}),
		QueryNodeBackpressure: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace:  "druid",
			Subsystem:  "broker",
			Name:       "query_node_backpressure",
			Help:       "milliseconds that the channel to this process has spent suspended due to backpressure",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}),
		QueryCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_count",
			Help:      "number of total queries",
		}, []string{}),
		QuerySuccessCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_success_count",
			Help:      "number of queries successfully processed",
		}, []string{}),
		QueryFailedCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_failed_count",
			Help:      "number of failed queries",
		}, []string{}),
		QueryInterruptedCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "druid",
			Subsystem: "broker",
			Name:      "query_interrupted_count",
			Help:      "number of queries interrupted due to cancellation or timeout",
		}, []string{}),
		SQLQueryTime: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace:  "druid",
			Subsystem:  "broker",
			Name:       "sql_query_time",
			Help:       "milliseconds taken to complete a SQL query",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}),
		SQLQueryBytes: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace:  "druid",
			Subsystem:  "broker",
			Name:       "sql_query_bytes",
			Help:       "number of bytes returned in SQL query response",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}),
	}

	// register all the prometheus metrics
	prometheus.MustRegister(qb.QueryTime)
	prometheus.MustRegister(qb.QueryBytes)
	prometheus.MustRegister(qb.QueryNodeTime)
	prometheus.MustRegister(qb.QueryNodeBytes)
	prometheus.MustRegister(qb.QueryNodetTtfb)
	prometheus.MustRegister(qb.QueryNodeBackpressure)
	prometheus.MustRegister(qb.QueryCount)
	prometheus.MustRegister(qb.QuerySuccessCount)
	prometheus.MustRegister(qb.QueryFailedCount)
	prometheus.MustRegister(qb.QueryInterruptedCount)
	prometheus.MustRegister(qb.SQLQueryTime)
	prometheus.MustRegister(qb.SQLQueryBytes)

	return qb
}

// SetQueryTime .
func (bc *QueryBrokerExporter) SetQueryTime(source string, val float64) {
	bc.QueryTime.With(prometheus.Labels{"dataSource": source}).Observe(val)
}

// SetQueryBytes .
func (bc *QueryBrokerExporter) SetQueryBytes(source string, val float64) {
	bc.QueryBytes.With(prometheus.Labels{"dataSource": source}).Observe(val)
}

// SetQueryNodeTime .
func (bc *QueryBrokerExporter) SetQueryNodeTime(val float64) {
	bc.QueryNodeTime.Observe(val)
}

// SetQueryNodeBytes .
func (bc *QueryBrokerExporter) SetQueryNodeBytes(val float64) {
	bc.QueryNodeBytes.Observe(val)
}

// SetQueryNodetTtfb .
func (bc *QueryBrokerExporter) SetQueryNodetTtfb(val float64) {
	bc.QueryNodetTtfb.Observe(val)
}

// SetQueryNodeBackpressure .
func (bc *QueryBrokerExporter) SetQueryNodeBackpressure(val float64) {
	bc.QueryNodeBackpressure.Observe(val)
}

// SetQueryCount .
func (bc *QueryBrokerExporter) SetQueryCount(val float64) {
	bc.QueryCount.WithLabelValues().Add(val)
}

// SetQuerySuccessCount .
func (bc *QueryBrokerExporter) SetQuerySuccessCount(val float64) {
	bc.QuerySuccessCount.WithLabelValues().Add(val)
}

// SetQueryFailedCount .
func (bc *QueryBrokerExporter) SetQueryFailedCount(val float64) {
	bc.QueryFailedCount.WithLabelValues().Add(val)
}

// SetQueryInterruptedCount .
func (bc *QueryBrokerExporter) SetQueryInterruptedCount(val float64) {
	bc.QueryInterruptedCount.WithLabelValues().Add(val)
}

// SetSQLQueryTime .
func (bc *QueryBrokerExporter) SetSQLQueryTime(val float64) {
	bc.SQLQueryTime.Observe(val)
}

// SetSQLQueryBytes .
func (bc *QueryBrokerExporter) SetSQLQueryBytes(val float64) {
	bc.SQLQueryBytes.Observe(val)
}
