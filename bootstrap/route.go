package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/middlewares"
	"gohub/initPack/config"
	"gohub/routes"
	"net/http"
)

func SetUpRouter(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(router)

	// 注册API路由
	routes.RegisterAPIRoutes(router)

	// 配置404路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
	err := router.SetTrustedProxies([]string{
		config.Get("app.host"),
	})
	if err != nil {
		panic(err.Error())
	}
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error_code":    404,
			"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
		})

	})
}
