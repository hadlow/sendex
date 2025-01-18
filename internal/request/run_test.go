package request

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hadlow/sendex/config"
)

func TestExecute(t *testing.T) {

}

func TestCompileId5(t *testing.T) {
	// Test with ID 5
	var request = config.RequestSchema{
		Args: []map[string]string{
			{
				"id": "1",
			},
		},
		Method:   "GET",
		Endpoint: "http://jsonplaceholder.typicode.com/todos/{id}",
		Headers: []map[string]string{
			{
				"TestId": "{id}",
			},
		},
	}

	var expectedRequestId5 = config.RequestSchema{
		Args: []map[string]string{
			{
				"id": "1",
			},
		},
		Method:   "GET",
		Endpoint: "http://jsonplaceholder.typicode.com/todos/5",
		Headers: []map[string]string{
			{
				"TestId": "5",
			},
		},
	}

	requestId5, err := compile(request, map[string]string{"id": "5"})

	if err != nil {
		t.Fatalf("error compiling request")
	}

	if !cmp.Equal(requestId5, expectedRequestId5) {
		t.Fatalf("requestId5 is not the same as expectedRequestId5")
	}
}

func TestCompileDefaultArg(t *testing.T) {
	var requestDefault = config.RequestSchema{
		Args: []map[string]string{
			{
				"id": "1",
			},
		},
		Method:   "GET",
		Endpoint: "http://jsonplaceholder.typicode.com/todos/{id}",
		Headers: []map[string]string{
			{
				"TestId": "{id}",
			},
		},
	}

	var expectedRequestId1 = config.RequestSchema{
		Args: []map[string]string{
			{
				"id": "1",
			},
		},
		Method:   "GET",
		Endpoint: "http://jsonplaceholder.typicode.com/todos/5",
		Headers: []map[string]string{
			{
				"TestId": "5",
			},
		},
	}

	requestId1, err := compile(requestDefault, map[string]string{"id": "5"})

	if err != nil {
		t.Fatalf("error compiling request")
	}

	if !cmp.Equal(requestId1, expectedRequestId1) {
		t.Fatalf("requestId1 is not the same as expectedRequestId5")
	}
}

func TestGetMethod(t *testing.T) {
	// get
	if m, _ := getMethod("get"); m != http.MethodGet {
		t.Fatalf("GET method not correct")
	}

	if m, _ := getMethod("GET"); m != http.MethodGet {
		t.Fatalf("GET method not correct")
	}

	if m, _ := getMethod("gEt"); m != http.MethodGet {
		t.Fatalf("GET method not correct")
	}

	// post
	if m, _ := getMethod("post"); m != http.MethodPost {
		t.Fatalf("POST method not correct")
	}

	if m, _ := getMethod("POST"); m != http.MethodPost {
		t.Fatalf("POST method not correct")
	}

	if m, _ := getMethod("pOsT"); m != http.MethodPost {
		t.Fatalf("POST method not correct")
	}

	// put
	if m, _ := getMethod("put"); m != http.MethodPut {
		t.Fatalf("PUT method not correct")
	}

	if m, _ := getMethod("PUT"); m != http.MethodPut {
		t.Fatalf("PUT method not correct")
	}

	if m, _ := getMethod("pUt"); m != http.MethodPut {
		t.Fatalf("PUT method not correct")
	}

	// patch
	if m, _ := getMethod("patch"); m != http.MethodPatch {
		t.Fatalf("PATCH method not correct")
	}

	if m, _ := getMethod("PATCH"); m != http.MethodPatch {
		t.Fatalf("PATCH method not correct")
	}

	if m, _ := getMethod("pAtCh"); m != http.MethodPatch {
		t.Fatalf("PATCH method not correct")
	}

	// delete
	if m, _ := getMethod("delete"); m != http.MethodDelete {
		t.Fatalf("DELETE method not correct")
	}

	if m, _ := getMethod("DELETE"); m != http.MethodDelete {
		t.Fatalf("DELETE method not correct")
	}

	if m, _ := getMethod("dElEtE"); m != http.MethodDelete {
		t.Fatalf("DELETE method not correct")
	}

	// options
	if m, _ := getMethod("options"); m != http.MethodOptions {
		t.Fatalf("OPTIONS method not correct")
	}

	if m, _ := getMethod("OPTIONS"); m != http.MethodOptions {
		t.Fatalf("OPTIONS method not correct")
	}

	if m, _ := getMethod("oPtIoNs"); m != http.MethodOptions {
		t.Fatalf("OPTIONS method not correct")
	}

	// head
	if m, _ := getMethod("head"); m != http.MethodHead {
		t.Fatalf("HEAD method not correct")
	}

	if m, _ := getMethod("HEAD"); m != http.MethodHead {
		t.Fatalf("HEAD method not correct")
	}

	if m, _ := getMethod("hEaD"); m != http.MethodHead {
		t.Fatalf("HEAD method not correct")
	}

	if _, err := getMethod("nope"); err == nil {
		t.Fatalf("no error when there should be")
	}
}

func TestSetHeaders(t *testing.T) {
	var headers = []map[string]string{
		{"Content-Type": "application/json"},
	}
	req, err := http.NewRequest(http.MethodGet, "", nil)

	if err != nil {
		t.Fatalf("could not create request")
	}

	setHeaders(headers, req)

	if req.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("header not set")
	}
}

func TestGetKeyVal(t *testing.T) {
	var m = map[string]string{"Content-Type": "application/json", "something": "else"}

	k, v := getKeyVal(m)

	if k != "Content-Type" {
		t.Fatalf("key incorrect")
	}

	if v != "application/json" {
		t.Fatalf("value incorrect")
	}
}
