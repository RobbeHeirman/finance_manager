package transactions

import (
	"finance_manager/src/transactions/rest"
	"github.com/gin-gonic/gin"
)

type App struct {
	restClient *rest.Client
}

func NewRestApp() *App {
	return &App{
		restClient: rest.CreateClient(),
	}
}

func (a App) Init() error {
	return nil
}

func (a App) AddRoutes(group *gin.RouterGroup) {
	a.restClient.RegisterRoutes(group)
}
