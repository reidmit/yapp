package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

func runYTT(yttPath string, config *appConfig, dataValues map[string]interface{}) error {
	dataValuesBytes, _ := yaml.Marshal(dataValues)
	dataValuesYAML := "#@data/values\n---\n" +
		string(dataValuesBytes)

	cmd := exec.Command(yttPath, "-f", "-", "-f", config.Path)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdin = strings.NewReader(dataValuesYAML)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()

	if err != nil {
		fmt.Fprintln(os.Stderr, stderrBuf.String())
		return err
	}

	yaml.Unmarshal(stdoutBuf.Bytes(), &config)

	return nil
}
