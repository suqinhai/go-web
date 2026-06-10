package middleware

import (
	"net/http"
	"strings"

	"go-web/internal/response"

	"github.com/gin-gonic/gin"
)

const TenantContextKey = "tenant_context"

// TenantContext 表示当前请求的租户和身份上下文。
// 多商户、多代理或后台角色相关的权限判断，都可以基于这个结构继续扩展。
type TenantContext struct {
	UserID      string   `json:"user_id"`
	RoleType    string   `json:"role_type"`
	MerchantID  string   `json:"merchant_id,omitempty"`
	AgentID     string   `json:"agent_id,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

// Tenant 生成当前请求的租户上下文。
// 它优先使用 JWT claims 中的信息，也支持从 X-User-ID、X-Role-Type 等请求头读取，便于本地调试。
func Tenant() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenant := tenantFromClaims(c)
		if tenant.UserID == "" {
			tenant.UserID = c.GetHeader("X-User-ID")
		}
		if tenant.RoleType == "" {
			tenant.RoleType = c.GetHeader("X-Role-Type")
		}
		if tenant.MerchantID == "" {
			tenant.MerchantID = c.GetHeader("X-Merchant-ID")
		}
		if tenant.AgentID == "" {
			tenant.AgentID = c.GetHeader("X-Agent-ID")
		}
		if len(tenant.Permissions) == 0 {
			tenant.Permissions = splitHeader(c.GetHeader("X-Permissions"))
		}

		if tenant.UserID == "" || tenant.RoleType == "" {
			response.FailWithStatus(c, http.StatusUnauthorized, response.CodeUnauthorized, "missing tenant context")
			c.Abort()
			return
		}

		c.Set(TenantContextKey, tenant)
		c.Next()
	}
}

// GetTenantContext 从 Gin Context 中读取 Tenant 中间件生成的租户上下文。
func GetTenantContext(c *gin.Context) (TenantContext, bool) {
	value, exists := c.Get(TenantContextKey)
	if !exists {
		return TenantContext{}, false
	}

	tenant, ok := value.(TenantContext)
	return tenant, ok
}

func tenantFromClaims(c *gin.Context) TenantContext {
	claims, ok := GetJWTClaims(c)
	if !ok {
		return TenantContext{}
	}

	return TenantContext{
		UserID:      claims.UserID,
		RoleType:    claims.RoleType,
		MerchantID:  claims.MerchantID,
		AgentID:     claims.AgentID,
		Permissions: claims.Permissions,
	}
}

func splitHeader(value string) []string {
	if value == "" {
		return nil
	}

	parts := strings.Split(value, ",")
	items := make([]string, 0, len(parts))
	for _, part := range parts {
		item := strings.TrimSpace(part)
		if item != "" {
			items = append(items, item)
		}
	}

	return items
}
