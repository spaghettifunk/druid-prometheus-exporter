package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"
)

// SysMetrics is a middleware to initialise the Brokers exporter
func SysMetrics() gin.HandlerFunc {
	b := export.NewSysExporter()
	return func(c *gin.Context) {
		c.Set("SysExporter", b)
		c.Next()
	}
}
