package main

import (
	"finance_manager/src/auth"
	"finance_manager/src/core/config"
	"finance_manager/src/core/persistence"
	"finance_manager/src/core/rest"
	"finance_manager/src/transactions"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"time"
)

func CreateGinServer() *gin.Engine {
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

func registerApp(server *gin.Engine, app rest.App, route string, handlers ...gin.HandlerFunc) rest.App {
	group := server.Group(route, handlers...)
	app.AddRoutes(group)
	return app
}

// @title           Finance manager
// @version 0.0.1
func main() {
	envConfig := config.NewEnvironmentRepository()
	pool, err := persistence.CreateConnectionPool(envConfig)
	server := CreateGinServer()
	jwtMiddleware := rest.JWTMiddleware(envConfig.GetPublicKey())
	apps := []rest.App{
		registerApp(server, auth.NewRestApp(envConfig, pool), "/auth"),
		registerApp(server, transactions.NewRestApp(pool), "/transaction", jwtMiddleware),
	}
	for _, app := range apps {
		err := app.Init()
		if err != nil {
			panic(err)
		}
	}

	if err != nil {
		log.Fatalf("Failed to create connection pool. Error: %s", err)
	}
	_ = server.Run()
}
