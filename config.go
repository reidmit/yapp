package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type appConfig struct {
	Path   string
	Routes map[string]routeConfig
}

type routeConfig struct {
	Status *int
	Body   map[string]interface{}
}

func loadConfig(path string) (*appConfig, error) {
	yamlConfig, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	parsedConfig := appConfig{}

	err = yaml.Unmarshal(yamlConfig, &parsedConfig)
	if err != nil {
		return nil, err
	}

	parsedConfig.Path = path

	return &parsedConfig, nil
}
