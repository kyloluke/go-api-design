package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/auth"
)

type UserRequest struct {
	Name         string `valid:"name" json:"name"`
	City         string `valid:"city" json:"city"`
	Introduction string `valid:"introduction" json:"introduction"`
}

func UserSave(data interface{}, c *gin.Context) map[string][]string {
	uid := auth.CurrentUID(c)
	rules := govalidator.MapData{
		"name":         []string{"required", "alpha_num", "between:3,20", "not_exists:users,name," + uid},
		"city":         []string{"min_cn:2", "max_cn:20"},
		"introduction": []string{"min_cn:4", "max_cn:240"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"alpha_num:格式错误，只允许英文和数字",
			"between:名称长度3-20之间",
			"not_exists:名称已存在",
		},
		"city": []string{
			"min_cn:描述长度需至少 2 个字",
			"max_cn:描述长度不能超过 20 个字",
		},
		"introduction": []string{
			"min_cn:描述长度需至少 4 个字",
			"max_cn:描述长度不能超过 240 个字",
		},
	}
	return validate(data, rules, messages)
}
