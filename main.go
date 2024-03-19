package main

import (
	"ownapihub/bootstrap"
	"ownapihub/config"
	"ownapihub/models/user"
	"ownapihub/pkg"
	"ownapihub/pkg/database"

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
	database.Db, database.SqlDB = bootstrap.Initdatabase()
	database.Db.AutoMigrate(&user.User{})
	

	r.Run(":" + pkg.Get("app.port"))
}
