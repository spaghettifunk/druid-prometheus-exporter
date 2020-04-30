package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"
)

// Coordination is a middleware to initialise the Brokers exporter
func Coordination() gin.HandlerFunc {
	b := export.NewCoordinationExporter()
	return func(c *gin.Context) {
		c.Set("CoordinationExporter", b)
		c.Next()
	}
}
