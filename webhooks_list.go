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
	var response ListWebhookResponse
	response.StatusResponse.Status = 200

	url := ResourceServer + "/notifications/webhooks"
	op := newOperation([]byte(""), url, methodGET, createOperationHeaders(oauthToken, resourceToken))

	body, status, err := request(op)

	if err != nil {
		return nil, err
	}

	if status != response.StatusResponse.Status {
		if err = json.Unmarshal(body, &response.StatusResponse); err != nil {
			return nil, err
		}

		return &response, fmt.Errorf("%s", response.Error)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
