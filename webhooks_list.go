package gojuno

import (
	"encoding/json"
	"fmt"
)

type ListWebhookResponse struct {
	Embedded struct {
		Webhooks []CreateWebhookResponse `json:"webhooks"`
	} `json:"_embedded"`

	StatusResponse
}

func ListWebhook(oauthToken, resourceToken string) (*ListWebhookResponse, error) {
	url := ResourceServer + "/notifications/webhooks"
	op := newOperation([]byte(""), url, methodGET, createOperationHeaders(oauthToken, resourceToken))

	body, err := request(op)

	if err != nil {
		return nil, err
	}

	var response ListWebhookResponse
	response.StatusResponse.Status = 200
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	if response.StatusResponse.Status != 200 {
		return &response, fmt.Errorf("%s", response.Error)
	}

	return &response, nil
}
