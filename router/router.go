package router

import "github.com/gin-gonic/gin"
import "ginProject/ginProject/controller"

func init() {
	gin.DisableConsoleColor()
	//使用默认中间件创建一个gin路由
	//日志与恢复中间件
	router := gin.Default()
	//为multipart 表单设置一个较低的内存限制
	router.MaxMultipartMemory = 8 << 20
	v1 := router.Group("/v1")
	{
		v1.POST("/login")
		v1.GET("/someGet/:name", controller.Getting)
		v1.GET("/someGet/:name/*action", controller.AnotherGetting)
		v1.GET("/welcome", controller.StrGetting)
		v1.POST("/formPost", controller.PostFunc)
		v1.POST("/aformPost", controller.ApostFunc)
		v1.POST("/upload", controller.UploadFunc)
		v1.POST("/multiUpload", controller.MultiUploadFunc)
	}
	router.Run(":3000")
}
