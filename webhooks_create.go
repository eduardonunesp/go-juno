package gojuno

import (
	"encoding/json"
	"fmt"
)

const (
	WebhookTypeDocumentStatusChanged      = "DOCUMENT_STATUS_CHANGED"
	WebhookTypeDigitalAccountStatusChaged = "DIGITAL_ACCOUNT_STATUS_CHANGED"
	WebhookTypeTransferStatusChanged      = "TRANSFER_STATUS_CHANGED"
	WebhookTypeP2PTransferStatusChanged   = "P2P_TRANSFER_STATUS_CHANGED"
	WebhookTypePaymentNotification        = "PAYMENT_NOTIFICATION"
	WebhookTypeChargeStatusChanged        = "CHARGE_STATUS_CHANGED"
)

type CreateWebhookParams struct {
	URL        string   `json:"url"`
	EventTypes []string `json:"eventTypes"`
}

type EventTypesResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Label  string `json:"label"`
	Status string `json:"status"`
}

type CreateWebhookResponse struct {
	ID         string               `json:"id"`
	URL        string               `json:"url"`
	Secret     string               `json:"secret"`
	Status     string               `json:"status"`
	EventTypes []EventTypesResponse `json:"eventTypes"`

	StatusResponse
}

func CreateWebhook(createWebookParams CreateWebhookParams, oauthToken, resourceToken string) (*CreateWebhookResponse, error) {
	var response CreateWebhookResponse
	response.StatusResponse.Status = 200

	bs, err := json.Marshal(createWebookParams)

	if err != nil {
		return nil, err
	}

	url := ResourceServer + "/notifications/webhooks"
	op := newOperation(bs, url, methodPOST, createOperationHeaders(oauthToken, resourceToken))

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

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
