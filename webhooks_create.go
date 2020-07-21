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
	EventTypes []EventTypesResponse `json:"eventTypes"`

	StatusResponse
}

func CreateWebhook(createWebookParams CreateWebhookParams, oauthToken, resourceToken string) (*CreateWebhookResponse, error) {
	bs, err := json.Marshal(createWebookParams)

	if err != nil {
		return nil, err
	}

	url := ResourceServer + "/notifications/webhooks"
	op := newOperation(bs, url, methodPOST, createOperationHeaders(oauthToken, resourceToken))

	body, err := request(op)

	if err != nil {
		return nil, err
	}

	var response CreateWebhookResponse
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
