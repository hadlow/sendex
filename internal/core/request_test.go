package core

import (
	"net/http"
	"testing"
	// "github.com/google/go-cmp/cmp"
	// "github.com/hadlow/sendex/config"
)

func TestExecute(t *testing.T) {

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
