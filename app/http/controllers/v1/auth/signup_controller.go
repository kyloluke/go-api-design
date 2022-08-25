package auth

import (
	v1 "gohub/app/http/controllers/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseApiController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 初始化请求对象
	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidatePhoneIsExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exists": user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {

	request := requests.SignupEmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateIsEmailExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})

}
