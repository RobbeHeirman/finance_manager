package main

import (
	"finance_manager/src/auth"
	"finance_manager/src/core/config"
	"finance_manager/src/core/persistence"
	"finance_manager/src/core/rest"
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

// @title           Finance manager
// @version 0.0.1
func main() {
	envConfig := config.NewEnvironmentRepository()
	pool, err := persistence.CreateConnectionPool(envConfig)
	apps := []rest.App{auth.NewRestApp(envConfig, pool)}
	r := CreateRestEndpoint()
	for _, app := range apps {
		err := app.Init()
		if err != nil {
			panic(err)
		}
		app.AddRoutes(r)
	}

	if err != nil {
		log.Fatalf("Failed to create connection pool. Error: %s", err)
	}
	_ = r.Run()
}
