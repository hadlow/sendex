package request

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hadlow/sendex/config"
	"github.com/hadlow/sendex/internal/helpers"
	"github.com/hadlow/sendex/internal/output"
)

// get file contents, parse yaml, call endpoint, return raw response
func Run(request *config.RequestSchema, args map[string]string) (*http.Response, error) {
	// replace args in request with values
	requestWithArgs := compile(request, args)

	// make request to endpoint
	response, err := execute(requestWithArgs, args)

	if err != nil {
		output.Error(err)
		return nil, fmt.Errorf("error calling endpoint")
	}

	return response, nil
}

func execute(request *config.RequestSchema, args map[string]string) (*http.Response, error) {
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

func compile(request *config.RequestSchema, args map[string]string) *config.RequestSchema {
	compiledRequest := request

	// override default args with CLI args
	finalArgs := helpers.FlattenMaps(request.Args)

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		finalArgs["env."+pair[0]] = pair[1]
	}

	for key, value := range args {
		if _, exists := finalArgs[key]; exists {
			finalArgs[key] = value
		}
	}

	// replace args
	for key, value := range finalArgs {
		compiledRequest.Endpoint = strings.Replace(compiledRequest.Endpoint, "{"+key+"}", value, -1)
		compiledRequest.Headers = replaceHeaders(compiledRequest.Headers, key, value)
		compiledRequest.Body = strings.Replace(compiledRequest.Body, "{"+key+"}", value, -1)
	}

	return compiledRequest
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
