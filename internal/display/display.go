package display

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/tidwall/pretty"
)

type DisplayConfig struct {
	ShowStatus bool
	ShowHead   bool
	ShowBody   bool
}

func NewDisplayConfig(showStatus bool, showHead bool, showBody bool) *DisplayConfig {
	c := DisplayConfig{
		ShowStatus: showStatus,
		ShowHead:   showHead,
		ShowBody:   showBody,
	}

	return &c
}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"

func Plain(text string) {
	fmt.Println(text)
}

func Error(err error) {
	fmt.Fprintf(os.Stderr, Red+"%v\n"+Reset, err)
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

func HeaderItem(header string, value []string) {
	fmt.Println(Cyan + header + Reset + ": " + strings.Join(value, ", "))
}

func YmlError() {

}

func Response(response *http.Response, config *DisplayConfig) error {
	if config.ShowStatus {
		Status(response)
	}

	if config.ShowHead {
		err := Head(response)
		if err != nil {
			return err
		}
	}

	if config.ShowBody {
		err := Body(response)
		if err != nil {
			return err
		}
	}

	return nil
}

func Status(response *http.Response) {
	if response.StatusCode < 200 {
		Warning(response.Status)
	}

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		Warning(response.Status)
	}

	if response.StatusCode >= 300 && response.StatusCode < 400 {
		Warning(response.Status)
	}

	if response.StatusCode >= 400 && response.StatusCode < 500 {
		Error(errors.New(response.Status))
	}

	if response.StatusCode >= 500 {
		Error(errors.New(response.Status))
	}
}

func Head(response *http.Response) error {
	for header, value := range response.Header {
		HeaderItem(header, value)
	}

	return nil
}

func Body(response *http.Response) error {
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	prettyJSON := pretty.Color([]byte(body), nil)
	fmt.Println(string(prettyJSON[:]))

	return nil
}
