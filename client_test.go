package gojuno

import (
	"testing"
)

func TestBasicDispatch(t *testing.T) {
	url := "https://postman-echo.com/post"
	header := make(map[string]string)
	operation := newOperation([]byte(""), url, methodPOST, header)
	if _, _, err := request(operation); err != nil {
		t.Error("Failed to dispatch to test url")
	}
}
