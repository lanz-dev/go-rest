# go-rest

[![Go Reference](https://pkg.go.dev/badge/github.com/lanz-dev/go-rest.svg)](https://pkg.go.dev/github.com/lanz-dev/go-rest)
[![Coverage Status](https://coveralls.io/repos/github/lanz-dev/go-rest/badge.svg?branch=main)](https://coveralls.io/github/lanz-dev/go-rest?branch=main)
[![Github Action](https://github.com/lanz-dev/go-mygrate/actions/workflows/main.yml/badge.svg)](https://github.com/lanz-dev/go-rest/actions/workflows/main.yml)


go-rest is inspired by the [REST API Tutorial](https://www.restapitutorial.com/resources.html). It will help to create a
RESTful service with stdlib in go. It's considered to be stable, feature complete, and it will not receive any breaking
changes.

One idea of the [REST API Tutorial](https://www.restapitutorial.com/resources.html) is to have "Wrapped Responses" (
see
[PDF](https://github.com/tfredrich/RestApiTutorial.com/raw/master/media/RESTful%20Best%20Practices-v1_2.pdf)
page 21).

```json
// Wrapped.Response

{
  "code": 200,
  "status": "success",
  "data": {

  }
}
```

This package provides methods and helpers to render wrapped.Responses and to test them.

```go
// Example Usage

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/lanz-dev/go-rest/handler"
	"github.com/lanz-dev/go-rest/rest"
)

func main() {
	c := chi.NewMux() // go-rest doesn't depend on chi!
	c.Route("/api", func(r chi.Router) {
		r.NotFound(handler.NotFoundHandler)
		r.MethodNotAllowed(handler.MethodNotAllowedHandler)

		r.Get("/demo", func(w http.ResponseWriter, r *http.Request) {
			yourStructOrNil := "demo"
			rest.Ok(w, r, yourStructOrNil)
		})
	})
}
```
