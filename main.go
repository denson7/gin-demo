package main

import (
	"gin-demo/middleware"
	"gin-demo/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Engine
	r := gin.Default()
	// 加载中间件
	middleware.InitMiddleware(r)

	// 加载路由
	routers.InitRouter(r)

	// 默认:8080端口
	r.Run()

}
