package router

import (
	"go-web/internal/api"
	"go-web/internal/config"
	"go-web/internal/middleware"

	"github.com/gin-gonic/gin"
)

func New(cfg config.Config) (*gin.Engine, error) {
	engine := gin.New()
	engine.Use(middleware.Logger(), middleware.Recovery(), middleware.CORS())

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

	authz, err := middleware.NewCasbin(cfg.Casbin.ModelPath, cfg.Casbin.PolicyPath)
	if err != nil {
		return nil, err
	}

	protected := engine.Group("/api")
	protected.Use(middleware.JWT(cfg.JWT.Secret, cfg.JWT.Issuer), middleware.Tenant(), authz)
	{
		protected.GET("/profile", api.Profile)
	}

	return engine, nil
}
