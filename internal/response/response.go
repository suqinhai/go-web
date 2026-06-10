package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess      = 0
	CodeFail         = 1
	CodeUnauthorized = 40101
	CodeForbidden    = 40301
	CodeServerError  = 50001
)

const (
	MessageSuccess = "success"
	MessageFail    = "fail"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(c *gin.Context, data any) {
	JSON(c, http.StatusOK, CodeSuccess, MessageSuccess, data)
}

func SuccessWithMessage(c *gin.Context, message string, data any) {
	JSON(c, http.StatusOK, CodeSuccess, message, data)
}

func Fail(c *gin.Context, message string) {
	JSON(c, http.StatusOK, CodeFail, message, gin.H{})
}

func FailWithCode(c *gin.Context, code int, message string) {
	JSON(c, http.StatusOK, code, message, gin.H{})
}

func FailWithStatus(c *gin.Context, statusCode, code int, message string) {
	JSON(c, statusCode, code, message, gin.H{})
}

func JSON(c *gin.Context, statusCode, code int, message string, data any) {
	if data == nil {
		data = gin.H{}
	}

	c.JSON(statusCode, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
