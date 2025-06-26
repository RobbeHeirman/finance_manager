package main

import (
	"finance_manager/src/auth/domain"
	"finance_manager/src/auth/persistence"
	"finance_manager/src/auth/rest"
	"finance_manager/src/core"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"time"
)

func CreateRestEndpoint() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	err := r.SetTrustedProxies(nil)
	if err != nil {
		slog.Error("Fatal", err)
		panic(err)
	}
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	return r
}

func AddAuthEndpoint(r *gin.Engine) {
	client := rest.CreateRestClient(domain.NewAuthServiceImpl())
	client.RegisterRoutes(r.Group("/auth"))
}

func InstallApps() {
	env, err := core.CreateConnectionPoolFromEnv()
	if err != nil {
		log.Fatalf("Failed to create connection pool. Error: %s", err)
	}

	repo := persistence.CreateUserRepo(env)
	if err = repo.Init(); err != nil {
		log.Fatalf("Failed to initialize repo. Error: %s", err)
	}
}

func main() {
	//InstallApps()
	r := CreateRestEndpoint()
	AddAuthEndpoint(r)
	_ = r.Run()
}
