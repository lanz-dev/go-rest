// Package handler provides default handlers.
//
// The handlers response with a wrapped.Response.
//
// 	// Example with chi
//
// 	c := chi.NewMux()
//	c.Route("/api", func(r chi.Router) {
//	    r.NotFound(handler.NotFoundHandler)
//	    r.MethodNotAllowed(handler.MethodNotAllowedHandler)
//
//	    r.Get("/service", func(w http.ResponseWriter, r *http.Request) {
//
//	    })
//	})
package handler

import (
	"net/http"

	"github.com/lanz-dev/go-rest/rest"
	"github.com/lanz-dev/go-rest/wrapped"
)

// NotFoundHandler will render a wrapped.Response with http.StatusNotFound.
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	rest.Render(w, r, &wrapped.Response{Code: http.StatusNotFound})
}

// MethodNotAllowedHandler will render a wrapped.Response with http.StatusMethodNotAllowed.
func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	rest.Render(w, r, &wrapped.Response{Code: http.StatusMethodNotAllowed})
}
