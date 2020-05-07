package export

import "github.com/prometheus/client_golang/prometheus"

// SQLExporter contains all the Prometheus metrics that are possible to gather from the SQL service
type SQLExporter struct {
	SQLQueryTime  prometheus.Summary `description:"milliseconds taken to complete a SQL query"`
	SQLQueryBytes prometheus.Summary `description:"number of bytes returned in SQL query response"`
}

// NewSQLExporter returns a new SQL exporter object
func NewSQLExporter() *SQLExporter {
	qj := &SQLExporter{
		SQLQueryTime: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace:  "druid",
			Subsystem:  "sql",
			Name:       "query_time",
			Help:       "milliseconds taken to complete a SQL query",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}),
		SQLQueryBytes: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace:  "druid",
			Subsystem:  "sql",
			Name:       "query_bytes",
			Help:       "number of bytes returned in SQL query response",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}),
	}

	// register the prometheus metrics
	prometheus.MustRegister(qj.SQLQueryTime)
	prometheus.MustRegister(qj.SQLQueryBytes)

	return qj
}

// SetSQLQueryTime .
func (se *SQLExporter) SetSQLQueryTime(val float64) {
	se.SQLQueryTime.Observe(val)
}

// SetSQLQueryBytes .
func (se *SQLExporter) SetSQLQueryBytes(val float64) {
	se.SQLQueryBytes.Observe(val)
}
