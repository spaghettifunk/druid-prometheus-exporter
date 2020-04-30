package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	logDebugFlag = "debug"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "export",
	Short: "Export root command for initializing druid-exporter",
	Long: `The Druid Prometheus Exporter is capable of scraping the metrics produced
by Apache Druid and pack them in such a way that Prometheus is able to read. This is useful not only
for alarms and generic monitoring, but also for auto-scaling the various Apache Druid services.`,
}

// Execute will start the application
func Execute() {
	cobra.OnInitialize(initConfig)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

// initConfig sets AutomaticEnv in viper to true.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
