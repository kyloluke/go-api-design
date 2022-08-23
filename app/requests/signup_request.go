package requests

import (
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}
type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

// ValidatePhoneIsExist 验证手机输入是否正确
func ValidatePhoneIsExist(data interface{}) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required: 手机号为必填项，参数名称为 phone",
			"digits: 手机号长度必须为 11 位数字",
		},
	}

	return validate(data, rules, messages)
}

// ValidateIsEmailExist 验证邮箱输入是否正确
func ValidateIsEmailExist(data interface{}) map[string][]string {
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email为必填项，参数名称 email",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}

	return validate(data, rules, messages)
}
