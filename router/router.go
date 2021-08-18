package router

import (
	"github.com/DavinciUI/backend/config"
	"github.com/DavinciUI/backend/router/middleware"
	"github.com/gin-gonic/gin"
)

func CreateRouter(config config.Config) (*gin.Engine, error) {
	router := gin.Default()

	// Middleware for Rate-limit
	router.Use(middleware.Limiter(config.RateLimit))

	// Middleware for prevent CSRF attacks
	router.Use(middleware.CSRF(config.AuthKey))

	// Middleware for Logs
	//router.Use(middleware.Logger()) Gin already includes a Logger middleware

	router.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello")
	})

	router.POST("/test", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello")
	})

	return router, router.Run(config.RouterPort)
}