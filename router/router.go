package router

import (
	"github.com/DavinciUI/backend/config"
	"github.com/DavinciUI/backend/router/middleware"
	"github.com/gin-gonic/gin"
)

func CreateRouter(config config.Config) (*gin.Engine, error) {
	router := gin.New()

	// Middleware for Rate-limit
	router.Use(middleware.MaxRequests(config.RateLimit))

	// Middleware for Logs
	router.Use(middleware.Logger())

	return router, router.Run(config.RouterPort)
}