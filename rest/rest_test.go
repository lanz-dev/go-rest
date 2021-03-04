package rest_test

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lanz-dev/go-rest/rest"
	"github.com/lanz-dev/go-rest/wrapped"
)

func parseBodyToResponse(t *testing.T, r io.Reader) wrapped.Response {
	var res wrapped.Response
	if err := json.NewDecoder(r).Decode(&res); err != nil {
		t.Fatalf("%s: could not parse body to Response, err: '%s'", t.Name(), err)
	}
	return res
}

func TestRender(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)
	rr := httptest.NewRecorder()

	expected := "ok."
	rest.Render(rr, req, &wrapped.Response{Data: expected})
	res := parseBodyToResponse(t, rr.Body)

	if res.Code != http.StatusOK {
		t.Fatalf(`expected Status to be '%d', got: '%d'`, http.StatusOK, res.Code)
	}

	data, _ := res.Data.(string)
	if data != expected {
		t.Fatalf(`expected Data to be '%s', got: '%s'`, expected, res.Data)
	}
}

func TestRender_Error(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)
	rr := httptest.NewRecorder()

	rest.Render(rr, req, &wrapped.Response{Data: math.Inf(1)})
	res := parseBodyToResponse(t, rr.Body)

	if res.Code != http.StatusInternalServerError {
		t.Fatalf(`expected Status to be '%d', got: '%d'`, http.StatusInternalServerError, res.Code)
	}
}
