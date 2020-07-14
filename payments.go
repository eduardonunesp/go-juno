package gojuno

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street   string `json:"street"`
	Number   string `json:"number"`
	City     string `json:"city"`
	State    string `json:"state"`
	PostCode string `json:"postCode"`
}

type PaymentBilling struct {
	Email   string `json:"email"`
	Address `json:"address"`
}

type CreditCardDetails struct {
	CreditCardID   string `json:"creditCardId,omitempty"`
	CreditCardHash string `json:"creditCardHash,omitempty"`
}

type PaymentParams struct {
	ChargeID          string `json:"chargeId"`
	PaymentBilling    `json:"billing"`
	CreditCardDetails `json:"creditCardDetails"`
}

type CreatePaymentResponse struct {
	TransactionID string `json:"transactionId"`
	Installments  int    `json:"installments"`
	Payments      struct {
		ID          string  `json:"id"`
		ChargeID    string  `json:"chargeId"`
		Date        string  `json:"date"`
		ReleaseDate string  `json:"releaseDate"`
		Amount      float32 `json:"amount"`
		Fee         float32 `json:"fee"`
		Type        string  `json:"type"`
		Status      string  `json:"status"`
		FailReason  string  `json:"failReason"`
	} `json:"payments"`

	ErrorResponse
}

func CreatePayment(paymentParams PaymentParams, oauthToken string) (*CreatePaymentResponse, error) {
	bs, err := json.Marshal(paymentParams)

	if err != nil {
		return nil, err
	}

	op := newOperation(bs, ResourceServer+"/payments", methodPOST, resourceHeaders(oauthToken))

	body, err := dispatch(op)

	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	var response CreatePaymentResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
