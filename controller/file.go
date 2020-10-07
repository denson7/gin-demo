package controller

import (
	"gin-demo/utils"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strings"
	//"time"
)

func UploadImg(c *gin.Context) {
	//config := config.CreateConfig()
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "上传失败!",
		})
		return
	} else {
		fileExt := strings.ToLower(path.Ext(file.Filename))
		if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "上传失败!只允许png,jpg,gif,jpeg文件",
			})
			return
		}
		//fileName := tools.Md5(fmt.Sprintf("%s%s", f.Filename, time.Now().String()))
		//fileDir := fmt.Sprintf("%s%d%s/", config.Upload, time.Now().Year(), time.Now().Month().String())
		fileDir := "./uploads/"
		dst := path.Join(fileDir, file.Filename)
		isExist, _ := utils.IsFileExist(fileDir)
		if !isExist {
			os.Mkdir(fileDir, os.ModePerm)
		}
		_ = c.SaveUploadedFile(file, dst)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "上传成功!",
			"result": gin.H{
				"path": dst,
			},
		})
	}
}
