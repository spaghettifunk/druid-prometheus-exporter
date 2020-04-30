package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"
)

// HealthFirehose is a middleware to initialise the Brokers exporter
func HealthFirehose() gin.HandlerFunc {
	b := export.NewHealthFirehoseExporter()
	return func(c *gin.Context) {
		c.Set("HealthFirehoseExporter", b)
		c.Next()
	}
}

// HealthHistorical is a middleware to initialise the Brokers exporter
func HealthHistorical() gin.HandlerFunc {
	b := export.NewHealthHistoricalExporter()
	return func(c *gin.Context) {
		c.Set("HealthHistoricalExporter", b)
		c.Next()
	}
}

// HealthJVM is a middleware to initialise the Brokers exporter
func HealthJVM() gin.HandlerFunc {
	b := export.NewHealthJVMExporter()
	return func(c *gin.Context) {
		c.Set("HealthJVMExporter", b)
		c.Next()
	}
}
