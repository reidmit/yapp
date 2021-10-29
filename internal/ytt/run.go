package ytt

import (
	"fmt"
	"io/ioutil"

	yttlib "github.com/k14s/ytt/pkg/cmd/template"
	"github.com/k14s/ytt/pkg/cmd/ui"
	"github.com/k14s/ytt/pkg/files"
	"github.com/reidmit/yapp/internal/config"
	"gopkg.in/yaml.v2"
)

const generatedDataValuesFile = "request-data-values.yml"
const generatedInputFile = "generated-app-config.yml"

func Run(
	appConfig *config.AppConfig,
	route config.HandledRoute,
	dataValues map[string]interface{},
) (*config.RouteConfig, error) {
	yttOptions := yttlib.NewOptions()

	yttOptions.Debug = appConfig.Debug
	yttOptions.DataValuesFlags.FromFiles = []string{generatedDataValuesFile}
	yttOptions.DataValuesFlags.ReadFileFunc = func(path string) ([]byte, error) {
		if path != generatedDataValuesFile {
			return nil, fmt.Errorf("unknown file to read: %s", path)
		}

		return yaml.Marshal(dataValues)
	}

	configFileBytes, _ := yaml.Marshal(appConfig)
	configFileBytes, err := ioutil.ReadFile(appConfig.Path)
	if err != nil {
		return nil, err
	}

	result := yttOptions.RunWithFiles(yttlib.Input{
		Files: []*files.File{
			files.MustNewFileFromSource(files.NewBytesSource(generatedInputFile, configFileBytes)),
		},
	}, ui.NewTTY(appConfig.Debug))

	if result.Err != nil {
		return nil, result.Err
	}

	if len(result.Files) == 0 {
		return nil, fmt.Errorf("expected to return an output file but saw zero files")
	}

	file := result.Files[0]
	if file.RelativePath() != generatedInputFile {
		return nil, fmt.Errorf("unexpected result file: %s", file.RelativePath())
	}

	var newAppConfig config.AppConfig
	yaml.Unmarshal(file.Bytes(), &newAppConfig)

	routeKey := fmt.Sprintf("%s %s", route.Method, route.Path)
	renderedRouteConfig := newAppConfig.Routes[routeKey]

	return &renderedRouteConfig, nil
}
