package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/reidmit/yapp/internal/config"
	"github.com/reidmit/yapp/internal/server"
	"github.com/spf13/cobra"
)

const appName = "yapp"
const appVersion = "0.0.1"
const yttPath = "/usr/local/bin/ytt"

func main() {
	rootCmd := &cobra.Command{
		Use:     appName,
		Version: appVersion,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd:   true,
			DisableDescriptions: true,
			DisableNoDescFlag:   true,
		},
	}

	runCmd := &cobra.Command{
		Use:   "run <path/to/yapp.yml>",
		Short: "run your yapp",
		Long:  "run your yapp using the given yapp.yml file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configPath := args[0]
			port, _ := cmd.Flags().GetInt64("port")

			appConfig, err := config.Load(configPath)
			if err != nil {
				fmt.Printf("error reading config: %v", err)
				os.Exit(1)
			}

			server.Serve(appConfig, port, yttPath)
		},
	}

	runCmd.PersistentFlags().Int64(
		"port",
		getDefaultPort(),
		"port to listen on (can also set $PORT env var)",
	)

	rootCmd.AddCommand(runCmd)

	rootCmd.Execute()
}

func getDefaultPort() int64 {
	if portVar, isSet := os.LookupEnv("PORT"); isSet {
		port, _ := strconv.ParseInt(portVar, 10, 64)
		return port
	}

	return 7000
}
