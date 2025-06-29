package auth

import (
	"finance_manager/src/auth/domain"
	"finance_manager/src/auth/persistence"
	"finance_manager/src/auth/rest"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRestApp(config domain.ConfigRepo, pool *pgxpool.Pool) *App {
	userRepo := persistence.NewUserRepo(pool)
	logic := domain.NewAuthServiceImpl(config, &userRepo)
	client := rest.CreateRestClient(logic)
	return &App{
		persist: userRepo,
		logic:   logic,
		client:  client,
	}
}

type App struct {
	persist persistence.UserRepo
	logic   domain.AuthService
	client  *rest.Client
}

func (a *App) Init() error {
	return a.persist.Init()
}

func (a *App) AddRoutes(engine *gin.Engine) {
	group := engine.Group("/auth")
	a.client.RegisterRoutes(group)
}
