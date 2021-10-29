package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Path   string
	Routes map[string]RouteConfig
}

type RouteConfig struct {
	Status *int
	Body   map[string]interface{}
}

func Load(path string) (*AppConfig, error) {
	yamlConfig, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	parsedConfig := AppConfig{}

	err = yaml.Unmarshal(yamlConfig, &parsedConfig)
	if err != nil {
		return nil, err
	}

	parsedConfig.Path = path

	return &parsedConfig, nil
}
