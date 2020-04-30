package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"
)

// SQLMetrics is a middleware to initialise the Brokers exporter
func SQLMetrics() gin.HandlerFunc {
	b := export.NewSQLExporter()
	return func(c *gin.Context) {
		c.Set("SQLExporter", b)
		c.Next()
	}
}
