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
	var response CancelChargeResponse
	response.StatusResponse.Status = 204

	url := fmt.Sprintf("%s/charges/%s/cancelation", ResourceServer, cancelChargeParams.ID)
	op := newOperation([]byte(""), url, methodPUT, createOperationHeaders(oauthToken, resourceToken))

	body, status, err := request(op)

	if err != nil {
		return nil, err
	}

	if status != response.StatusResponse.Status {
		if err = json.Unmarshal(body, &response.StatusResponse); err != nil {
			return nil, err
		}

		return &response, fmt.Errorf("%s", response.Error)
	}

	return &response, nil
}
