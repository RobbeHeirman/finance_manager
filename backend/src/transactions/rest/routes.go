package rest

import "github.com/gin-gonic/gin"

type Client struct {
	// domain domain.AuthService
}

func CreateClient() *Client {
	return &Client{}
}

func (client *Client) receiveKbcTransactionsCsv(gin *gin.Context) {

}
