package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"
)

// QueryBroker is a middleware to initialise the Brokers exporter
func QueryBroker() gin.HandlerFunc {
	b := export.NewQueryBrokerExporter()
	return func(c *gin.Context) {
		c.Set("QueryBrokerExporter", b)
		c.Next()
	}
}

// QueryHistorical is a middleware to initialise the Historical exporter
func QueryHistorical() gin.HandlerFunc {
	h := export.NewQueryHistoricalExporter()
	return func(c *gin.Context) {
		c.Set("QueryHistoricalExporter", h)
		c.Next()
	}
}

// QueryRealtime is a middleware to initialise the Tranquillity exporter
func QueryRealtime() gin.HandlerFunc {
	r := export.NewQueryRealtimeExporter()
	return func(c *gin.Context) {
		c.Set("QueryRealtimeExporter", r)
		c.Next()
	}
}

// QueryJetty is a middleware to initialise the Tranquillity exporter
func QueryJetty() gin.HandlerFunc {
	j := export.NewQueryJettyExporter()
	return func(c *gin.Context) {
		c.Set("QueryJettyExporter", j)
		c.Next()
	}
}

// QueryCache is a middleware to initialise the Tranquillity exporter
func QueryCache() gin.HandlerFunc {
	j := export.NewQueryCacheExporter()
	return func(c *gin.Context) {
		c.Set("QueryCacheExporter", j)
		c.Next()
	}
}
