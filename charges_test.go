package gojuno

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateCharge(t *testing.T) {
	ClientID = os.Getenv("JUNO_CLIENT_ID")
	ClientSecret = os.Getenv("JUNO_CLIENT_SECRET")
	ResourceToken = os.Getenv("JUNO_RESOURCE_TOKEN")

	result, err := newOauthToken(ClientID, ClientSecret)

	if err != nil {
		t.Errorf("Failed get authorization token %s\n", err)
		return
	}

	response, err := CreateCharge(ChargeParams{
		Charge: Charge{
			Description: "OK",
			Amount:      20.0,
			PaymentType: []string{PaymentTypeCreditCard},
		},
		ChargeBilling: ChargeBilling{
			Name:     "Foo Bar",
			Document: "96616796060",
		},
	}, result.AccessToken)

	if err != nil {
		t.Errorf("Failed to crete charge cause %v\n", err)
		return
	}

	if len(response.Embedded.Charges) == 0 {
		t.Errorf("No charges returned")
		return
	}

	fmt.Println(response.Embedded.Charges[0].ID)
}
