package api

import (
	"go-web/internal/response"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	response.Success(c, gin.H{
		"message": "go-web service",
	})
}

func Health(c *gin.Context) {
	response.Success(c, gin.H{
		"status": "ok",
	})
}
