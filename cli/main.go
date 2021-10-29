package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/reidmit/yapp/config"
	"github.com/reidmit/yapp/server"
)

const (
	defaultPort    = "7000"
	defaultYttPath = "/usr/local/bin/ytt"
)

func main() {
	configPath := flag.String("f", "yapp.yml", "relative path to yapp.yml file")
	yttPath := flag.String("ytt", os.Getenv("YTT_PATH"), "absolute path to ytt executable (can also set YTT_PATH env var)")

	flag.Parse()

	config, err := config.Load(*configPath)
	if err != nil {
		fmt.Printf("error reading config: %v", err)
		os.Exit(1)
	}

	if *yttPath == "" {
		*yttPath = defaultYttPath
	}

	port := defaultPort
	if envPort, isSet := os.LookupEnv("PORT"); isSet {
		port = envPort
	}

	server.Serve(config, port, *yttPath)
}
