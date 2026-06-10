package api

import (
	"net/http"

	"go-web/internal/response"

	"github.com/gin-gonic/gin"
)

func ResponseSuccessWithMessage(c *gin.Context) {
	response.SuccessWithMessage(c, "custom success message", gin.H{
		"example": "SuccessWithMessage",
	})
}

func ResponseFail(c *gin.Context) {
	response.Fail(c, "fail example")
}

func ResponseFailWithCode(c *gin.Context) {
	response.FailWithCode(c, 40001, "fail with custom business code")
}

func ResponseFailWithStatus(c *gin.Context) {
	response.FailWithStatus(c, http.StatusBadRequest, 40002, "fail with http status")
}
