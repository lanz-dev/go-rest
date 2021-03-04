// Package resttest provides helper methods for testing.
package resttest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lanz-dev/go-rest/wrapped"
)

// Request will build a REST http request.
func Request(t *testing.T, method, url string, body interface{}) *http.Request {
	j, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("%s: error with marshalling body (%s)", t.Name(), err)
	}
	req := httptest.NewRequest(method, url, bytes.NewReader(j))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	return req
}

// ParseToWrapped will parse r into wrapped.Response.
func ParseToWrapped(t *testing.T, r io.Reader) wrapped.Response {
	var res wrapped.Response
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		t.Fatalf("%s: could not parse body to Response, err: '%s'", t.Name(), err)
	}
	return res
}

// ExpectStatusCode expects statusCode in http.Response and in wrapped.Response.
func ExpectStatusCode(t *testing.T, resp *http.Response, wrap wrapped.Response, code int) {
	if resp.StatusCode != code {
		t.Fatalf("%s: expected StatusCode (in response) '%d', got '%d'", t.Name(), code, resp.StatusCode)
	}

	if wrap.Code != code {
		t.Fatalf("%s: expected StatusCode (in wrapped) '%d', got '%d'", t.Name(), code, wrap.Code)
	}
}

// ExpectMessage expects message in wrapped.Response.
func ExpectMessage(t *testing.T, wrap wrapped.Response, msg string) {
	if wrap.Message != msg {
		t.Fatalf("%s: expected Message '%s', got '%s'", t.Name(), msg, wrap.Message)
	}
}

// ExpectStatusAndMessage expects statusCode statusCode in http.Response and in wrapped.Response and message in wrapped.Response.
func ExpectStatusAndMessage(t *testing.T, resp *http.Response, wrap wrapped.Response, code int, msg string) {
	ExpectStatusCode(t, resp, wrap, code)
	ExpectMessage(t, wrap, msg)
}
