package user

import (
	"ownapihub/pkg/database"

)

func CheckUsersWithPhone(phone string) bool {
    var count int64
	database.Db.Model(&User{}).Where("phone = ?", phone).Count(&count)
    return count > 0
}

func CheckUsersWithEmail(email string) bool {
    var count int64
    database.Db.Model(&User{}).Where("email = ?", email).Count(&count)
    return count > 0
}