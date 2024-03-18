package main

import (
	"ownapihub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	bootstrap.InitRoute(r)
	r.Run(":3000")
}