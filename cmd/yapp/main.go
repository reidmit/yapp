package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/reidmit/yapp/internal/config"
	"github.com/reidmit/yapp/internal/server"
	"github.com/spf13/cobra"
)

const appName = "yapp"
const configFileName = "yapp.yml"
const defaultPort = 7000

// expected to be supplied at build time:
var appVersion string
var appCommit string

func main() {
	rootCmd := &cobra.Command{
		Use:     fmt.Sprintf("%s <path/to/%s>", appName, configFileName),
		Version: getAppVersion(),
		Short:   "Run your app",
		Long:    "Run your app using the given configuration file",
		Args:    cobra.ExactArgs(1),
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true,
			DisableDescriptions: true,
			DisableNoDescFlag:   true,
		},
		Run: run,
	}

	rootCmd.Flags().Int64(
		"port",
		getDefaultPort(),
		"port to listen on (can also set $PORT env var)",
	)

	rootCmd.Flags().Bool(
		"debug",
		false,
		"enable debug output",
	)

	rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) {
	configPath := args[0]

	appConfig, err := config.Load(configPath, configFileName)
	if err != nil {
		log.Printf("Error reading config: %v", err)
		os.Exit(1)
	}

	appConfig.Port, _ = cmd.Flags().GetInt64("port")
	appConfig.Debug, _ = cmd.Flags().GetBool("debug")

	if appConfig.Debug {
		log.Printf("Loaded config: %+v", appConfig)
	}

	server.Serve(appConfig)
}

func getAppVersion() string {
	var v string

	if appVersion != "" {
		v = appVersion
	} else {
		v = "0.0.0"
	}

	if appCommit != "" {
		v += "-" + appCommit
	} else {
		v += "-dev"
	}

	return v
}

func getDefaultPort() int64 {
	if portVar, isSet := os.LookupEnv("PORT"); isSet {
		port, _ := strconv.ParseInt(portVar, 10, 64)
		return port
	}

	return defaultPort
}
