package router

import (
	"go-web/internal/api"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	engine := gin.Default()

	engine.GET("/", api.Index)
	engine.GET("/health", api.Health)

	return engine
}
