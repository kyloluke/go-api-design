// Package factories 存放工厂方法
package factories

import (
	"github.com/bxcodec/faker/v3"
	"gohub/app/models/user"
	"gohub/pkg/helpers"
)

func MakeUsers(times int) []user.User {
	var objs []user.User
	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		userModel := user.User{
			Name:     faker.ChineseName(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		}
		objs = append(objs, userModel)
	}

	return objs
}
