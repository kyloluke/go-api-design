package auth

import (
	"fmt"
	v1 "gohub/app/http/controllers/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseApiController
}

// IsPhoneExists 检测手机号是否被注册
func (sc *SignupController) IsPhoneExists(c *gin.Context) {

	// 初始化请求对象
	request := requests.SignupPhoneExistRequest{}

	// 解析 JSON 格式，非json格式抛出错误
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}

	// 表单验证
	errs := requests.ValidatePhoneIsExists(&request, c)
	if len(errs) > 0 {
		// 验证失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exists": user.IsPhoneExists(request.Phone),
	})
}
