package output

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/tidwall/pretty"

	"github.com/hadlow/sendex/config"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestResponse(t *testing.T) {
	expectedOutput := "\033[33m200 OK\033[0m\n" + Cyan + "Content-Type" + Reset + ": application/json\n" + string(pretty.Color([]byte("{\n  \"userId\": 1\n}\n"), nil)[:])

	response := http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: io.NopCloser(bytes.NewBufferString("{\n  \"userId\": 1\n}")),
	}

	outputConfig := NewOutputConfig()
	outputConfig.Request = &config.RequestSchema{}

	output := captureStdout(func() {
		Print(&response, outputConfig)
	})

	if output != expectedOutput {
		t.Fatalf("Terminal output not correct")
	}
}

func TestResponseWithOnlyStatus(t *testing.T) {
	expectedOutput := "\033[33m200 OK\033[0m\n"

	response := http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: io.NopCloser(bytes.NewBufferString("{\n  \"userId\": 1\n}")),
	}

	outputConfig := NewOutputConfig()
	outputConfig.ShowBody = false
	outputConfig.ShowHead = false
	outputConfig.Request = &config.RequestSchema{}

	output := captureStdout(func() {
		Print(&response, outputConfig)
	})

	if output != expectedOutput {
		t.Fatalf("Terminal output status output not correct")
	}
}

func TestResponseWithOnlyHeaders(t *testing.T) {
	expectedOutput := Cyan + "Content-Type" + Reset + ": application/json\n"

	response := http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: io.NopCloser(bytes.NewBufferString("{\n  \"userId\": 1\n}")),
	}

	outputConfig := NewOutputConfig()
	outputConfig.ShowStatus = false
	outputConfig.ShowBody = false
	outputConfig.Request = &config.RequestSchema{}

	output := captureStdout(func() {
		Print(&response, outputConfig)
	})

	if output != expectedOutput {
		t.Fatalf("Terminal output headers output not correct")
	}
}

func TestResponseWithWhitelistedHeaders(t *testing.T) {
	expectedOutput := Cyan + "Content-Type" + Reset + ": application/json\n"

	response := http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
			"Etag":         []string{"123"},
		},
		Body: io.NopCloser(bytes.NewBufferString("{\n  \"userId\": 1\n}")),
	}

	outputConfig := NewOutputConfig()
	outputConfig.ShowBody = false
	outputConfig.ShowStatus = false
	outputConfig.Request = &config.RequestSchema{
		WhitelistHeaders: []string{"Content-Type"},
	}

	output := captureStdout(func() {
		Print(&response, outputConfig)
	})

	if output != expectedOutput {
		t.Fatalf("Terminal output headers output not correct")
	}
}

func TestResponseWithOnlyBody(t *testing.T) {
	expectedOutput := string(pretty.Color([]byte("{\n  \"userId\": 1\n}\n"), nil)[:])

	response := http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: io.NopCloser(bytes.NewBufferString("{\n  \"userId\": 1\n}")),
	}

	outputConfig := NewOutputConfig()
	outputConfig.ShowStatus = false
	outputConfig.ShowHead = false
	outputConfig.Request = &config.RequestSchema{}

	output := captureStdout(func() {
		Print(&response, outputConfig)
	})

	if output != expectedOutput {
		t.Fatalf("Terminal output body output not correct")
	}
}

func TestResponseOnSave(t *testing.T) {
	expectedOutput := "200 OK\nContent-Type: application/json\n" + string(pretty.Color([]byte("{\n  \"userId\": 1\n}\n"), FileStyle)[:])

	response := http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: io.NopCloser(bytes.NewBufferString("{\n  \"userId\": 1\n}")),
	}

	outputConfig := NewOutputConfig()
	outputConfig.Request = &config.RequestSchema{}

	output, err := GenerateOutput(&response, outputConfig, true)

	if err != nil {
		t.Fatalf("Error generating raw output")
	}

	if output != expectedOutput {
		t.Fatalf("Raw output not correct")
	}
}
