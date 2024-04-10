package requests

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func init() {
	govalidator.TagMap["fixedPhone"] = govalidator.Validator(func(str string) bool {
		return len(str) == 11
	})
	govalidator.TagMap["LengthBetween"] = govalidator.Validator(func(str string) bool {
		length := len(str)
		return length >= 4 && length <= 100
	})
}

// SignUpRequest represents the expected fields for a sign-up request
type SignUpRequestPhone struct {
	Phone string `json:"phone" valid:"required~手机号码不能为空,fixedPhone~手机号码格式不正确"`
}

type SignUpRequestEmail struct {
	Email string `json:"email" valid:"required~邮箱不能为空,email~邮箱格式不正确,LengthBetween~邮箱长度必须在4-100位之间"`
}

// ValidateSignUpRequest validates the sign-up request data
func ValidateSignUpRequestPhone(c *gin.Context, request SignUpRequestPhone) (SignUpRequestPhone, error) {
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return request, err
	}
	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		return request, err
	}

	return request, nil
}

func ValidateSignUpRequestEmail(c *gin.Context, request SignUpRequestEmail) (SignUpRequestEmail, error) {
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return request, err
	}
	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		return request, err
	}

	return request, nil
}
