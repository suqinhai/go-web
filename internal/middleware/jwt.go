package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"go-web/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	AuthClaimsKey  = "auth_claims"
	UserIDKey      = "user_id"
	RoleTypeKey    = "role_type"
	MerchantIDKey  = "merchant_id"
	AgentIDKey     = "agent_id"
	PermissionsKey = "permissions"
)

// JWTClaims 定义业务 JWT 中携带的用户身份信息。
// 这些字段会在 JWT 中间件解析成功后写入 Gin Context，供 Tenant 和业务 handler 使用。
type JWTClaims struct {
	UserID      string   `json:"user_id"`
	RoleType    string   `json:"role_type"`
	MerchantID  string   `json:"merchant_id,omitempty"`
	AgentID     string   `json:"agent_id,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	jwt.RegisteredClaims
}

// JWT 校验 Authorization: Bearer <token> 请求头。
// 校验通过后，它会把用户 ID、角色、商户、代理和权限等信息写入 Gin Context。
func JWT(secret, issuer string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if secret == "" {
			response.FailWithStatus(c, http.StatusInternalServerError, response.CodeServerError, "jwt secret is empty")
			c.Abort()
			return
		}

		tokenText := bearerToken(c.GetHeader("Authorization"))
		if tokenText == "" {
			response.FailWithStatus(c, http.StatusUnauthorized, response.CodeUnauthorized, "missing authorization token")
			c.Abort()
			return
		}

		claims := &JWTClaims{}
		options := []jwt.ParserOption{}
		if issuer != "" {
			options = append(options, jwt.WithIssuer(issuer))
		}

		token, err := jwt.ParseWithClaims(tokenText, claims, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		}, options...)
		if err != nil || !token.Valid {
			response.FailWithStatus(c, http.StatusUnauthorized, response.CodeUnauthorized, "invalid authorization token")
			c.Abort()
			return
		}

		c.Set(AuthClaimsKey, claims)
		c.Set(UserIDKey, claims.UserID)
		c.Set(RoleTypeKey, claims.RoleType)
		c.Set(MerchantIDKey, claims.MerchantID)
		c.Set(AgentIDKey, claims.AgentID)
		c.Set(PermissionsKey, claims.Permissions)

		c.Next()
	}
}

// GetJWTClaims 从 Gin Context 中读取 JWT 中间件解析后的 claims。
func GetJWTClaims(c *gin.Context) (*JWTClaims, bool) {
	value, exists := c.Get(AuthClaimsKey)
	if !exists {
		return nil, false
	}

	claims, ok := value.(*JWTClaims)
	return claims, ok
}

func bearerToken(header string) string {
	fields := strings.Fields(header)
	if len(fields) != 2 || !strings.EqualFold(fields[0], "Bearer") {
		return ""
	}

	return fields[1]
}
