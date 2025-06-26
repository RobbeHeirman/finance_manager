package rest

import (
	"context"
	"finance_manager/src/auth/domain"
	"finance_manager/src/core/data_structures"
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

// ValidateGoogleUserToken Takes in a Google JWT token and returns a GoogleClaims object.
// A structured object of the JWT claims.
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

func GoogleClaimsToUserAdapter(claims *GoogleClaims) (*domain.User, error) {
	email, err := data_structures.NewEmail(claims.Email)
	if err != nil {
		return nil, err
	}

	url, err := data_structures.CreateUrl(claims.Picture)
	if err != nil {
		return nil, err
	}
	user, err := domain.NewUser(
		nil,
		&email,
		&claims.GivenName,
		&claims.FamilyName,
		&url,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
