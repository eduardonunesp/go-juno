package gojuno

import (
	"os"
	"testing"
)

func TestCreateRefund(t *testing.T) {
	ClientID := os.Getenv("JUNO_CLIENT_ID")
	ClientSecret := os.Getenv("JUNO_CLIENT_SECRET")
	ResourceToken := os.Getenv("JUNO_RESOURCE_TOKEN")
	PaymentID := os.Getenv("PAYMENT_ID")
	AuthServer = os.Getenv("JUNO_AUTH_SERVER")
	ResourceServer = os.Getenv("JUNO_RESOURCE_SERVER")

	result, err := NewOauthToken(ClientID, ClientSecret)

	if err != nil {
		t.Errorf("Failed get authorization token")
	}

	refunderResponse, err := CreateRefund(CreateRefundParams{
		PaymentID: PaymentID,
	}, result.AccessToken, ResourceToken)

	if err != nil {
		t.Errorf("Failed to create refund cause %+v %+v\n", err, result)
	}

	if refunderResponse.Status != 200 {
		t.Errorf("Failed to create refund cause %+v %+v\n", err, result)
	}
}
