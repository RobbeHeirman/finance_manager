package _interface

import (
	"finance_manager/src/auth/domain"
	"finance_manager/src/core"
	"finance_manager/src/core/data_structures"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

type RestClient struct {
	domain domain.AuthService
}

func CreateRestClient(authDomain domain.AuthService) *RestClient {
	return &RestClient{domain: authDomain}
}

func (adapter *RestClient) RegisterRoutes(router *gin.RouterGroup) *RestClient {
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

func (adapter *RestClient) googleAuthHandler(request *TokenRequest) (*UserResponse, *core.HttpError) {
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

	user := adapter.domain.CreateUpdateUser(googleUser)
	id, _ := user.GetId().Get()
	userEmail := user.GetEmail()

	return &UserResponse{
		UserId:    id,
		UserEmail: userEmail,
	}, nil
}
