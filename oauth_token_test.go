package gojuno

import (
	"os"
	"testing"
)

func TestAuthorizationToken(t *testing.T) {
	ClientID = os.Getenv("JUNO_CLIENT_ID")
	ClientSecret = os.Getenv("JUNO_CLIENT_SECRET")

	_, err := NewOauthToken(ClientID, ClientSecret)

	if err != nil {
		t.Error("Failed to request token")
	}
}

func TestAuthorizationFailed(t *testing.T) {
	ClientID = "foo"
	ClientSecret = "bar"

	oauthToken, err := NewOauthToken(ClientID, ClientSecret)

	if err != nil {
		t.Error("Failed to request token")
	}

	if oauthToken.AccessToken == "ZXhlbXBsby1jbGllbnQtaWQ6ZXhlbXBsby1jbGllbnQtc2VjcmV0" {
		t.Error("Authorization should fail")
	}
}
