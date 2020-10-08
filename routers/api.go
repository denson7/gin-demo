package routers

import (
	"gin-demo/controller"
	"gin-demo/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine)  {
	//首页
	r.GET("/", controller.Index)
	r.POST("/check", middleware.IpBlack, controller.Index)

	//上传文件
	r.POST("/upload", controller.UploadImg)

	// IP
	v1 := r.Group("/v1")
	{
		v1.GET("/getIpBlackInfo", controller.GetIpBlackInfo)
		v1.POST("/addIpBlack", controller.AddIpBlack)
		v1.PUT("/updateIpBlack", controller.UpdateIpBlack)
		v1.DELETE("/delIpBlack", controller.DelIpBlack)
		v1.GET("/getIpBlackList", controller.GetIpBlackList)
	}

}
