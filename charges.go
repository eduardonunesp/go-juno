package gojuno

import (
	"encoding/json"
	"fmt"
)

const (
	PaymentTypeCreditCard = "CREDIT_CARD"
)

type ChargeParams struct {
	Description string   `json:"description"`
	Amount      float64  `json:"amount"`
	PaymentType []string `json:"paymentTypes"`
}

type ChargeBillingParams struct {
	Name     string `json:"name"`
	Document string `json:"document"`
}

type CreateChargeParams struct {
	ChargeParams        `json:"charge"`
	ChargeBillingParams `json:"billing"`
}

type CreateChargeResponse struct {
	Embedded struct {
		Charges []struct {
			ID   string `json:"id"`
			Code int64  `json:"code"`
		} `json:"charges"`
	} `json:"_embedded"`

	StatusResponse
}

type CancelChargeParams struct {
	ID string `json:"id"`
}

type CancelChargeResponse struct {
	StatusResponse
}

func CreateCharge(chargeParams CreateChargeParams, oauthToken, resourceToken string) (*CreateChargeResponse, error) {
	bs, err := json.Marshal(chargeParams)

	if err != nil {
		return nil, err
	}

	url := ResourceServer + "/charges"
	op := newOperation(bs, url, methodPOST, createOperationHeaders(oauthToken, resourceToken))

	body, err := request(op)

	if err != nil {
		return nil, err
	}

	var response CreateChargeResponse
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

func CancelCharge(cancelChargeParams CancelChargeParams, oauthToken, resourceToken string) (*CancelChargeResponse, error) {
	url := fmt.Sprintf("%s/charges/%s/cancelation", ResourceServer, cancelChargeParams.ID)
	op := newOperation([]byte(""), url, methodPUT, createOperationHeaders(oauthToken, resourceToken))

	body, err := request(op)

	if err != nil {
		return nil, err
	}

	var response CancelChargeResponse
	response.Status = 204

	if len(body) != 0 {
		err = json.Unmarshal(body, &response)

		if err != nil {
			return nil, err
		}
	}

	if response.Status != 204 {
		return &response, fmt.Errorf("%s", response.Error)
	}

	return &response, nil
}
