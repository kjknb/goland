package router

import (
	"GInchat/docs"
	"GInchat/service" // 导入你存放处理函数的包
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	// 创建默认路由
	r := gin.Default()

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 定义路由
	r.GET("/index", service.GetIndex)
	// 可以继续定义更多路由
	// r.POST("/users", service.CreateUser)
	r.GET("/user/GetUserList", service.GetUserList)

	return r
}
