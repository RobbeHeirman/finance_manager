package rest

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type Client struct {
	// domain domain.AuthService
}

func CreateClient() *Client {
	return &Client{}
}

func (client *Client) receiveKbcTransactionsCsv(g *gin.Context) {
	file, err := g.FormFile("file")
	if err != nil {
		slog.Error(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{"failed to get file %s": err.Error()})
	}

	// Refactor to seperate
	fileHeader, err := file.Open()
	defer func() {
		localErr := fileHeader.Close()
		if localErr != nil {
			err = localErr
			slog.Error("Problems closing file", "reason", err.Error())
		}
	}()

}
