package router

import (
	"go-web/internal/api"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	engine := gin.Default()

	engine.GET("/", api.Index)
	engine.GET("/health", api.Health)

	example := engine.Group("/example")
	{
		responseExample := example.Group("/response")
		{
			responseExample.GET("/", api.Index)
			responseExample.GET("/success-message", api.ResponseSuccessWithMessage)
			responseExample.GET("/fail", api.ResponseFail)
			responseExample.GET("/fail-code", api.ResponseFailWithCode)
			responseExample.GET("/fail-status", api.ResponseFailWithStatus)
		}
	}

	return engine
}
