package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var rateLimiter *rate.Limiter

// Limiter middleware to use in Global Router
func Limiter(max int) gin.HandlerFunc {
	rateLimiter = rate.NewLimiter(rate.Limit(max), max)
	return limiterHandler
}

func limiterHandler(ctx *gin.Context) {
	if rateLimiter.Allow() {
		ctx.Next()
	} else {
		ctx.JSON(429, "LIMIT_REQUEST_EXCEED")
		ctx.Abort()
	}
}
