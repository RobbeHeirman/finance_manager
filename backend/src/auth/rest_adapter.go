package auth

import (
	"finance_manager/src/auth/domain"
	"finance_manager/src/auth/persistence"
	"finance_manager/src/core"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type RestAdapter struct {
	repo persistence.Interface
}

func CreateRestAdapter(repo persistence.Interface) *RestAdapter {
	return &RestAdapter{repo: repo}
}

func (adapter *RestAdapter) RegisterRoutes(router *gin.RouterGroup) *RestAdapter {
	router.POST("/google_auth", core.RestPostWrapper(adapter.googleAuthHandler))
	return adapter
}

type TokenRequest struct {
	Token string `json:"idToken"`
}

type UserResponse struct {
	Message string `json:"message"`
}

func (adapter *RestAdapter) googleAuthHandler(request *TokenRequest) (*UserResponse, *core.HttpError) {
	_, err := domain.ValidateGoogleUserToken(request.Token)
	if err != nil {
		res := core.HttpError{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		}
		slog.Error("Error validating google user token", "stacktrace", err.Error())
		return nil, &res
	}
	return &UserResponse{Message: "Login success"}, nil
}
