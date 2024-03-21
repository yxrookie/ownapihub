package routes

import (
	"ownapihub/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func GetRoutes(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		suc := new(auth.SingUp)
		// 判断手机是否已注册
		v1.POST("auth/signup/phone/exist", suc.SingUpUsingPhone)
		// 判断邮箱是否已注册
		v1.POST("auth/signup/email/exist", suc.SingUpUsingEmail)	
	}

}
