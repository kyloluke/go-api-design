package auth

import (
	"github.com/gin-gonic/gin"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/jwt"
	"gohub/pkg/response"
)

type LoginController struct {
}

func (lc *LoginController) LoginByPassword(c *gin.Context) {

	// 1. 各类验证
	var request = requests.LoginByPasswordRequest{}

	ok := requests.Validate(c, &request, requests.LoginByPassword)
	if !ok {
		return
	}

	// 2. 登录
	user, err := auth.Attempt(request.LoginID, request.Password)

	if err != nil {
		// 失败，显示错误提示
		// c.AbortWithStatusJSON
		response.Unauthorized(c, "账号不存在或密码错误")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

func (lc LoginController) RefreshToken(c *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
