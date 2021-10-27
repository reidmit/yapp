package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type appConfig struct {
	Port   *int
	Routes map[string]routeConfig
}

type routeConfig struct {
	Status *int
	Body   map[string]interface{}
}

func loadConfig(path string) appConfig {
	yamlConfig, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	parsedConfig := appConfig{}

	err = yaml.Unmarshal(yamlConfig, &parsedConfig)
	if err != nil {
		log.Fatalf("error unmarshalling yapp.yml: %v", err)
	}

	return parsedConfig
}
