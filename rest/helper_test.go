package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lanz-dev/go-rest/rest"
	"github.com/lanz-dev/go-rest/wrapped"
)

func TestHelper(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		code    int
		status  string
		data    interface{}
		msg     string
		handler http.HandlerFunc
	}{
		// 4xx
		"bad request": {
			http.StatusBadRequest, wrapped.StatusError, nil, "msg",
			func(w http.ResponseWriter, r *http.Request) {
				rest.BadRequest(w, r, "msg")
			},
		},
		"unauthorized": {
			http.StatusUnauthorized, wrapped.StatusError, nil, "msg",
			func(w http.ResponseWriter, r *http.Request) {
				rest.Unauthorized(w, r, "msg")
			},
		},
		"forbidden": {
			http.StatusForbidden, wrapped.StatusError, nil, "msg",
			func(w http.ResponseWriter, r *http.Request) {
				rest.Forbidden(w, r, "msg")
			},
		},
		"not found": {
			http.StatusNotFound, wrapped.StatusError, nil, "msg",
			func(w http.ResponseWriter, r *http.Request) {
				rest.NotFound(w, r, "msg")
			},
		},
		"conflict": {
			http.StatusConflict, wrapped.StatusError, nil, "msg",
			func(w http.ResponseWriter, r *http.Request) {
				rest.Conflict(w, r, "msg")
			},
		},
		"gone": {
			http.StatusGone, wrapped.StatusError, nil, "msg",
			func(w http.ResponseWriter, r *http.Request) {
				rest.Gone(w, r, "msg")
			},
		},
		"unsupported MediaType": {
			http.StatusUnsupportedMediaType, wrapped.StatusError, nil, "msg",
			func(w http.ResponseWriter, r *http.Request) {
				rest.UnsupportedMediaType(w, r, "msg")
			},
		},
		"too many requests": {
			http.StatusTooManyRequests, wrapped.StatusError, nil, "msg",
			func(w http.ResponseWriter, r *http.Request) {
				rest.TooManyRequests(w, r, "msg")
			},
		},
		"unavailable legal reasons": {
			http.StatusUnavailableForLegalReasons, wrapped.StatusError, nil, "msg",
			func(w http.ResponseWriter, r *http.Request) {
				rest.UnavailableLegal(w, r, "msg")
			},
		},
		// 5xx
		"internal server error": {
			http.StatusInternalServerError, wrapped.StatusFail, nil, "msg",
			func(w http.ResponseWriter, r *http.Request) {
				rest.InternalServerError(w, r, "msg")
			},
		},
		"not implemented": {
			http.StatusNotImplemented, wrapped.StatusFail, nil, "msg",
			func(w http.ResponseWriter, r *http.Request) {
				rest.NotImplemented(w, r, "msg")
			},
		},
		"service unavailable": {
			http.StatusServiceUnavailable, wrapped.StatusFail, nil, "msg",
			func(w http.ResponseWriter, r *http.Request) {
				rest.ServiceUnavailable(w, r, "msg")
			},
		},
		// 2xx
		"ok with data": {
			http.StatusOK, wrapped.StatusSuccess, "data", "",
			func(w http.ResponseWriter, r *http.Request) {
				rest.Ok(w, r, "data")
			},
		},
		"ok": {
			http.StatusOK, wrapped.StatusSuccess, nil, "",
			func(w http.ResponseWriter, r *http.Request) {
				rest.Ok(w, r, nil)
			},
		},
		"created with data": {
			http.StatusCreated, wrapped.StatusSuccess, "data", "",
			func(w http.ResponseWriter, r *http.Request) {
				rest.Created(w, r, "data")
			},
		},
		"created": {
			http.StatusCreated, wrapped.StatusSuccess, nil, "",
			func(w http.ResponseWriter, r *http.Request) {
				rest.Created(w, r, nil)
			},
		},
		"accepted": {
			http.StatusAccepted, wrapped.StatusSuccess, nil, "",
			func(w http.ResponseWriter, r *http.Request) {
				rest.Accepted(w, r, nil)
			},
		},
		"accepted with data": {
			http.StatusAccepted, wrapped.StatusSuccess, "data", "",
			func(w http.ResponseWriter, r *http.Request) {
				rest.Accepted(w, r, "data")
			},
		},
		"no content": {
			http.StatusNoContent, wrapped.StatusSuccess, nil, "",
			func(w http.ResponseWriter, r *http.Request) {
				rest.NoContent(w, r)
			},
		},
		"reset content": {
			http.StatusResetContent, wrapped.StatusSuccess, nil, "",
			func(w http.ResponseWriter, r *http.Request) {
				rest.ResetContent(w, r)
			},
		},
	}

	for name, tc := range tests {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest("GET", "/", nil)
			rr := httptest.NewRecorder()
			tc.handler(rr, req)
			res := parseBodyToResponse(t, rr.Body)

			if res.Code != tc.code {
				t.Fatalf(`expected Code to be '%d', got: '%d'`, tc.code, res.Code)
			}
			if res.Data != tc.data {
				t.Fatalf(`expected Data to be '%s', got: '%s'`, tc.data, res.Data)
			}
			if res.Message != tc.msg {
				t.Fatalf(`expected Message to be '%s', got: '%s'`, tc.msg, res.Message)
			}
			if res.Status != tc.status {
				t.Fatalf(`expected Status to be '%s', got: '%s'`, tc.status, res.Status)
			}
		})
	}
}
