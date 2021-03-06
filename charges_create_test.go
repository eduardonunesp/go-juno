package gojuno

import (
	"os"
	"testing"
)

func TestCreateCharge(t *testing.T) {
	ClientID := os.Getenv("JUNO_CLIENT_ID")
	ClientSecret := os.Getenv("JUNO_CLIENT_SECRET")
	ResourceToken := os.Getenv("JUNO_RESOURCE_TOKEN")

	result, err := NewOauthToken(ClientID, ClientSecret)

	if err != nil {
		t.Errorf("Failed get authorization token %s\n", err)
		return
	}

	response, err := CreateCharge(CreateChargeParams{
		ChargeParams: ChargeParams{
			Description: "OK",
			Amount:      20.0,
			PaymentType: []string{PaymentTypeCreditCard},
		},
		ChargeBillingParams: ChargeBillingParams{
			Name:     "Foo Bar",
			Document: "96616796060",
		},
	}, result.AccessToken, ResourceToken)

	if err != nil {
		t.Errorf("Failed to create charge cause %+v %+v\n", err, response)
		return
	}

	if len(response.Embedded.Charges) == 0 {
		t.Errorf("No charges returned")
		return
	}
}
