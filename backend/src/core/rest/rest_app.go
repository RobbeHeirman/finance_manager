package rest

import "github.com/gin-gonic/gin"

type App interface {
	Init() error
	AddRoutes(group *gin.RouterGroup)
}
