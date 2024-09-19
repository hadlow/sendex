package display

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/tidwall/pretty"
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

	config := NewDisplayConfig(true, true, true)

	output := captureStdout(func() {
		Response(&response, config)
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

	config := NewDisplayConfig(true, false, false)

	output := captureStdout(func() {
		Response(&response, config)
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

	config := NewDisplayConfig(false, true, false)

	output := captureStdout(func() {
		Response(&response, config)
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

	config := NewDisplayConfig(false, false, true)

	output := captureStdout(func() {
		Response(&response, config)
	})

	if output != expectedOutput {
		t.Fatalf("Display body output not correct")
	}
}
