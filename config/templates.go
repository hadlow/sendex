package config

var DefaultTemplate []byte = []byte("args:\n  - id: 1 # specify 1 as default\nmethod: GET\nendpoint: http://localhost:8000/blog/{id} # we can use 'id' here\nheaders:\n  - Content-Type: application/json\n  - Accept: application/json\nallow-headers:\n  - Content-Type\n")
var PostTemplate []byte = []byte("args:\n  - id: 1 # specify 1 as default\nmethod: POST\nendpoint: http://localhost:8000/blog/{id} # we can use 'id' here\nheaders:\n  - Content-Type: application/json\n  - Accept: application/json\nallow-headers:\n  - Content-Type\n")

type RequestSchema struct {
	Args         []map[string]string `yaml:"args"`
	Method       string              `yaml:"method"`
	Endpoint     string              `yaml:"endpoint"`
	Headers      []map[string]string `yaml:"headers"`
	Body         string              `yaml:"body"`
	AllowHeaders []string            `yaml:"allow-headers"`
}
