package gojuno

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestListWebhook(t *testing.T) {
	ClientID := os.Getenv("JUNO_CLIENT_ID")
	ClientSecret := os.Getenv("JUNO_CLIENT_SECRET")
	ResourceToken := os.Getenv("JUNO_RESOURCE_TOKEN")

	result, err := NewOauthToken(ClientID, ClientSecret)

	if err != nil {
		t.Errorf("Failed get authorization token %s\n", err)
		return
	}

	response, err := ListWebhook(result.AccessToken, ResourceToken)

	if err != nil {
		t.Errorf("Failed to list webhooks cause %+v\n", err)
		return
	}

	prettyJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatal("Failed to generate json", err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}
