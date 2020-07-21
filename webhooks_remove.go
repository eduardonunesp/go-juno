package gojuno

import (
	"encoding/json"
	"fmt"
)

type RemoveWebhookParams struct {
	ID string `json:"id"`
}

type RemoveWebhookResponse struct {
	StatusResponse
}

func RemoveWebhook(removeWebhookParams RemoveWebhookParams, oauthToken, resourceToken string) (*RemoveWebhookResponse, error) {
	url := fmt.Sprintf("%s/notifications/webhooks/%s", ResourceServer, removeWebhookParams.ID)
	op := newOperation([]byte(""), url, methodDELETE, createOperationHeaders(oauthToken, resourceToken))

	body, err := request(op)

	if err != nil {
		return nil, err
	}

	var response RemoveWebhookResponse
	response.Status = 204

	if len(body) != 0 {
		err = json.Unmarshal(body, &response)

		if err != nil {
			return nil, err
		}
	}

	if response.Status != 204 {
		return &response, fmt.Errorf("%s", response.Error)
	}

	return &response, nil
}
