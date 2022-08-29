package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/database"
	"gohub/pkg/verifycode"
	"net/http"
)

type ResetByEmailRequest struct {
	Email string `json:"email" valid:"email"`

	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

func ResetByEmail(c *gin.Context, data interface{}) bool {

	// 1. 验证数据格式
	if err := c.ShouldBindJSON(data); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"error":   err.Error(),
		})

		return false
	}
	// 2 验证表单输入
	rules := govalidator.MapData{
		"email":       []string{"required", "email", "min:4", "max:30"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:邮箱为必填项",
			"email:输入为邮箱格式",
			"min:最少输入4个字符",
			"max:最多输入30个长度",
		},
		"verify_code": []string{
			"required:验证码为必填项",
			"digits:请输入6位验证",
		},
		"password": []string{
			"required:密码为必填项",
			"min:最少输入6个字符",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*ResetByEmailRequest)

	// 3. 验证邮箱是否存在
	var count int64
	database.DB.Where("email = ?", _data.Email).Count(&count)
	if count == 0 {
		errs["password"] = append(errs["password"], "该邮箱不存在")
	}

	// 4. 验证验证码是否正确
	if ok := verifycode.NewVerifyCode().CheckAnswer(_data.Email, _data.VerifyCode); !ok {
		errs["verify_code"] = append(errs["verify_code"], "邮箱验证码错误")
		return false
	}

	return true
}
