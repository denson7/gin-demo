package middleware

import (
	"gin-demo/models"
	"github.com/gin-gonic/gin"
)

func IpBlack(c *gin.Context) {
	ip := c.ClientIP()
	ipBlack := models.FindIp(ip)
	if ipBlack.IP != "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "IP已被加入黑名单",
		})
		c.Abort()
		return
	}
}

