package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var dateFormat = "00/00/0000 00:00:00"

func Logger() gin.HandlerFunc {
	return printInfo
}

func printInfo(ctx *gin.Context) {
	requestCopy := ctx.Request

	fmt.Println(
		"\n",
		"date: ", time.Now().Format(dateFormat), "\n",
		"url:", requestCopy.URL, "\n",
		"address:", requestCopy.Host, "\n",
		"method:", requestCopy.Method, "\n",
		"body:", requestCopy.Body,
	)
}
