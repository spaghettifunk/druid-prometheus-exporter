package collect

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CommonMetric represents the common fields of the metric feed from Druid
type CommonMetric struct {
	Timestamp time.Time `json:"timestamp"`
	Metric    string    `json:"metric"`
	Service   string    `json:"service"`
	Host      string    `json:"host"`
	Value     int       `json:"value"`
}

// DruidCollectMetrics receives the metrics from Druid and exports them in Prometheus
func DruidCollectMetrics(c *gin.Context) {

	var cm CommonMetric
	if err := c.ShouldBindJSON(&cm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// channel the body to the appropriate exporter in a go func
	// parse the body in the struct that is more convenient
	// export to Prometheus

	c.JSON(http.StatusOK, "ok")
}
