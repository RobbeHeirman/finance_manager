package rest

import (
	"crypto"
	"finance_manager/src/core/security"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type Function[U any, V any] func(*U) (*V, *HttpError)

type HttpError struct {
	Code    int
	Message string
}

func (err *HttpError) Error() string {
	return fmt.Sprintf("code %d: %s", err.Code, err.Message)
}

func PostWrapper[U any, V any](function Function[U, V]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req U
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			slog.Error("Could not Bind Json", err)
			return
		}
		result, err := function(&req)
		if err != nil {
			c.JSON(err.Code, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

const UserIdKey = "user_id"

func JWTMiddleware(publicKey crypto.PublicKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		claims, err := security.DecodeAndValidateJWT(publicKey, token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}
		c.Set(UserIdKey, claims.Id)
		c.Next()
	}
}
