package gojuno

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
)

type OauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
	UserName    string `json:"user_name"`
	JTI         string `json:"jti"`

	StatusResponse
}

func NewOauthToken(clientID, clientSecret string) (*OauthTokenResponse, error) {
	var response OauthTokenResponse
	response.StatusResponse.Status = 200

	basicToken := b64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))

	headers := make(map[string]string)
	headers["Authorization"] = "Basic " + basicToken
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	body, status, err := request(operationParams{
		headers: headers,
		path:    AuthServer + "/oauth/token",
		body:    []byte("grant_type=client_credentials"),
		method:  methodPOST,
	})

	if err != nil {
		return nil, err
	}

	if status != response.StatusResponse.Status {
		if err := json.Unmarshal(body, &response.StatusResponse); err != nil {
			return nil, err
		}

		return &response, fmt.Errorf("%s", response.Error)
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
