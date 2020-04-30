package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"
)

// IngestionKafka is a middleware to initialise the Brokers exporter
func IngestionKafka() gin.HandlerFunc {
	b := export.NewIngestionKafkaExporter()
	return func(c *gin.Context) {
		c.Set("IngestionKafkaExporter", b)
		c.Next()
	}
}

// IngestionRealtime is a middleware to initialise the Brokers exporter
func IngestionRealtime() gin.HandlerFunc {
	b := export.NewIngestionRealtimeExporter()
	return func(c *gin.Context) {
		c.Set("IngestionRealtimeExporter", b)
		c.Next()
	}
}

// IngestionRealtimeIndexing is a middleware to initialise the Brokers exporter
func IngestionRealtimeIndexing() gin.HandlerFunc {
	b := export.NewIngestionRealtimeIndexingExporter()
	return func(c *gin.Context) {
		c.Set("IngestionRealtimeIndexingExporter", b)
		c.Next()
	}
}
