package rest

import (
	"encoding/csv"
	"finance_manager/src/transactions/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"net/http"
)

type Client struct {
	domain domain.TransactionService
}

func CreateClient(domain domain.TransactionService) *Client {
	return &Client{
		domain: domain,
	}
}

func (adapter *Client) RegisterRoutes(router *gin.RouterGroup) *Client {
	router.POST("/upload_kbc_csv", adapter.receiveKbcTransactionsCsv)
	return adapter
}

// UploadCSV godoc
// @Summary Upload a CSV file
// @Description Upload a CSV file via multipart/form-data
// @Tags transactions
// @Accept mpfd
// @Produce json
// @Param file formData file true "CSV file"
// @Success 200 {string} string "ok"
// @Failure 400 {string} string "bad request"
// @Id kbcTransactionsUpload
// @Router /transaction/upload_kbc_csv [post]
func (adapter *Client) receiveKbcTransactionsCsv(g *gin.Context) {
	file, err := g.FormFile("file")
	if err != nil {
		slog.Error(err.Error())
		g.JSON(http.StatusBadRequest, gin.H{"failed to get file %s": err.Error()})
	}

	fileHeader, err := file.Open()
	defer func() {
		localErr := fileHeader.Close()
		if localErr != nil {
			err = localErr
			slog.Error("Problems closing file", "reason", err.Error())
		}
	}()

	csvReader := csv.NewReader(fileHeader)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println("CSV Line:", record)
	}
	g.String(http.StatusOK, "CSV file processed successfully")
}
