// Package wrapped provides the wrapped.Response.
package wrapped

import (
	"context"
)

type ctxKey string

// DataResponder will set the Data field on Response.
type DataResponder interface {
	Data() interface{}
}

// MsgResponder will set the Message field on Response.
type MsgResponder interface {
	Message() string
}

// StatusCodeResponder will set the Status field on Response.
type StatusCodeResponder interface {
	StatusCode() int
}

const (
	// StatusFail will be set if StatusCode is 5XX.
	StatusFail = "fail"
	// StatusError will be set if StatusCode is 4XX.
	StatusError = "error"
	// StatusSuccess will be set if StatusCode is 1XX, 2XX or 3XX.
	StatusSuccess = "success"

	ctxKeyShowError = ctxKey("showError")
)

// CtxSetShowError will set ctxKeyShowError on ctx.
//
// If set to true, Response.Data will contain details about the error.
func CtxSetShowError(ctx context.Context, show bool) context.Context {
	return context.WithValue(ctx, ctxKeyShowError, show)
}

// ShowErrorFromCtx will get ctxKeyShowError on ctx.
func ShowErrorFromCtx(ctx context.Context) bool {
	showError, _ := ctx.Value(ctxKeyShowError).(bool)
	return showError
}
