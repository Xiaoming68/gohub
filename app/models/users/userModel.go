package users

import (
	"gohub/app/models"
)

type User struct {
	models.BaseModel // 主键默认id

	Name string `json:"name,omitempty"`
	// "-" 表示JSON 解析器忽略字段，返回用户数据时候，以下三个字段都会被隐藏
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField // 时间默认 created_at、updated_at
}
