package wrapped

import (
	"context"
	"errors"
	"net/http"
)

// Response represents a response from a rest endpoint.
//
// The idea comes from https://www.restapitutorial.com/resources.html
//
// Details
//
// Code value will determined base on the following order:
//  - If field is already set, the already set value will be used
//  - If Err will implement StatusCodeResponder the result will be used
//  - If Err != nil the value will be http.StatusInternalServerError
//  - Else the value will be http.StatusSuccess
//
// Status will be one of
//  - “fail” for HTTP status code response values from 500-599,
//  - “error” for HTTP status code response values from 400-499,
//  - “success” for everything else (e.g. 1XX, 2XX and 3XX responses).
//
// Message value will be determined based on the following order:
//  - If status is "success" this field will always be "" and not present in json
//  - If the field is already set, the already set value will be used
//  - If status is not "success" and Err will implement MsgResponder, this field will be MsgResponder.ResponseMessage()
//  - If status is "error" this will be Err.Error()
//  - If status is "fail" and ctxKeyShowError is true this will be Err.Error()
//  - Else it will contain http.StatusText(Code)
//
// Data value will be determined base on the following order:
//  - if status is "success" it could contain the response body
//  - If the field is already set, the already set value will be used
//  - If the status is not "success" it could contain the cause/exception (depends on ShowErrorFromCtx).
type Response struct {
	// Code contains the HTTP response status code as an integer.
	Code int `json:"code"`
	// Status contains the text: “success”, “fail”, or “error”.
	Status string `json:"status"`
	// Message is only used for “fail” and “error” statuses to contain the error message.
	//
	// For internationalization (i18n) purposes, this could contain a message number or code,
	// either alone or contained within delimiters.
	Message string `json:"message,omitempty"`
	// Data can contain user provides data,
	Data interface{} `json:"data,omitempty"`
	// Err contains the error
	Err error `json:"-"`
}

func (res *Response) setCode() {
	if res.Code != 0 {
		return
	}

	var responder StatusCodeResponder
	if errors.As(res.Err, &responder) {
		res.Code = responder.StatusCode()
		return
	}

	if res.Err != nil {
		res.Code = http.StatusInternalServerError
		return
	}

	res.Code = http.StatusOK
}

func (res *Response) setData() {
	if res.Data != nil {
		return
	}

	var responder DataResponder
	if errors.As(res.Err, &responder) {
		res.Data = responder.Data()
		return
	}
}

func (res *Response) setMessage(showError bool) {
	if res.Status == StatusSuccess {
		res.Message = ""
		return
	}

	if res.Message != "" {
		return
	}

	var responder MsgResponder
	if errors.As(res.Err, &responder) {
		res.Message = responder.Message()
		return
	}

	if res.Status == StatusError && res.Err != nil {
		res.Message = res.Err.Error()
		return
	}

	if res.Status == StatusFail && res.Err != nil && showError {
		res.Message = res.Err.Error()
		return
	}

	res.Message = http.StatusText(res.Code)
}

// Parse will prepare Response for rendering.
func (res *Response) Parse(ctx context.Context) {
	res.setCode()

	switch {
	case res.Code >= 500 && res.Code <= 599:
		res.Status = StatusFail
	case res.Code >= 400 && res.Code <= 499:
		res.Status = StatusError
	default:
		res.Status = StatusSuccess
	}

	res.setMessage(ShowErrorFromCtx(ctx))
	res.setData()
}

// Render implements chi's render.Renderer interface.
func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	res.Parse(r.Context())
	return nil
}

// Reset will clear and reset the struct for reuse (e.g. object pool).
func (res *Response) Reset() {
	res.Code = 0
	res.Status = ""
	res.Message = ""
	res.Data = nil
	res.Err = nil
}
