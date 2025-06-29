package rest

import "github.com/gin-gonic/gin"

type App interface {
	Init() error
	AddRoutes(*gin.Engine)
}
