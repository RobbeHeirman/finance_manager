package auth

import (
	"finance_manager/src/core"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func AddRouteGroup(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.POST("/google_auth", core.RestPostWrapper(googleAuthHandler))

}

type TokenRequest struct {
	Token string `json:"idToken"`
}

type UserResponse struct {
	Message string `json:"message"`
}

func googleAuthHandler(request *TokenRequest) (*UserResponse, *core.HttpError) {
	_, err := ValidateGoogleUserToken(request.Token)
	if err != nil {
		res := core.HttpError{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		}
		slog.Error("Error validating google user token", "stacktrace", err.Error())
		return nil, &res
	}
	return &UserResponse{Message: "Login succes"}, nil
}
