package transactions

import (
	"finance_manager/src/transactions/domain"
	"finance_manager/src/transactions/peristence"
	"finance_manager/src/transactions/rest"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type App struct {
	repo       *peristence.TransactionPostgresRepository
	restClient *rest.Client
}

func NewRestApp(pool *pgxpool.Pool) *App {
	repo := peristence.CreateNewTransactionRepository(pool)
	return &App{
		repo:       repo,
		restClient: rest.CreateClient(domain.CreateNewTransactionService(repo)),
	}
}

func (a *App) Init() error {
	err := a.repo.Init()
	if err != nil {
		slog.Error("Could not init tables", "cause", err.Error())
		return err
	}
	return nil
}

func (a *App) AddRoutes(group *gin.RouterGroup) {
	a.restClient.RegisterRoutes(group)
}
