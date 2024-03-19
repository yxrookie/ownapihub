package user

import "ownapihub/models"

type User struct {
	models.BaseModel
	
	Name string
	Password string
	Phone string
	Email string
	
	models.CommonTimestampsField// 嵌入BaseModel
}