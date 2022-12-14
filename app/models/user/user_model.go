package user

import models "gohub/app/models"

type User struct {
	models.BaseModel

	// 敏感信息不想输出给用户，"-" 这是在指示 JSON 解析器忽略字段 。后面接口返回用户数据时候，这三个字段都会被隐藏
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
