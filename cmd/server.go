package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spaghettifunk/druid-prometheus-exporter/collect"
	"github.com/spaghettifunk/druid-prometheus-exporter/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

const (
	serverURLFlag = "server-url"
)

var (
	queryBrokerMetricsFlag                        = false
	queryHistoricalMetricsFlag                    = false
	queryJettyMetricsFlag                         = false
	queryRealtimeMetricsFlag                      = false
	queryCacheMetricsFlag                         = false
	sqlMetricsFlag                                = false
	ingestionKafkaMetricsFlag                     = false
	ingestionRealtimeMetricsFlag                  = false
	ingestionRealtimeIndexingMetricsFlag          = false
	coordinationMetricsFlag                       = false
	generalHealthHistoricalMetricsFlag            = false
	generalHealthJVMMetricsFlag                   = false
	generalHealthEventReceiverFirehoseMetricsFlag = false
	sysMetricsFlags                               = false
)

// serverCmd represents the internal command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server for reading the metrics from Apche Druid",
	Long: `This command will start the server for polling the 
metrics from Apche Druid and convert them into the Prometheus format.`,
	Run: func(cmd *cobra.Command, args []string) {
		serverURL := viper.GetString(serverURLFlag)
		logDebug := viper.GetBool(logDebugFlag)

		// log level debug
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		if logDebug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}

		if err := startServer(serverURL); err != nil {
			panic(err)
		}
	},
}

func startServer(addr string) error {
	r := gin.Default()
	// prometheus middleware
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	// inject all the middlewares based on the flags in input
	for _, md := range initMetrics() {
		r.Use(md)
	}

	// collect endpoint for receiving the metrics from Druid
	r.POST("/collect", collect.DruidCollectMetrics)

	return r.Run(addr)
}

func initMetrics() []gin.HandlerFunc {
	var middlewares []gin.HandlerFunc
	// query middlewares
	if queryBrokerMetricsFlag {
		middlewares = append(middlewares, middleware.QueryBroker())
	}
	if queryHistoricalMetricsFlag {
		middlewares = append(middlewares, middleware.QueryHistorical())
	}
	if queryJettyMetricsFlag {
		middlewares = append(middlewares, middleware.QueryJetty())
	}
	if queryRealtimeMetricsFlag {
		middlewares = append(middlewares, middleware.QueryRealtime())
	}
	if queryCacheMetricsFlag {
		middlewares = append(middlewares, middleware.QueryCache())
	}
	// sql middleware
	if sqlMetricsFlag {
		middlewares = append(middlewares, middleware.SQLMetrics())
	}
	// kafka ingestion middleware
	if ingestionKafkaMetricsFlag {
		middlewares = append(middlewares, middleware.IngestionKafka())
	}
	// realtime ingestion middlewares
	if ingestionRealtimeMetricsFlag {
		middlewares = append(middlewares, middleware.IngestionRealtime())
	}
	if ingestionRealtimeIndexingMetricsFlag {
		middlewares = append(middlewares, middleware.IngestionRealtimeIndexing())
	}
	// coordinator middleware
	if coordinationMetricsFlag {
		middlewares = append(middlewares, middleware.Coordination())
	}
	// general health middlewares
	if generalHealthJVMMetricsFlag {
		middlewares = append(middlewares, middleware.HealthJVM())
	}
	if generalHealthHistoricalMetricsFlag {
		middlewares = append(middlewares, middleware.HealthHistorical())
	}
	if generalHealthEventReceiverFirehoseMetricsFlag {
		middlewares = append(middlewares, middleware.HealthFirehose())
	}
	// sys middleware
	if sysMetricsFlags {
		middlewares = append(middlewares, middleware.SysMetrics())
	}
	return middlewares
}

func init() {
	rootCmd.AddCommand(serverCmd)

	f := serverCmd.PersistentFlags()

	f.String(serverURLFlag, ":7000", "exporter server url")

	// optional query metrics flags - ref: https://druid.apache.org/docs/latest/operations/metrics.html#query-metrics
	f.BoolVarP(&queryBrokerMetricsFlag, "query-broker", "b", false, "export the query brokers metrics")
	f.BoolVarP(&queryHistoricalMetricsFlag, "query-historical", "i", false, "export the query historical metrics")
	f.BoolVarP(&queryJettyMetricsFlag, "query-jetty", "j", false, "export the query jetty metrics")
	f.BoolVarP(&queryRealtimeMetricsFlag, "query-realtime", "r", false, "export the query tranquillity metrics")
	f.BoolVarP(&queryCacheMetricsFlag, "query-cache", "c", false, "export the query cache metrics")

	// optional sql metrics flag - ref: https://druid.apache.org/docs/latest/operations/metrics.html#sql-metrics
	f.BoolVarP(&sqlMetricsFlag, "sql", "s", false, "export the SQL metrics")

	// optional kafka ingestion metrics flag - ref: https://druid.apache.org/docs/latest/operations/metrics.html#ingestion-metrics-kafka-indexing-service
	f.BoolVarP(&ingestionKafkaMetricsFlag, "kafka-ingestion", "k", false, "export the query cache metrics")

	// optional realtime ingestion metrics flags - ref: https://druid.apache.org/docs/latest/operations/metrics.html#ingestion-metrics-realtime-process
	f.BoolVarP(&ingestionRealtimeMetricsFlag, "realtime-ingestion", "t", false, "export the query cache metrics")
	f.BoolVarP(&ingestionRealtimeIndexingMetricsFlag, "realtime-indexing", "n", false, "export the query cache metrics")

	// optional coordinator metrics flag - ref: https://druid.apache.org/docs/latest/operations/metrics.html#coordination
	f.BoolVarP(&coordinationMetricsFlag, "coordination", "o", false, "export the query cache metrics")

	// optional general health metrics flag - ref: https://druid.apache.org/docs/latest/operations/metrics.html#general-health
	f.BoolVarP(&generalHealthJVMMetricsFlag, "health-jvm", "v", false, "export the query cache metrics")
	f.BoolVarP(&generalHealthHistoricalMetricsFlag, "health-historical", "l", false, "export the query cache metrics")
	f.BoolVarP(&generalHealthEventReceiverFirehoseMetricsFlag, "health-firehose", "f", false, "export the query cache metrics")

	// optional Sys metrics flag - ref: https://druid.apache.org/docs/latest/operations/metrics.html#sys
	f.BoolVarP(&sysMetricsFlags, "sys", "g", false, "export the query cache metrics")

	viper.BindEnv(serverURLFlag, "EXPORTER_SERVER_URL")
	viper.BindPFlags(f)
}
