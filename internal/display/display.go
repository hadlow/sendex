package display

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tidwall/pretty"
)

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
	fmt.Fprintf(os.Stderr, Red+"error: %v\n"+Reset, err)
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

func YmlError() {

}

func Response(response *http.Response) error {
	if response.StatusCode < 200 {
		Warning(response.Status)
	}

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		Warning(response.Status)
	}

	if response.StatusCode >= 300 && response.StatusCode < 400 {
		Warning(response.Status)
	}

	if response.StatusCode >= 400 && response.StatusCode < 400 {
		Error(errors.New(response.Status))
	}

	if response.StatusCode >= 500 {
		Error(errors.New(response.Status))
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	prettyJSON := pretty.Color([]byte(body), nil)
	fmt.Println(string(prettyJSON[:]))

	return nil
}
