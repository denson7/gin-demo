package controller

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.Redirect(302, "/index/en")
}


func Main(c *gin.Context) {
	//c.Redirect(302, "/index/en")
}