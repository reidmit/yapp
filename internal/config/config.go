package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Path   string
	Port   int64
	Debug  bool
	Routes map[string]RouteConfig
}

type RouteConfig struct {
	Status  *int
	Headers map[string][]string
	Body    interface{}
}

type HandledRoute struct {
	Method string
	Path   string
	Config RouteConfig
}

func Load(configPath string, expectedFileName string) (*AppConfig, error) {
	stats, err := os.Stat(configPath)
	if err != nil {
		return nil, err
	}

	if stats.IsDir() {
		configPath = filepath.Join(configPath, expectedFileName)
	}

	yamlConfig, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	log.Printf("Loading config file %s", configPath)

	parsedConfig := AppConfig{}

	err = yaml.Unmarshal(yamlConfig, &parsedConfig)
	if err != nil {
		return nil, err
	}

	parsedConfig.Path = configPath

	return &parsedConfig, nil
}

func GetHandledRoutes(routes map[string]RouteConfig) []HandledRoute {
	var handledRoutes []HandledRoute

	for routeWithMethod, routeConfig := range routes {
		parts := strings.Split(routeWithMethod, " ")

		handledRoutes = append(handledRoutes, HandledRoute{
			Method: parts[0],
			Path:   parts[1],
			Config: routeConfig,
		})
	}

	return handledRoutes
}
