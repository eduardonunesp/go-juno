package gojuno

import (
	"os"
	"testing"
)

func TestCreatePayment(t *testing.T) {
	ClientID := os.Getenv("JUNO_CLIENT_ID")
	ClientSecret := os.Getenv("JUNO_CLIENT_SECRET")
	ResourceToken := os.Getenv("JUNO_RESOURCE_TOKEN")
	ChargeID := os.Getenv("CHARGE_ID")
	CreditCardHash := os.Getenv("CREDIT_CARD_HASH")

	result, err := NewOauthToken(ClientID, ClientSecret)

	if err != nil {
		t.Errorf("Failed get authorization token")
	}

	response, err := CreatePayment(CreatePaymentParams{
		ChargeID: ChargeID,
		PaymentBillingParams: PaymentBillingParams{
			Email: "eduardonunesp@gmail.com",
			AddressParams: AddressParams{
				Street:   "Acacia Avenue",
				Number:   "22",
				City:     "Londom",
				State:    "SC",
				PostCode: "08226021",
			},
		},
		CreditCardDetailsParams: CreditCardDetailsParams{
			CreditCardHash: CreditCardHash,
		},
	}, result.AccessToken, ResourceToken)

	if err != nil {
		t.Errorf("Failed to crete Payment cause %+v\n", err)
	}

	if response.Status != 200 {
		t.Errorf("Failed to create payment cause %+v\n", response)
	}
}
