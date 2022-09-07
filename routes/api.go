// Package routes 注册路由
package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/v1/auth"
)

func RegisterAPIRoutes(r *gin.Engine) {

	// v1 版本的路由组
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/user")
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	}
}
