package gojuno

import (
	"encoding/json"
	"fmt"
)

type AddressParams struct {
	Street   string `json:"street"`
	Number   string `json:"number"`
	City     string `json:"city"`
	State    string `json:"state"`
	PostCode string `json:"postCode"`
}

type PaymentBillingParams struct {
	Email         string `json:"email"`
	AddressParams `json:"address"`
}

type CreditCardDetailsParams struct {
	CreditCardID   string `json:"creditCardId,omitempty"`
	CreditCardHash string `json:"creditCardHash,omitempty"`
}

type CreatePaymentParams struct {
	ChargeID                string `json:"chargeId"`
	PaymentBillingParams    `json:"billing"`
	CreditCardDetailsParams `json:"creditCardDetails"`
}

type CreatePaymentResponse struct {
	TransactionID string `json:"transactionId"`
	Installments  int    `json:"installments"`
	Payments      []struct {
		ID          string  `json:"id"`
		ChargeID    string  `json:"chargeId"`
		Date        string  `json:"date"`
		ReleaseDate string  `json:"releaseDate"`
		Amount      float64 `json:"amount"`
		Fee         float64 `json:"fee"`
		Type        string  `json:"type"`
		Status      string  `json:"status"`
		FailReason  string  `json:"failReason"`
	} `json:"payments"`

	StatusResponse
}

func CreatePayment(paymentParams CreatePaymentParams, oauthToken, ResourceToken string) (*CreatePaymentResponse, error) {
	bs, err := json.Marshal(paymentParams)

	if err != nil {
		return nil, err
	}

	op := newOperation(bs, ResourceServer+"/payments", methodPOST, createOperationHeaders(oauthToken, ResourceToken))

	body, err := request(op)

	if err != nil {
		return nil, err
	}

	var response CreatePaymentResponse
	response.Status = 200
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	if response.Status != 200 {
		return &response, fmt.Errorf("%s", response.Error)
	}

	return &response, nil
}
