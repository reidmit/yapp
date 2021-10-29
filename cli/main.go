package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

const (
	defaultPort    = "7000"
	defaultYttPath = "/usr/local/bin/ytt"
)

func main() {
	configPath := flag.String("f", "yapp.yml", "relative path to yapp.yml file")
	yttPath := flag.String("ytt", os.Getenv("YTT_PATH"), "absolute path to ytt executable (can also set YTT_PATH env var)")

	flag.Parse()

	config, err := loadConfig(*configPath)
	if err != nil {
		fmt.Printf("error reading config: %v", err)
		os.Exit(1)
	}

	if *yttPath == "" {
		*yttPath = defaultYttPath
	}

	setUpHandlers(*config, *yttPath)

	port := defaultPort
	if envPort, isSet := os.LookupEnv("PORT"); isSet {
		port = envPort
	}

	fmt.Printf("Listening on port %v...\n", port)

	http.ListenAndServe(":"+port, nil)
}
