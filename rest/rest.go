// Package rest provides methods to render and to help implementing a REST api.
//
// The explanations for the helper methods are copied from https://developer.mozilla.org/en-US/docs/Web/HTTP/Status.
package rest

import (
	"encoding/json"
	"net/http"

	"github.com/lanz-dev/go-rest/wrapped"
)

// MarshalFn will be used to marshal wrapped.
//
// It defaults to json.Marshal and allows to change the marshaller.
var MarshalFn = json.Marshal

// Render will call Render on wrapped.Response to prepare the response and json.Marshal it to w.
func Render(w http.ResponseWriter, r *http.Request, wrapped *wrapped.Response) {
	wrapped.Parse(r.Context())

	data, err := MarshalFn(wrapped)
	if err != nil {
		Error(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(wrapped.Code)

	_, _ = w.Write(data)
}
