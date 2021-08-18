package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	wraphh "github.com/turtlemonvh/gin-wraphh"
)

func CSRF(key string) gin.HandlerFunc {
	return wraphh.WrapHH(
		csrf.Protect(
			[]byte(key),
			csrf.RequestHeader("Auth-Token"),
			csrf.FieldName("token"),
		))
}