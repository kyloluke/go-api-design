package auth

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"
)

type PasswordResetController struct {
}

// ResetByEmail 重置密码
func (prc *PasswordResetController) ResetByEmail(c *gin.Context) {
	request := requests.ResetByEmailRequest{}

	if ok := requests.ResetByEmail(c, &request); !ok {
		return
	}

	userModel := user.GetByEmail(request.Email)

	userModel.Password = request.Password
	userModel.Save()
	response.Success(c)
}
