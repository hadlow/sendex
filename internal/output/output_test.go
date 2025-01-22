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

	displayConfig := NewDisplayConfig(true, true, true)
	displayConfig.Request = &config.RequestSchema{}

	output := captureStdout(func() {
		Response(&response, displayConfig)
	})

	if output != expectedOutput {
		t.Fatalf("Display response output not correct")
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

	displayConfig := NewDisplayConfig(true, false, false)
	displayConfig.Request = &config.RequestSchema{}

	output := captureStdout(func() {
		Response(&response, displayConfig)
	})

	if output != expectedOutput {
		t.Fatalf("Display status output not correct")
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

	displayConfig := NewDisplayConfig(false, true, false)
	displayConfig.Request = &config.RequestSchema{}

	output := captureStdout(func() {
		Response(&response, displayConfig)
	})

	if output != expectedOutput {
		t.Fatalf("Display headers output not correct")
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

	displayConfig := NewDisplayConfig(false, true, false)
	displayConfig.Request = &config.RequestSchema{
		WhitelistHeaders: []string{"Content-Type"},
	}

	output := captureStdout(func() {
		Response(&response, displayConfig)
	})

	if output != expectedOutput {
		t.Fatalf("Display headers output not correct")
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

	displayConfig := NewDisplayConfig(false, false, true)
	displayConfig.Request = &config.RequestSchema{}

	output := captureStdout(func() {
		Response(&response, displayConfig)
	})

	if output != expectedOutput {
		t.Fatalf("Display body output not correct")
	}
}
