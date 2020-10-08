package middleware

import (
	"fmt"
	"gin-demo/utils"
	"github.com/gin-gonic/gin"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for incoming header, use it if exists
		requestId := c.Request.Header.Get("X-Request-Id")
		fmt.Println("requestId:-------->", requestId)
		// Create request id with UUID
		if requestId == "" {
			uuid, _ := utils.NewUUID()
			requestId = uuid.String()
		}
		fmt.Println("requestId:-------->", requestId)
		// Expose it for use in the application
		c.Set("X-Request-Id", requestId)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
