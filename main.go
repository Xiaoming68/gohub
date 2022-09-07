package main

import (
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
	_ "gohub/configs"
	baseCon "gohub/initPack/config"
)

func init() {
	// 加载CONFIG配置信息
	baseCon.InitConfig("")
}

func main() {

	// 初始化 Logger
	bootstrap.SetupLogger()

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// 初始化GIN擎实例
	router := gin.New()

	// 初始化DB定义
	bootstrap.SetupDB()

	// 初始化 Redis
	bootstrap.SetupRedis()

	// 初始化路由定义
	bootstrap.SetUpRouter(router)

	err := router.Run(":" + baseCon.Get("app.port"))
	if err != nil {
		panic(err.Error())
	}
}
