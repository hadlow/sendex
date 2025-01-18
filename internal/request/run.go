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
	"github.com/hadlow/sendex/internal/helpers"
)

// get file contents, parse yaml, call endpoint, return raw response
func Run(path string, args map[string]string) (*http.Response, error) {
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

	fmt.Println("request")
	fmt.Println(request)

	// replace args in request with values
	requestWithArgs, err := compile(request, args)

	fmt.Println("requestWithArgs")
	fmt.Println(requestWithArgs)

	if err != nil {
		display.Error(err)
		return nil, fmt.Errorf("error compiling request")
	}

	// make request to endpoint
	response, err := execute(requestWithArgs, args)

	if err != nil {
		display.Error(err)
		return nil, fmt.Errorf("error calling endpoint")
	}

	return response, nil
}

func execute(request config.RequestSchema, args map[string]string) (*http.Response, error) {
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

	setHeaders(request.Headers, req)

	res, err := http.DefaultClient.Do(req)

	return res, err
}

func compile(request config.RequestSchema, args map[string]string) (config.RequestSchema, error) {
	compiledRequest := request

	// override default args with CLI args
	finalArgs := helpers.FlattenMaps(request.Args)

	for key, value := range args {
		if _, exists := finalArgs[key]; exists {
			finalArgs[key] = value
		}
	}

	// replace args
	for a, b := range finalArgs {
		compiledRequest.Endpoint = strings.Replace(compiledRequest.Endpoint, "{"+a+"}", b, -1)
		compiledRequest.Headers = replaceHeaders(compiledRequest.Headers, a, b)
		compiledRequest.Body = strings.Replace(compiledRequest.Body, "{"+a+"}", b, -1)
	}

	return compiledRequest, nil
}

func replaceHeaders(headers []map[string]string, a string, b string) []map[string]string {
	for index, header := range headers {
		for ha, hb := range header {
			headers[index] = map[string]string{ha: strings.Replace(hb, "{"+a+"}", b, -1)}
		}
	}

	return headers
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
