package middleware

import (
	"fmt"
	"net/http"

	"go-web/internal/response"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// NewCasbin 根据配置文件初始化 Casbin 权限校验中间件。
// modelPath 是权限模型文件，policyPath 是权限策略文件。
func NewCasbin(modelPath, policyPath string) (gin.HandlerFunc, error) {
	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		return nil, fmt.Errorf("new casbin enforcer: %w", err)
	}

	if err := enforcer.LoadPolicy(); err != nil {
		return nil, fmt.Errorf("load casbin policy: %w", err)
	}

	return Casbin(enforcer), nil
}

// Casbin 根据当前请求的角色、路径和方法判断是否允许访问。
// 它依赖 Tenant 中间件提前写入 role_type，所以使用顺序应为 JWT -> Tenant -> Casbin。
func Casbin(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		tenant, ok := GetTenantContext(c)
		if !ok || tenant.RoleType == "" {
			response.FailWithStatus(c, http.StatusUnauthorized, response.CodeUnauthorized, "missing tenant context")
			c.Abort()
			return
		}

		path := c.Request.URL.Path
		method := c.Request.Method

		allowed, err := enforcer.Enforce(tenant.RoleType, path, method)
		if err != nil {
			response.FailWithStatus(c, http.StatusInternalServerError, response.CodeServerError, "permission check failed")
			c.Abort()
			return
		}
		if !allowed {
			response.FailWithStatus(c, http.StatusForbidden, response.CodeForbidden, "permission denied")
			c.Abort()
			return
		}

		c.Next()
	}
}
