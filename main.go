package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func testing(c *gin.Context) {
	var person *Person
	if err := c.ShouldBind(&person); err == nil {
		c.String(200, "%v", person)
	} else {
		c.String(200, "person bind error: %v", err )
	}
}

func main()  {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 请求参数
	//r.GET("/:name/:id", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"name": context.Param("name"),
	//		"id": context.Param("id"),
	//	})
	//})
	// 绑定前缀
	//r.GET("/user/*action", func(context *gin.Context) {
	//	context.String(200, "hello")
	//})
	// 获取GET URL中的参数
	r.GET("/test", func(context *gin.Context) {
		firstName := context.Query("first_name")
		lastName := context.DefaultQuery("last_name", "last_default_name")
		context.String(http.StatusOK, "%s, %s", firstName, lastName)
	})
	r.POST("/post", func(context *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.String(http.StatusBadRequest, err.Error())
			context.Abort()
		}
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		firstName := context.PostForm("first_name")
		lastName := context.DefaultPostForm("last_name", "default_last_name")
		fmt.Println(firstName)
		fmt.Println(lastName)
		context.String(http.StatusOK,  string(bodyBytes))
	})
	// GET
	r.GET("/get", func(context *gin.Context) {
		context.String(200, "get")
	})
	// POST
	r.POST("/post", func(context *gin.Context) {
		context.String(200, "post")
	})
	// DELETE
	r.Handle("DELETE","/delete", func(context *gin.Context) {
		context.String(200, "delete")
	})
	// 允许所有请求
	r.Any("/any", func(context *gin.Context) {
		context.String(200, "any")
	})
	r.Static("/assets", "./assets")
	r.StaticFS("/static", http.Dir("static"))
	r.StaticFile("/favicon.ico", "./favicon.ico")

	r.GET("/testing", testing)
	r.POST("/testing", testing)
	r.Run()
}
