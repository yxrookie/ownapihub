package routes

import "github.com/gin-gonic/gin"

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func GetRoutes(engine *gin.Engine) {
	v1 := engine.Group("/v1")
	{
		v1.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"new": "life",
			})
		})
	}

}
