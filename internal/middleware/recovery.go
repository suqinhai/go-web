package middleware

import (
	"log"
	"net/http"

	"go-web/internal/response"

	"github.com/gin-gonic/gin"
)

// Recovery 捕获 handler 或 middleware 中的 panic。
// 发生 panic 时，它会记录错误，并返回统一格式的 500 JSON 响应，避免服务进程直接崩溃。
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[Recovery] panic recovered: %v", err)
				response.FailWithStatus(c, http.StatusInternalServerError, response.CodeFail, "internal server error")
				c.Abort()
			}
		}()

		c.Next()
	}
}
