package main

import (
	"ownapihub/bootstrap"
	"ownapihub/config"
	"ownapihub/pkg"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Initialize()
}

func main() {
	// load env file
	var env string
	pkg.InitConfig(env)

	r := gin.Default()
	bootstrap.InitRoute(r)
	bootstrap.Initdatabase()
	

	r.Run(":" + pkg.Get("app.port"))
}
