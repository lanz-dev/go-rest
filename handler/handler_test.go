package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lanz-dev/go-rest/handler"
	"github.com/lanz-dev/go-rest/resttest"
)

func TestNotFoundHandler(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)
	w := httptest.NewRecorder()

	handler.NotFoundHandler(w, req)

	resp := w.Result()
	wrap := resttest.ParseToWrapped(t, resp.Body)

	resttest.ExpectStatusAndMessage(t, resp, wrap, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

func TestMethodNotAllowedHandler(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)
	w := httptest.NewRecorder()

	handler.MethodNotAllowedHandler(w, req)

	resp := w.Result()
	wrap := resttest.ParseToWrapped(t, resp.Body)

	resttest.ExpectStatusAndMessage(t, resp, wrap, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
}
