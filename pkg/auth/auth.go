// Package auth 授权相关逻辑
package auth

import (
	"errors"
	"gohub/app/models/user"
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
