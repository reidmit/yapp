package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

func runYTT(configPath string, config *appConfig, dataValues map[string]interface{}) {
	dataValuesBytes, _ := yaml.Marshal(dataValues)
	dataValuesYAML := fmt.Sprintf("#@data/values\n---\n%s", string(dataValuesBytes))

	cmd := exec.Command(defaultYttPath, "-f", "-", "-f", configPath)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdin = strings.NewReader(dataValuesYAML)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()

	if err != nil {
		log.Printf("ytt stderr: %v", stderrBuf.String())
		log.Fatalf("error running ytt: %v", err)
	}

	yaml.Unmarshal(stdoutBuf.Bytes(), &config)
}
