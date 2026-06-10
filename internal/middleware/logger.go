package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 记录每一次 HTTP 请求的基础信息。
// 日志包含时间、状态码、耗时、客户端 IP、请求方法和路径，方便开发阶段排查接口问题。
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		line := fmt.Sprintf("[GIN] %s | %3d | %13v | %15s | %-7s %s",
			param.TimeStamp.Format(time.RFC3339),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
		)

		if param.ErrorMessage != "" {
			line += " | " + param.ErrorMessage
		}

		return line + "\n"
	})
}
