package gojuno

import (
	"os"
	"testing"
)

func TestRemoveWebhook(t *testing.T) {
	ClientID := os.Getenv("JUNO_CLIENT_ID")
	ClientSecret := os.Getenv("JUNO_CLIENT_SECRET")
	ResourceToken := os.Getenv("JUNO_RESOURCE_TOKEN")
	WebhookTest := os.Getenv("JUNO_WEBHOOK_ID")

	result, err := NewOauthToken(ClientID, ClientSecret)

	if err != nil {
		t.Errorf("Failed get authorization token %s\n", err)
		return
	}

	_, err = RemoveWebhook(RemoveWebhookParams{
		ID: WebhookTest,
	}, result.AccessToken, ResourceToken)

	if err != nil {
		t.Errorf("Failed to remove webhook cause %+v\n", err)
		return
	}
}
