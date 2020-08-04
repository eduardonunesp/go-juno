package gojuno

import (
	"encoding/json"
	"fmt"
)

type SplitParams struct {
	RecipientToken  string  `json:"recipientToken"`
	Amount          float64 `json:"amount"`
	Percentage      float64 `json:"percentage"`
	AmountRemainder bool    `json:"amountRemainder"`
	ChargeFee       bool    `json:"chargeFee"`
}

type CreateRefundParams struct {
	PaymentID   string
	Amount      float64       `json:"amount"`
	SplitParams []SplitParams `json:"split"`
}

type CreateRefundResponse struct {
	TransactionID string `json:"transactionId"`
	Installments  int    `json:"installments"`
	Refunds       []struct {
		ID            string  `json:"id"`
		ChargeID      string  `json:"chargeId"`
		ReleaseDate   string  `json:"releaseDate"`
		PaybackDate   string  `json:"paybackDate"`
		PaybackAmount float64 `json:"paybackAmount"`
		Status        string  `json:"status"`
	}

	StatusResponse
}

func CreateRefund(RefundParams CreateRefundParams, oauthToken, resourceToken string) (*CreateRefundResponse, error) {
	var response CreateRefundResponse
	response.StatusResponse.Status = 200

	url := fmt.Sprintf("%s/payments/%s/refunds", ResourceServer, RefundParams.PaymentID)
	op := newOperation([]byte(""), url, methodPOST, createOperationHeaders(oauthToken, resourceToken))

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
