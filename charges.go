package gojuno

import (
	"encoding/json"
)

const (
	PaymentTypeCreditCard = "CREDIT_CARD"
)

type Charge struct {
	Description string   `json:"description"`
	Amount      float32  `json:"amount"`
	PaymentType []string `json:"paymentTypes"`
}

type ChargeBilling struct {
	Name     string `json:"name"`
	Document string `json:"document"`
}

type ChargeParams struct {
	Charge        `json:"charge"`
	ChargeBilling `json:"billing"`
}

type CreateChargeResponse struct {
	Embedded struct {
		Charges []struct {
			ID   string `json:"id"`
			Code int64  `json:"code"`
		} `json:"charges"`
	} `json:"_embedded"`

	ErrorResponse
}

func CreateCharge(chargeParams ChargeParams, oauthToken string) (*CreateChargeResponse, error) {
	bs, err := json.Marshal(chargeParams)

	if err != nil {
		return nil, err
	}

	op := newOperation(bs, ResourceServer+"/charges", methodPOST, resourceHeaders(oauthToken))

	body, err := dispatch(op)

	if err != nil {
		return nil, err
	}

	var response CreateChargeResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
