package gojuno

import (
	"os"
	"testing"
)

func TestCreateCharge(t *testing.T) {
	ClientID = os.Getenv("JUNO_CLIENT_ID")
	ClientSecret = os.Getenv("JUNO_CLIENT_SECRET")
	ResourceToken = os.Getenv("JUNO_RESOURCE_TOKEN")

	result, err := newOauthToken(ClientID, ClientSecret)

	if err != nil {
		t.Errorf("Failed get authorization token")
	}

	response, err := CreateCharge(ChargeParams{
		Charge: Charge{
			Description: "OK",
			Amount:      20.0,
		},
		Billing: Billing{
			Name:     "Foo Bar",
			Document: "96616796060",
		},
	}, result.AccessToken)

	if err != nil {
		t.Errorf("Failed to crete charge cause %v\n", err)
	}

	if len(response.Embedded.Charges) == 0 {
		t.Errorf("No charges returned")
	}
}
