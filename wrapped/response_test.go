package wrapped_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lanz-dev/go-rest/wrapped"
)

func TestWrapped_Parse(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		code    int
		status  string
		message string
		err     error
		data    interface{}
	}{
		"InternalServerError": {
			code:    http.StatusInternalServerError,
			status:  wrapped.StatusFail,
			message: http.StatusText(http.StatusInternalServerError),
			err:     nil,
			data:    nil,
		},
		"BadRequest": {
			code:    http.StatusBadRequest,
			status:  wrapped.StatusError,
			message: http.StatusText(http.StatusBadRequest),
			err:     nil,
			data:    nil,
		},
		"Success": {
			code:    http.StatusOK,
			status:  wrapped.StatusSuccess,
			message: "",
			err:     nil,
			data:    nil,
		},
		"Error": {
			code:    http.StatusOK,
			status:  wrapped.StatusSuccess,
			message: "",
			err:     errors.New("unittest"),
			data:    nil,
		},
	}

	for name, tc := range tests {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest("GET", "/unittest", nil)

			res := wrapped.Response{Code: tc.code, Err: tc.err}
			res.Parse(req.Context())

			if status := res.Status; status != tc.status {
				t.Errorf("expected status '%s' but got '%s`", tc.status, status)
			}

			if message := res.Message; message != tc.message {
				t.Errorf("expected message '%s' but got '%s`", tc.message, message)
			}

			if data := res.Data; data != tc.data {
				t.Errorf("expected data '%+v' but got '%+v`", tc.data, data)
			}
		})
	}
}

// If status is "success" Message will always be "" and not present in json.
func TestWrapped_Parse_SetMessage_Success(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)

	res := wrapped.Response{Code: http.StatusOK, Message: "unittest"}
	res.Parse(req.Context())

	if res.Message != "" {
		t.Fatal(`expected Message to be empty`)
	}
}

// If Message is already set, the already set value will be used.
func TestWrapped_Parse_SetMessage_ValueWillBeUsed(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)

	expected := "unittest"
	res := wrapped.Response{Code: http.StatusForbidden, Message: expected}
	res.Parse(req.Context())

	if res.Message != expected {
		t.Fatalf(`expected Message to be '%s', got: '%s'`, expected, res.Message)
	}
}

// If status is not "success" and Err will implement MsgResponder, Message will be MsgResponder.ResponseMessage().
func TestWrapped_Parse_SetMessage_MsgResponder(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)

	mockErr := &MockError{}
	res := wrapped.Response{Code: http.StatusForbidden, Err: mockErr}
	res.Parse(req.Context())

	if res.Message != mockErr.Message() {
		t.Fatalf(`expected Message to be '%s', got: '%s'`, mockErr.Message(), res.Message)
	}
}

// If Message is set it will still win.
func TestWrapped_Parse_SetMessage_MessageIsAlreadySet(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)

	expected := "unittest"
	mockErr := &MockError{}
	res := wrapped.Response{Code: http.StatusForbidden, Err: mockErr, Message: expected}
	res.Parse(req.Context())

	if res.Message != expected {
		t.Fatalf(`expected Message to be '%s', got: '%s'`, expected, res.Message)
	}
}

// If status is "error" this will be Err.Error().
func TestWrapped_Parse_SetMessage_StatusError(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)

	mockErr := errors.New("unittest")
	res := wrapped.Response{Code: http.StatusBadRequest, Err: mockErr}
	res.Parse(req.Context())

	if res.Message != mockErr.Error() {
		t.Fatalf(`expected Message to be '%s', got: '%s'`, mockErr.Error(), res.Message)
	}
}

// If status is "fail" and ctxKeyShowError is true, Message will be Err.Error().
func TestWrapped_Parse_SetMessage_CtxShowError(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)

	ctx := req.Context()
	ctx = wrapped.CtxSetShowError(ctx, true)
	req = req.WithContext(ctx)

	mockErr := errors.New("unittest")
	res := wrapped.Response{Code: http.StatusInternalServerError, Err: mockErr}
	res.Parse(req.Context())

	if res.Message != mockErr.Error() {
		t.Fatalf(`expected Message to be '%s', got: '%s'`, mockErr.Error(), res.Message)
	}
}

// Default Message will contain http.StatusText(Code).
func TestWrapped_Parse_SetMessage_StatusFail(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)

	res := wrapped.Response{Code: http.StatusInternalServerError, Err: errors.New("unittest")}
	res.Parse(req.Context())

	if res.Message != http.StatusText(http.StatusInternalServerError) {
		t.Fatalf(`expected Message to be '%s', got: '%s'`, http.StatusText(http.StatusInternalServerError), res.Message)
	}
}

func TestWrapped_Parse_SetData(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)

	mockErr := &MockError{}
	res := wrapped.Response{Code: http.StatusInternalServerError, Err: mockErr}
	res.Parse(req.Context())

	if res.Message != mockErr.Message() {
		t.Fatalf(`expected Message to be '%s', got: '%s'`, mockErr.Message(), res.Message)
	}

	data, ok := res.Data.([]string)
	if !ok {
		t.Fatal(`expected Data to be []string`)
	}
	if len(data) != 2 && data[0] != "msg1" && data[1] != "msg2" {
		t.Fatal(`expected different Data`)
	}
}

func TestWrapped_Parse_DontSetDataIfAlreadySet(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)

	expected := "alreadySet"
	mockErr := &MockError{}
	res := wrapped.Response{Code: http.StatusInternalServerError, Data: expected, Err: mockErr}
	res.Parse(req.Context())

	if res.Message != mockErr.Message() {
		t.Fatalf(`expected Message to be '%s', got: '%s'`, mockErr.Message(), res.Message)
	}

	data, ok := res.Data.(string)
	if !ok {
		t.Fatal(`expected Data to be string`)
	}
	if data != expected {
		t.Fatalf(`expected Data to be '%s', got: '%s'`, expected, data)
	}
}

func TestWrapped_Parse_StatusDefaultsTo200(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)

	res := wrapped.Response{}
	res.Parse(req.Context())

	if res.Code != http.StatusOK {
		t.Fatalf(`expected Status to be '%d', got: '%d'`, http.StatusOK, res.Code)
	}
}

func TestWrapped_Parse_UnsetStatusWithErrorIs500(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)

	res := wrapped.Response{Err: errors.New("unittest")}
	res.Parse(req.Context())

	if res.Code != http.StatusInternalServerError {
		t.Fatalf(`expected Status to be '%d', got: '%d'`, http.StatusInternalServerError, res.Code)
	}
}

func TestWrapped_Parse_SetStatusWithResponder(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)

	mockErr := &MockError{}
	res := wrapped.Response{Err: mockErr}
	res.Parse(req.Context())

	if res.Code != http.StatusBadGateway {
		t.Fatalf(`expected Status to be '%d', got: '%d'`, http.StatusBadGateway, res.Code)
	}
}

func TestWrapped_Reset(t *testing.T) {
	t.Parallel()

	res := wrapped.Response{
		Code:    http.StatusInternalServerError,
		Status:  wrapped.StatusError,
		Message: "msg",
		Data:    "data",
		Err:     errors.New("unittest"),
	}
	res.Reset()

	if res.Code != 0 {
		t.Fatalf(`expected Code to be '%d', got: '%d'`, 0, res.Code)
	}
	if res.Status != "" {
		t.Fatalf(`expected Status to be '%s', got: '%s'`, "", res.Status)
	}
	if res.Message != "" {
		t.Fatalf(`expected Message to be '%s', got: '%s'`, "", res.Message)
	}
	if res.Data != nil {
		t.Fatalf(`expected Data to be 'nil', got: '%s'`, res.Data)
	}
	if res.Err != nil {
		t.Fatalf(`expected Err to be 'nil', got: '%s'`, res.Err)
	}
}

func TestWrapped_Render_WillCallParse(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/unittest", nil)
	w := httptest.NewRecorder()

	res := wrapped.Response{Code: http.StatusOK, Message: "unittest"}
	err := res.Render(w, req)
	if err != nil {
		t.Fatalf("did not expected error '%s'", err)
	}

	if res.Message != "" {
		t.Fatal(`expected Message to be empty`)
	}
}
