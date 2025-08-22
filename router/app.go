package router

import (
	"GInchat/service" // 导入你存放处理函数的包
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// 创建默认路由
	r := gin.Default()

	// 定义路由
	r.GET("/index", service.GetIndex)
	// 可以继续定义更多路由
	// r.POST("/users", service.CreateUser)
	// r.GET("/users/:id", service.GetUserByID)

	return r
}
