package rest

import (
	"finance_manager/src/auth/domain"
	"finance_manager/src/core"
	"finance_manager/src/core/data_structures"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

type Client struct {
	domain domain.AuthService
}

func CreateRestClient(authDomain domain.AuthService) *Client {
	return &Client{domain: authDomain}
}

func (adapter *Client) RegisterRoutes(router *gin.RouterGroup) *Client {
	router.POST("/google_auth", core.RestPostWrapper(adapter.googleAuthHandler))
	return adapter
}

type TokenRequest struct {
	Token string `json:"idToken"`
}

type UserResponse struct {
	UserId    *uuid.UUID             `json:"userId"`
	UserEmail *data_structures.Email `json:"userEmail"`
}

func (adapter *Client) googleAuthHandler(request *TokenRequest) (*UserResponse, *core.HttpError) {
	claims, err := ValidateGoogleUserToken(request.Token)
	if err != nil {
		res := core.HttpError{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		}
		slog.Error("Error validating google user token", "stacktrace", err.Error())
		return nil, &res
	}
	googleUser, err := GoogleClaimsToUserAdapter(claims)
	if err != nil {
		res := core.HttpError{
			Code:    http.StatusInternalServerError,
			Message: "Could not handle google claims",
		}
		return nil, &res
	}

	user, err := adapter.domain.CreateUpdateUser(googleUser)
	if err != nil {
		res := core.HttpError{
			Code:    http.StatusInternalServerError,
			Message: "Could not save to the database",
		}
		slog.Error("Error updating user", "stacktrace", err.Error())
		return nil, &res
	}
	id, _ := user.GetId().Get()
	userEmail := user.GetEmail()

	return &UserResponse{
		UserId:    id,
		UserEmail: userEmail,
	}, nil
}
