// Package middleware providing middlewares.
package middleware

import (
	"net/http"

	"github.com/lanz-dev/go-rest/wrapped"
)

// ShowError will allow wrapped.Response to show details in "message" for an error.
//
//  wrapped.Response{Code: http.StatusInternalServerError, Err: errors.New("myErrorMsg")}
//
//  // if ShowError(true)
//  {
//    "code": 500,
//    "status": "fail",
//    "message": "myErrorMsg",
//  }
//
// Wrapped response will check for an optional context key that ShowError middleware will set.
//
// Notice: The main purpose is to show details about errors in responses for a dev or internal environment.
func ShowError(showError bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(wrapped.CtxSetShowError(r.Context(), showError)))
		})
	}
}
