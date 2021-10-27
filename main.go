package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort    = "7000"
	defaultYttPath = "/usr/local/bin/ytt"
)

func main() {
	configPath := flag.String("f", "yapp.yml", "path to yapp.yml file")

	flag.Parse()

	config := loadConfig(*configPath)

	setUpHandlers(config, *configPath)

	port := defaultPort
	if envPort, isSet := os.LookupEnv("PORT"); isSet {
		port = envPort
	}

	log.Printf("Listening on port %v...\n", port)

	http.ListenAndServe(":"+port, nil)
}
