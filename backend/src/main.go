package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

func main() {
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

	//auth.CreateRestAdapter().RegisterRoutes(r.Group("/auth"))
	_ = r.Run()
}
