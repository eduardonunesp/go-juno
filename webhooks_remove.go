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
	var response RemoveWebhookResponse
	response.Status = 204

	url := fmt.Sprintf("%s/notifications/webhooks/%s", ResourceServer, removeWebhookParams.ID)
	op := newOperation([]byte(""), url, methodDELETE, createOperationHeaders(oauthToken, resourceToken))

	body, status, err := request(op)

	if err != nil {
		return nil, err
	}

	if status != response.StatusResponse.Status {
		if err := json.Unmarshal(body, &response.StatusResponse); err != nil {
			return nil, err
		}

		return &response, fmt.Errorf("%s", response.Error)
	}

	return &response, nil
}
