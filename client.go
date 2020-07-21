package gojuno

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	methodPOST   = "POST"
	methodPUT    = "PUT"
	methodGET    = "GET"
	methodDELETE = "DELETE"
	timeoutSecs  = 60
)

var (
	AuthServer     = "https://sandbox.boletobancario.com/authorization-server"
	ResourceServer = "https://sandbox.boletobancario.com/api-integration"
)

type StatusResponse struct {
	Status  interface{} `json:"status,omitempty"`
	Error   string      `json:"error,omitempty"`
	Details []struct {
		Message   string `json:"message,omitempty"`
		ErrorCode string `json:"errorCode,omitempty"`
	} `json:"details,omitempty"`
}

type operationParams struct {
	headers map[string]string
	method  string
	path    string
	body    []byte
}

func newOperation(body []byte, path, method string, headers map[string]string) operationParams {
	return operationParams{body: body, path: path, method: method, headers: headers}
}

func createOperationHeaders(authToken, resourceToken string) map[string]string {
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + authToken
	headers["X-Api-Version"] = "2"
	headers["X-Resource-Token"] = resourceToken
	headers["Content-Type"] = "application/json"
	return headers
}

func request(operation operationParams) ([]byte, int, error) {
	timeout := time.Duration(timeoutSecs * time.Second)

	httpClient := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest(operation.method, operation.path, bytes.NewBuffer(operation.body))

	for k, v := range operation.headers {
		request.Header.Set(k, v)
	}

	if err != nil {
		return nil, 500, err
	}

	response, err := httpClient.Do(request)

	if err != nil {
		return nil, response.StatusCode, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, response.StatusCode, err
	}

	return body, response.StatusCode, nil
}
