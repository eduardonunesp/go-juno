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

func CreateCharge(chargeParams CreateChargeParams, oauthToken, resourceToken string) (*CreateChargeResponse, error) {
	var response CreateChargeResponse
	response.Status = 200

	bs, err := json.Marshal(chargeParams)

	if err != nil {
		return nil, err
	}

	url := ResourceServer + "/charges"
	op := newOperation(bs, url, methodPOST, createOperationHeaders(oauthToken, resourceToken))

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
