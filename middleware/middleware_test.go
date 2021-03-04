package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lanz-dev/go-rest/middleware"
	"github.com/lanz-dev/go-rest/wrapped"
)

func TestShowError(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)
	rr := httptest.NewRecorder()

	middleware.ShowError(true)(func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if !wrapped.ShowErrorFromCtx(r.Context()) {
				t.Fatal(`expected ctxKeyShowError to be true`)
			}
		}
	}()).ServeHTTP(rr, req)
}
