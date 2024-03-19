package auth

import (
	v1 "ownapihub/app/http/controllers/api/v1"
	"ownapihub/models/user"

	"github.com/gin-gonic/gin"
)

type SingUp struct {
	v1.BaseAPIController
}

func (s *SingUp) SingUpUsingPhone(c *gin.Context)  {
	type phoneRequest struct {
		Phone string `json:"phone"`
	}
	request := &phoneRequest{}
	// partse json request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"answer": user.CheckUsersWithPhone(request.Phone),
	})
	
}