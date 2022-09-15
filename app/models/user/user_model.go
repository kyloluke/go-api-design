package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"gohub/pkg/hash"
)

type User struct {
	models.BaseModel

	// 敏感信息不想输出给用户，"-" 这是在指示 JSON 解析器忽略字段 。后面接口返回用户数据时候，这三个字段都会被隐藏
	Name string `json:"name,omitempty"`

	City         string `json:"city,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Avatar       string `json:"avatar,omitempty"`

	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

func (userModel *User) Create() {
	database.DB.Create(userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(password string) bool {
	return hash.BcryptCheck(password, userModel.Password)
}

// Get 通过 ID 获取用户
func Get(idStr string) (userModel User) {
	database.DB.Where("id = ?", idStr).First(&userModel)
	return
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(userModel)
	return result.RowsAffected
}

func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}

func All() (users []User) {
	database.DB.Find(&users)
	return
}
