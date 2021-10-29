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

const debug = false
const generatedDataValuesFile = "request-data-values.yml"
const generatedInputFile = "yapp.yml"

func Run(appConfig *config.AppConfig, dataValues map[string]interface{}) error {
	yttOptions := yttlib.NewOptions()

	yttOptions.Debug = debug
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
		return err
	}

	result := yttOptions.RunWithFiles(yttlib.Input{
		Files: []*files.File{
			files.MustNewFileFromSource(files.NewBytesSource(generatedInputFile, configFileBytes)),
		},
	}, ui.NewTTY(debug))

	if result.Err != nil {
		return result.Err
	}

	if len(result.Files) == 0 {
		return fmt.Errorf("expected to return an output file but saw zero files")
	}

	file := result.Files[0]
	if file.RelativePath() != generatedInputFile {
		return fmt.Errorf("unexpected result file: %s", file.RelativePath())
	}

	yaml.Unmarshal(file.Bytes(), &appConfig)

	return nil
}
