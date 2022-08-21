package auth

import (
	"fmt"
	v1 "gohub/app/http/controllers/v1"
	"gohub/app/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseApiController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExists(c *gin.Context) {
	// 请求对象
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}

	request := PhoneExistRequest{}

	// 解析 JSON 请求
	if err := c.ShouldBindJSON(&request); err != nil {
		// 解析失败，返回 422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exists": user.IsPhoneExists(request.Phone),
	})
}
