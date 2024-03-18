package bootstrap

import (
	"fmt"
	"net/http"
	"ownapihub/routes"

	"github.com/gin-gonic/gin"
)

func InitRoute(engine *gin.Engine) {
	Register404Route(engine)
	engine.Use(CustomErrorHandler())
	routes.GetRoutes(engine)
}

// Register404Route 用于注册处理 404 错误的路由处理函数
func Register404Route(engine *gin.Engine) {
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})
}

// CustomErrorHandler 自定义错误处理中间件
func CustomErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 发生异常时的处理逻辑
				fmt.Println("Internal Server Error:", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				c.Abort() // 终止后续中间件和处理函数的执行
			}
		}()
		c.Next() // 执行下一个中间件或处理函数
	}
}

