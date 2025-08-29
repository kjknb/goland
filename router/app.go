package router

import (
	"GInchat/docs"
	"GInchat/service" // 导入你存放处理函数的包
	"GInchat/utils"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	// 创建默认路由
	r := gin.Default()

	//首页
	r.GET("/", service.GetIndex)

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 定义路由
	r.GET("/index", service.GetIndex)

	// 公开路由 - 不需要token认证
	public := r.Group("/user")
	{
		public.POST("/findUserByNameAndPwd", service.FindUserByNameAndPwd)
		public.POST("/createUser", service.CreateUser)
	}

	// 受保护的路由 - 需要token认证
	protected := r.Group("/user")
	protected.Use(utils.JWTAuthMiddleware()) // 添加JWT中间件
	{
		protected.POST("/GetUserList", service.GetUserList)
		protected.DELETE("/deleteUser", service.DeleteUser)
		protected.PUT("/updateUser", service.UpdateUser)
	}

	return r
}
