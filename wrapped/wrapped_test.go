package wrapped_test

import (
	"net/http"
)

type MockError struct {
}

// StatusCode implements rest.StatusCodeResponder.
func (m *MockError) StatusCode() int {
	return http.StatusBadGateway
}

// Message implements rest.MsgResponder.
func (m *MockError) Message() string {
	return "errorMsg"
}

// Data implements rest.DataResponder.
func (m *MockError) Data() interface{} {
	return []string{"msg1", "msg2"}
}

func (m *MockError) Error() string {
	return "errorMsg"
}
