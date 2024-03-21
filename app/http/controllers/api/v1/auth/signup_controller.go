package auth

import (
	"net/http"
	v1 "ownapihub/app/http/controllers/api/v1"
	"ownapihub/app/requests"
	"ownapihub/models/user"

	"github.com/gin-gonic/gin"
)

type SingUp struct {
	v1.BaseAPIController
}

func (s *SingUp) SingUpUsingPhone(c *gin.Context)  {
	// partse json request
	request, err := requests.ValidateSignUpRequestPhone(c, requests.SignUpRequestPhone{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"exist": user.CheckUsersWithPhone(request.Phone),
	})
	
}

func (s *SingUp) SingUpUsingEmail(c *gin.Context)  {
	// partse json request
	request, err := requests.ValidateSignUpRequestEmail(c, requests.SignUpRequestEmail{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"exist": user.CheckUsersWithEmail(request.Email),
	})
	
}