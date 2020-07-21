package gojuno

import (
	"os"
	"testing"
)

func TestCreateWebhook(t *testing.T) {
	ClientID := os.Getenv("JUNO_CLIENT_ID")
	ClientSecret := os.Getenv("JUNO_CLIENT_SECRET")
	ResourceToken := os.Getenv("JUNO_RESOURCE_TOKEN")
	WebhookTest := os.Getenv("JUNO_WEBHOOK_TEST")

	result, err := NewOauthToken(ClientID, ClientSecret)

	if err != nil {
		t.Errorf("Failed get authorization token %s\n", err)
		return
	}

	response, err := CreateWebhook(CreateWebhookParams{
		URL: WebhookTest,
		EventTypes: []string{
			WebhookTypeChargeStatusChanged,
		},
	}, result.AccessToken, ResourceToken)

	if err != nil {
		t.Errorf("Failed to create charge cause %+v\n %+v\n", err, response)
		return
	}
}
