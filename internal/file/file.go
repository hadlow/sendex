package file

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/hadlow/sendex/config"
	"github.com/hadlow/sendex/internal/display"
)

func Get(path string) (*config.RequestSchema, error) {
	// get file contents
	contents, err := os.ReadFile(path)

	if err != nil {
		display.Error(err)
		return nil, fmt.Errorf("error reading file")
	}

	// parse yaml into object
	request, err := ParseYaml(contents)

	if err != nil {
		display.Error(err)
		return nil, fmt.Errorf("error parsing YAML")
	}

	return request, nil
}

func NewWithTemplate(path string, template []byte) error {
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("this file already exists")
	}

	err := os.WriteFile(path, template, 0644)

	return err
}

func ParseYaml(contents []byte) (*config.RequestSchema, error) {
	var request config.RequestSchema

	err := yaml.Unmarshal(contents, &request)

	if err != nil {
		return &request, err
	}

	return &request, nil
}

func Save(response *http.Response) error {
	return nil
}
