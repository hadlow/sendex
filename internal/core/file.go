package core

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/hadlow/sendex/config"
)

func NewWithTemplate(path string, template []byte) error {
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("this file already exists")
	}

	err := os.WriteFile(path, template, 0644)

	return err
}

func ParseYaml(contents []byte) (config.RequestSchema, error) {
	var request config.RequestSchema

	err := yaml.Unmarshal(contents, &request)

	if err != nil {
		return request, err
	}

	return request, nil
}
