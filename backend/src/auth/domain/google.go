package domain

import (
	"context"
	"github.com/goccy/go-json"
	"google.golang.org/api/idtoken"
	"log/slog"
)

const ClientId = "63575078815-kl24b59mf9adslcut5671amaqm05een3.apps.googleusercontent.com"

type GoogleClaims struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	Picture    string `json:"picture"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Locale     string `json:"locale"`
}

func ValidateGoogleUserToken(token string) (*GoogleClaims, error) {

	response, err := idtoken.Validate(context.Background(), token, ClientId)
	if err != nil {
		slog.Error("Couldn't validate Google user token:", err)
		return nil, err
	}

	data, err := json.Marshal(response.Claims)
	if err != nil {
		slog.Error("Failed to marshal", "trace", err.Error())
		return nil, err
	}

	var claims GoogleClaims
	if err := json.Unmarshal(data, &claims); err != nil {
		slog.Error("Failed to write claims", "trace", err.Error())
		return nil, err
	}
	return &claims, nil
}
