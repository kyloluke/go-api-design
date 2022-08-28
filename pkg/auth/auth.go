// Package auth 授权相关逻辑
package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gohub/app/models/user"
	"gohub/pkg/logger"
)

// Attempt 登录
func Attempt(loginId string, password string) (user.User, error) {

	// 1. 查找账号是否存在
	userModel := user.GetByMulti(loginId)
	if userModel.ID == 0 {
		return userModel, errors.New("账号不存在")
	}
	// 2. 验证密码是否正确
	ok := userModel.ComparePassword(password)
	if !ok {
		return userModel, errors.New("密码错误")
	}
	return userModel, nil
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
