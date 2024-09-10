package request

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hadlow/sendex/config"
	"github.com/hadlow/sendex/internal/display"
	"github.com/hadlow/sendex/internal/file"
)

// get file contents, parse yaml, call endpoint, return raw response
func Run(path string) (*http.Response, error) {
	// get file contents
	contents, err := os.ReadFile(path)

	if err != nil {
		display.Error(err)
		return nil, fmt.Errorf("error reading file")
	}

	// parse yaml into object
	request, err := file.ParseYaml(contents)

	if err != nil {
		display.Error(err)
		return nil, fmt.Errorf("error parsing YAML")
	}

	// make request to endpoint
	response, err := execute(request)

	if err != nil {
		display.Error(err)
		return nil, fmt.Errorf("error calling endpoint")
	}

	return response, nil
}

func execute(request config.RequestSchema) (*http.Response, error) {
	// validate method
	method, err := getMethod(request.Method)

	if err != nil {
		log.Fatalf("invalid method: %s\n", method)
	}

	// create request
	req, err := http.NewRequest(method, request.Endpoint, strings.NewReader(request.Body))

	if err != nil {
		log.Fatalf("could not create request: %s\n", err)
	}

	// set headers
	setHeaders(request.Headers, req)

	res, err := http.DefaultClient.Do(req)

	return res, err
}

// get valid HTTP method, or throw error
func getMethod(method string) (string, error) {
	switch strings.ToLower(method) {
	case "get":
		return http.MethodGet, nil
	case "post":
		return http.MethodPost, nil
	case "put":
		return http.MethodPut, nil
	case "patch":
		return http.MethodPatch, nil
	case "delete":
		return http.MethodDelete, nil
	case "options":
		return http.MethodOptions, nil
	case "head":
		return http.MethodHead, nil
	default:
		return method, fmt.Errorf("HTTP method not valid")
	}
}

// set headers
func setHeaders(headers []map[string]string, req *http.Request) {
	for _, value := range headers {
		req.Header.Set(getKeyVal(value))
	}
}

func getKeyVal(m map[string]string) (string, string) {
	for k, v := range m {
		return k, v
	}

	return "", ""
}
