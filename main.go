package main

import (
	"GInchat/router"
	"GInchat/utils"
)

func main() {
	utils.InitConfig()
	utils.Initmysql()
	utils.InitRedis()
	// 1. 初始化路由
	r := router.Router()
	// 2. 可以在这里做其他初始化工作，比如连接数据库等

	// 3. 启动服务器，监听端口
	r.Run(":8081") // listen and serve on 0.0.0.0:8081
}
