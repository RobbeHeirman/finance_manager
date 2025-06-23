package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type RestFunction[U any, V any] func(*U) (*V, *HttpError)

type HttpError struct {
	Code    int
	Message string
}

func (err *HttpError) Error() string {
	return fmt.Sprintf("code %d: %s", err.Code, err.Message)
}

func RestPostWrapper[U any, V any](function RestFunction[U, V]) gin.HandlerFunc {
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
