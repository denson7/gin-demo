package main

import (
	"gin-demo/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建Engine
	r := gin.Default()
	// 加载路由
	routers.InitApiRouter(r)
	// 默认:8080端口
	r.Run()

}
