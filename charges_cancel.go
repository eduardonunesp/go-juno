package gojuno

import (
	"encoding/json"
	"fmt"
)

type CancelChargeParams struct {
	ID string `json:"id"`
}

type CancelChargeResponse struct {
	StatusResponse
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
