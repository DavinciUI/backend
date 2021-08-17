package middleware

import (
	"github.com/gin-gonic/gin"
)

var requestChannel = make(chan byte)

func MaxRequests(max int) gin.HandlerFunc {
	if len(requestChannel) > max {
		return abort
	} else {
		return next
	}
}

func abort(ctx *gin.Context) {
	ctx.JSON(411, "LIMIT_REQUEST_EXCEED")
	ctx.Abort()
}

func next(ctx *gin.Context) {
	requestChannel <- 1
	ctx.Next()
	<- requestChannel
}