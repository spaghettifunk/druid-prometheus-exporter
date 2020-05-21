package collect

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spaghettifunk/druid-prometheus-exporter/pkg/export"
	"github.com/spaghettifunk/druid-prometheus-exporter/pkg/feed"
)

// DruidCollectMetrics receives the metrics from Druid and exports them in Prometheus
func DruidCollectMetrics(c *gin.Context) {
	var feeds []feed.Feed
	if err := c.ShouldBindJSON(&feeds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var sf feed.Service
	for _, f := range feeds {
		switch f.Service {
		case "druid/broker":
			sf = &feed.Broker{
				QueryBrokerExporter: c.MustGet("QueryBrokerExporter").(*export.QueryBrokerExporter),
				QueryCacheExporter:  c.MustGet("QueryCacheExporter").(*export.QueryCacheExporter),
				QueryJettyExporter:  c.MustGet("QueryJettyExporter").(*export.QueryJettyExporter),
				HealthJVMExporter:   c.MustGet("HealthJVMExporter").(*export.HealthJVMExporter),
				SQLExporter:         c.MustGet("SQLExporter").(*export.SQLExporter),
				SysExporter:         c.MustGet("SysExporter").(*export.SysExporter),
			}
			break
		case "druid/coordinator":
			sf = &feed.Coordinator{
				CoordinationExporter: c.MustGet("CoordinationExporter").(*export.CoordinationExporter),
				HealthJVMExporter:    c.MustGet("HealthJVMExporter").(*export.HealthJVMExporter),
				QueryJettyExporter:   c.MustGet("QueryJettyExporter").(*export.QueryJettyExporter),
			}
			break
		case "druid/historical":
			sf = &feed.Historical{
				QueryHistoricalExporter:  c.MustGet("QueryHistoricalExporter").(*export.QueryHistoricalExporter),
				HealthHistoricalExporter: c.MustGet("HealthHistoricalExporter").(*export.HealthHistoricalExporter),
				HealthJVMExporter:        c.MustGet("HealthJVMExporter").(*export.HealthJVMExporter),
				QueryCacheExporter:       c.MustGet("QueryCacheExporter").(*export.QueryCacheExporter),
				QueryJettyExporter:       c.MustGet("QueryJettyExporter").(*export.QueryJettyExporter),
				SQLExporter:              c.MustGet("SQLExporter").(*export.SQLExporter),
				SysExporter:              c.MustGet("SysExporter").(*export.SysExporter),
			}
			break
		case "druid/middleManager":
			sf = &feed.MiddleManager{
				HealthJVMExporter: c.MustGet("HealthJVMExporter").(*export.HealthJVMExporter),
			}
			break
		case "druid/router":
			sf = &feed.Router{
				HealthJVMExporter: c.MustGet("HealthJVMExporter").(*export.HealthJVMExporter),
			}
			break
		default:
			continue
		}
		sf.Evaluate(f)
	}
	c.JSON(http.StatusOK, "ok")
}
