package gojuno

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	methodPOST  = "POST"
	timeoutSecs = 10
)

var (
	ClientID       string
	ClientSecret   string
	ResourceToken  string
	AuthServer     = "https://sandbox.boletobancario.com/authorization-server"
	ResourceServer = "https://sandbox.boletobancario.com/api-integration"
)

type ErrorResponse struct {
	Status  int    `json:"status,omitempty"`
	Error   string `json:"error,omitempty"`
	Details []struct {
		Message   string `json:"message,omitempty"`
		ErrorCode string `json:"errorCode,omitempty"`
	} `json:"details,omitempty"`
}

type operation struct {
	headers map[string]string
	method  string
	path    string
	body    []byte
}

func newOperation(body []byte, path, method string, headers map[string]string) operation {
	return operation{body: body, path: path, method: method, headers: headers}
}

func newOperationWith(body []byte) operation {
	return operation{body: body}
}

func resourceHeaders(authToken string) map[string]string {
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + authToken
	headers["X-Api-Version"] = "2"
	headers["X-Resource-Token"] = ResourceToken
	headers["Content-Type"] = "application/json"
	return headers
}

func dispatch(operation operation) ([]byte, error) {
	timeout := time.Duration(timeoutSecs * time.Second)

	httpClient := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest(operation.method, operation.path, bytes.NewBuffer(operation.body))

	for k, v := range operation.headers {
		request.Header.Set(k, v)
	}

	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
