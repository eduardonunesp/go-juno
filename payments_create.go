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
	var response CreatePaymentResponse
	response.Status = 200

	bs, err := json.Marshal(paymentParams)

	if err != nil {
		return nil, err
	}

	op := newOperation(bs, ResourceServer+"/payments", methodPOST, createOperationHeaders(oauthToken, ResourceToken))

	body, status, err := request(op)

	if err != nil {
		return nil, err
	}

	if status != response.StatusResponse.Status {
		if err := json.Unmarshal(body, &response.StatusResponse); err != nil {
			return nil, err
		}

		return &response, fmt.Errorf("%s", response.Error)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
