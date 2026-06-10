package api

import (
	"go-web/internal/middleware"
	"go-web/internal/response"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	tenant, ok := middleware.GetTenantContext(c)
	if !ok {
		response.Fail(c, "missing tenant context")
		return
	}

	response.Success(c, tenant)
}
