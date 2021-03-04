package rest

import (
	"net/http"

	"github.com/lanz-dev/go-rest/wrapped"
)

func responseWithData(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	Render(w, r, &wrapped.Response{Code: code, Data: data})
}

func responseWithMessage(w http.ResponseWriter, r *http.Request, code int, msg string) {
	Render(w, r, &wrapped.Response{Code: code, Message: msg})
}

func responseWithCode(w http.ResponseWriter, r *http.Request, code int) {
	Render(w, r, &wrapped.Response{Code: code})
}

// Error is a generic method to set an error on a wrapped.Response. The implementation
// will try to detect the statusCode, msg and data.
func Error(w http.ResponseWriter, r *http.Request, err error) {
	Render(w, r, &wrapped.Response{Err: err})
}

// 4xx

// BadRequest The server could not understand the request due to invalid syntax.
func BadRequest(w http.ResponseWriter, r *http.Request, msg string) {
	responseWithMessage(w, r, http.StatusBadRequest, msg)
}

// Unauthorized Although the HTTP standard specifies "unauthorized", semantically
// this response means "unauthenticated". That is, the client must authenticate itself
// to get the requested response.
func Unauthorized(w http.ResponseWriter, r *http.Request, msg string) {
	responseWithMessage(w, r, http.StatusUnauthorized, msg)
}

// Forbidden The client does not have access rights to the content; that is, it is
// unauthorized, so the server is refusing to give the requested resource. Unlike 401,
// the client's identity is known to the server.
func Forbidden(w http.ResponseWriter, r *http.Request, msg string) {
	responseWithMessage(w, r, http.StatusForbidden, msg)
}

// NotFound The server can not find the requested resource. In the browser, this means
// the URL is not recognized. In an API, this can also mean that the endpoint is valid
// but the resource itself does not exist. Servers may also send this response instead
// of 403 to hide the existence of a resource from an unauthorized client.
func NotFound(w http.ResponseWriter, r *http.Request, msg string) {
	responseWithMessage(w, r, http.StatusNotFound, msg)
}

// Conflict This response is sent when a request conflicts with the current state of
// the server.
func Conflict(w http.ResponseWriter, r *http.Request, msg string) {
	responseWithMessage(w, r, http.StatusConflict, msg)
}

// Gone This response is sent when the requested content has been permanently deleted
// from server, with no forwarding address. Clients are expected to remove their caches
// and links to the resource. The HTTP specification intends this status code to be
// used for "limited-time, promotional services". APIs should not feel compelled to
// indicate resources that have been deleted with this status code.
func Gone(w http.ResponseWriter, r *http.Request, msg string) {
	responseWithMessage(w, r, http.StatusGone, msg)
}

// UnsupportedMediaType The media format of the requested data is not supported by
// the server, so the server is rejecting the request.
func UnsupportedMediaType(w http.ResponseWriter, r *http.Request, msg string) {
	responseWithMessage(w, r, http.StatusUnsupportedMediaType, msg)
}

// TooManyRequests The user has sent too many requests in a given amount of time
// ("rate limiting").
func TooManyRequests(w http.ResponseWriter, r *http.Request, msg string) {
	responseWithMessage(w, r, http.StatusTooManyRequests, msg)
}

// UnavailableLegal The user-agent requested a resource that cannot legally be
// provided, such as a web page censored by a government.
func UnavailableLegal(w http.ResponseWriter, r *http.Request, msg string) {
	responseWithMessage(w, r, http.StatusUnavailableForLegalReasons, msg)
}

// 5xx

// InternalServerError The server has encountered a situation it doesn't know
// how to handle.
func InternalServerError(w http.ResponseWriter, r *http.Request, msg string) {
	responseWithMessage(w, r, http.StatusInternalServerError, msg)
}

// NotImplemented The request method is not supported by the server and cannot
// be handled. The only methods that servers are required to support (and
// therefore that must not return this code) are GET and HEAD.
func NotImplemented(w http.ResponseWriter, r *http.Request, msg string) {
	responseWithMessage(w, r, http.StatusNotImplemented, msg)
}

// ServiceUnavailable The server is not ready to handle the request. Common causes
// are a server that is down for maintenance or that is overloaded. Note that together
// with this response, a user-friendly page explaining the problem should be sent.
// This responses should be used for temporary conditions and the
// Retry-After: HTTP header should, if possible, contain the estimated time before the
// recovery of the service. The webmaster must also take care about the caching-related
// headers that are sent along with this response, as these temporary condition
// responses should usually not be cached.
func ServiceUnavailable(w http.ResponseWriter, r *http.Request, msg string) {
	responseWithMessage(w, r, http.StatusServiceUnavailable, msg)
}

// 2xx

// Ok The request has succeeded. The meaning of the success depends on the HTTP method:
//
// GET: The resource has been fetched and is transmitted in the message body.
//
// HEAD: The entity headers are in the message body.
//
// PUT or POST: The resource describing the result of the action is transmitted
// in the message body.
//
// TRACE: The message body contains the request message as received by the server.
func Ok(w http.ResponseWriter, r *http.Request, data interface{}) {
	responseWithData(w, r, http.StatusOK, data)
}

// Created The request has succeeded and a new resource has been created as a result.
// This is typically the response sent after POST requests, or some PUT requests.
func Created(w http.ResponseWriter, r *http.Request, data interface{}) {
	responseWithData(w, r, http.StatusCreated, data)
}

// Accepted The request has been received but not yet acted upon. It is noncommittal,
// since there is no way in HTTP to later send an asynchronous response indicating
// the outcome of the request. It is intended for cases where another process or server
// handles the request, or for batch processing.
func Accepted(w http.ResponseWriter, r *http.Request, data interface{}) {
	responseWithData(w, r, http.StatusAccepted, data)
}

// NoContent There is no content to send for this request, but the headers may be useful.
// The user-agent may update its cached headers for this resource with the new ones.
func NoContent(w http.ResponseWriter, r *http.Request) {
	responseWithCode(w, r, http.StatusNoContent)
}

// ResetContent Tells the user-agent to reset the document which sent this request.
func ResetContent(w http.ResponseWriter, r *http.Request) {
	responseWithCode(w, r, http.StatusResetContent)
}
