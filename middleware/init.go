package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(r *gin.Engine) {
	// 跨域处理
	r.Use(Options)
	// Set X-Request-Id header
	r.Use(RequestId())
}
