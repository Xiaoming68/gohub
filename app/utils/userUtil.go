package utils

import (
	"gohub/app/models/users"
	"gohub/initPack/database"
)

// IsEmailExist 验证邮箱是否存在
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(users.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(users.User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
