package controller

import (
	"fmt"
	"gin-demo/models"
	"github.com/gin-gonic/gin"
	"strconv"
)
// curl -X GET "http://localhost:8080/v1/getIpBlackInfo?ip=191.168.1.111"
func GetIpBlackInfo(c *gin.Context)  {
	ip := c.Query("ip")
	if ip == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "please input ip!",
		})
		return
	}
	info := models.FindIp(ip)
	fmt.Println(info)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok!",
		"data": info,
	})
}

// curl -X POST "http://localhost:8080/v1/addIpBlack" -d "ip=191.168.1.111&user_id=a1"
func AddIpBlack(c *gin.Context) {
	ip := c.PostForm("ip")
	fmt.Println(ip)
	if ip == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "请输入IP!",
		})
		return
	}
	userId := c.PostForm("user_id")
	fmt.Println(userId)
	insertId := models.CreateIpBlack(ip, userId)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "add success!",
		"data": insertId,
	})
}

// curl -X PUT "http://localhost:8080/v1/updateIpBlack" -d "ip=191.168.1.999&user_id=a1&id=2"
func UpdateIpBlack(c *gin.Context)  {
	userId := c.PostForm("user_id")
	ip := c.PostForm("ip")
	id, _ := strconv.Atoi(c.PostForm("id"))
	updateId := models.UpdateIpBlack(ip, userId, id)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "update!",
		"data": updateId,
	})
}

// curl -X DELETE "http://localhost:8080/v1/delIpBlack?ip=191.168.1.111"
func DelIpBlack(c *gin.Context) {
	ip := c.Query("ip")
	if ip == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "please input IP!",
		})
		return
	}
	models.DeleteIpBlackByIp(ip)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "delete ip success!",
	})
}

// curl -X GET "http://localhost:8080/v1/getIpBlackList?page=1&pageSize=2"
func GetIpBlackList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	list := models.FindIps(nil, nil, uint(page), uint(pageSize))
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"result": gin.H{
			"list":     list,
			"pageSize": pageSize,
		},
	})
}
