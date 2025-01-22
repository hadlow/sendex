package file

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hadlow/sendex/config"
)

func TestParseYaml(t *testing.T) {
	var contents = []byte(`args:
  - id: 1 # specify 1 as default
method: GET
endpoint: http://jsonplaceholder.typicode.com/todos/1 # we can use 'id' here
headers:
  - Content-Type: application/json
  - Accept: application/json
`)
	var expectedRequest = config.RequestSchema{
		Args: []map[string]string{
			{
				"id": "1",
			},
		},
		Method:   "GET",
		Endpoint: "http://jsonplaceholder.typicode.com/todos/1",
		Headers: []map[string]string{
			{
				"Content-Type": "application/json",
			},
			{
				"Accept": "application/json",
			},
		},
	}

	request, err := ParseYaml(contents)

	if err != nil {
		t.Fatalf("Error parsing YAML")
	}

	if !cmp.Equal(request, expectedRequest) {
		t.Fatalf("Not parsed properly")
	}
}
