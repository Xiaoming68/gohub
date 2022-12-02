package utils

import (
	"gohub/app/models/users"
	"gohub/initPack/database"
)

// IsDataExist 验证数据是否存在
func IsDataExist(field string, parameter string) bool {
	var count int64
	database.DB.Model(users.User{}).Where("? = ?", field, parameter).Count(&count)
	return count > 0
}
