package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/auth"
	"gohub/pkg/verifycode"
	"net/http"
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

type UserUpdateEmailRequest struct {
	Email      string `valid:"email" json:"email,omitempty"`
	VerifyCode string `valid:"verify_code" json:"verify_code,omitempty"`
}

func UserUpdateEmailValidate(data interface{}, c *gin.Context) map[string][]string {
	// 违反的验证规则项都会同时抛出
	currentUser := auth.CurrentUser(c)
	rules := govalidator.MapData{
		"email": []string{
			"required",
			"max:30",
			"email",
			"not_exists:users,email," + currentUser.GetStringID(),
			"not_in:" + currentUser.Email},
		"verify_code": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:邮箱为必填项目",
			"max:长度需小于30",
			"email:请输入邮箱格式",
			"not_exists:邮箱已被占用",
			"not_in:新邮箱不能与老邮箱一致",
		},
		"verify_code": []string{
			"required:验证码必填项",
			"digits:请输入6位数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*UserUpdateEmailRequest)

	if ok := verifycode.NewVerifyCode().CheckAnswer(_data.Email, _data.VerifyCode); !ok {
		errs["verify_code"] = append(errs["verify_code"], "验证码错误")
	}

	return errs
}

type UpdatePasswordRequest struct {
	CurrentPassword    string `valid:"current_password" json:"current_password,omitempty"`
	NewPassword        string `valid:"new_password" json:"new_password,omitempty"`
	NewPasswordConfirm string `valid:"new_password_confirm" json:"new_password_confirm,omitempty"`
}

// UpdatePasswordValidate 密码修改
func UpdatePasswordValidate(data interface{}, c *gin.Context) bool {

	// 数据绑定
	if err := c.ShouldBindJSON(data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "数据解析错误",
			"err":     err.Error(),
		})
		return false
	}

	// 数据验证
	rules := govalidator.MapData{
		"current_password":     []string{"required", "min:6"},
		"new_password":         []string{"required", "min:6"},
		"new_password_confirm": []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"current_password":     []string{"required:密码是必填项", "min:最少6位"},
		"new_password":         []string{"required:新密码是必填项", "min:最少6位"},
		"new_password_confirm": []string{"required:确认新密码为必填项", "min:最少6位"},
	}

	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	errs := govalidator.New(opts).ValidateStruct()

	// 新密码的两次输入是否一致
	_data := data.(*UpdatePasswordRequest)
	if _data.NewPassword != _data.NewPasswordConfirm {
		errs["new_password_confirm"] = append(errs["new_password_confirm"], "新密码不一致")
	}

	if len(errs) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "数据验证通过",
			"errors":  errs,
		})
		return false
	}
	return true
}
