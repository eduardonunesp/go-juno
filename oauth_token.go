package gojuno

import (
	b64 "encoding/base64"
	"encoding/json"
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
	basicToken := b64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))

	headers := make(map[string]string)
	headers["Authorization"] = "Basic " + basicToken
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	body, err := request(operationParams{
		headers: headers,
		path:    AuthServer + "/oauth/token",
		body:    []byte("grant_type=client_credentials"),
		method:  methodPOST,
	})

	if err != nil {
		return nil, err
	}

	var response OauthTokenResponse
	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
