package output

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/hadlow/sendex/config"
)

type OutputConfig struct {
	Request    *config.RequestSchema
	Path       string
	ShowStatus bool
	ShowHead   bool
	ShowBody   bool
}

func NewOutputConfig() *OutputConfig {
	c := OutputConfig{
		ShowStatus: true,
		ShowHead:   true,
		ShowBody:   true,
	}

	return &c
}

func Info(text string) {
	fmt.Println(Blue + text + Reset)
}

func Success(text string) {
	fmt.Println(Green + text + Reset)
}

func Warning(text string) {
	fmt.Println(Yellow + text + Reset)
}

func Error(err error) {
	fmt.Fprintf(os.Stderr, Red+"%v\n"+Reset, err)
}

func GenerateOutput(response *http.Response, config *OutputConfig, raw bool) (string, error) {
	buff := Buffer{
		raw: raw,
	}

	if config.ShowStatus {
		buff.Status(response)
	}

	if config.ShowHead {
		err := buff.Head(response, config.Request.WhitelistHeaders)
		if err != nil {
			return "", err
		}
	}

	if config.ShowBody {
		err := buff.Body(response)
		if err != nil {
			return "", err
		}
	}

	return buff.buffer.String(), nil
}

func Print(response *http.Response, config *OutputConfig) error {
	out, err := GenerateOutput(response, config, false)

	if err != nil {
		return err
	}

	if response.StatusCode >= 300 {
		fmt.Fprintf(os.Stderr, Red+"%v\n"+Reset, errors.New(out))

		return nil
	}

	fmt.Print(out)

	return nil
}

func Save(response *http.Response, config *OutputConfig) error {
	out, err := GenerateOutput(response, config, true)

	if err != nil {
		return err
	}

	f, err := os.Create(config.Path)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err2 := f.WriteString(out)

	if err2 != nil {
		return err2
	}

	Success("File created at " + config.Path)

	return nil
}
