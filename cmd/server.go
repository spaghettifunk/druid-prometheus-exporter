package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/spaghettifunk/druid-prometheus-exporter/collect"
	"github.com/spaghettifunk/druid-prometheus-exporter/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

const (
	serverHostFlag = "server-host"
	serverPortFlag = "server-port"
)

// serverCmd represents the internal command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server for reading the metrics from Apche Druid",
	Long: `This command will start the server for polling the 
metrics from Apche Druid and convert them into the Prometheus format.`,
	Run: func(cmd *cobra.Command, args []string) {
		serverHost := viper.GetString(serverHostFlag)
		serverPort := viper.GetString(serverPortFlag)
		logDebug := viper.GetBool(logDebugFlag)

		// log level debug
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		if logDebug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}

		if err := startServer(serverHost, serverPort); err != nil {
			panic(err)
		}
	},
}

func startServer(addr, port string) error {
	r := gin.Default()
	// prometheus middleware
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	// inject all the middlewares based on the flags in input
	for _, md := range initMetrics() {
		r.Use(md)
	}

	// collect endpoint for receiving the metrics from Druid
	r.GET("/", collect.Healthz)
	r.POST("/collect", collect.DruidCollectMetrics)

	return r.Run(fmt.Sprintf("%s:%s", addr, port))
}

func initMetrics() []gin.HandlerFunc {
	var middlewares []gin.HandlerFunc
	middlewares = append(middlewares, middleware.QueryBroker(), middleware.QueryHistorical(),
		middleware.QueryJetty(), middleware.QueryRealtime(), middleware.QueryCache(), middleware.SQLMetrics(),
		middleware.IngestionKafka(), middleware.IngestionRealtime(), middleware.IngestionRealtimeIndexing(),
		middleware.Coordination(), middleware.HealthJVM(), middleware.HealthHistorical(),
		middleware.HealthFirehose(), middleware.SysMetrics())
	return middlewares
}

func init() {
	rootCmd.AddCommand(serverCmd)

	f := serverCmd.PersistentFlags()

	f.String(serverHostFlag, "localhost", "exporter server host")
	f.String(serverPortFlag, "7000", "exporter server port")

	viper.BindEnv(serverHostFlag, "EXPORTER_SERVER_HOST")
	viper.BindEnv(serverHostFlag, "EXPORTER_SERVER_PORT")

	viper.BindPFlags(f)
}
