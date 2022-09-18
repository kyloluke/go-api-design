package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"net/http"
)

type UsersController struct {
	BaseAPIController
}

func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

func (ctrl *UsersController) Index(c *gin.Context) {

	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10)

	//response.JSON(c, gin.H{"data": data, "pager": pager})
	c.JSON(http.StatusOK, gin.H{
		"message": true,
		"data":    gin.H{"data": data, "pager": pager},
	})
}

func (ctrl *UsersController) Update(c *gin.Context) {
	request := requests.UserRequest{}
	if ok := requests.Validate(c, &request, requests.UserSave); !ok {
		return
	}

	userModel := auth.CurrentUser(c)
	userModel.Name = request.Name
	userModel.City = request.City
	userModel.Introduction = request.Introduction

	rowsAffected := userModel.Save()
	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "内部错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userModel,
	})
}

func (*UsersController) UpdateEmail(c *gin.Context) {
	request := requests.UserUpdateEmailRequest{}

	if ok := requests.Validate(c, &request, requests.UserUpdateEmailValidate); !ok {
		return
	}

	currentUser, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "未知错误，请稍后再试",
		})
		return
	}

	currentUser.Email = request.Email
	rowsAffected := currentUser.Save()

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "更新失败，请稍后再试",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

// UpdatePassword 修改密码
func (ctrl *UsersController) UpdatePassword(c *gin.Context) {
	request := requests.UpdatePasswordRequest{}

	if !requests.UpdatePasswordValidate(&request, c) {
		return
	}

	// 验证原密码是否正确
	currentUser, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("未找到登录用户"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "未知错误",
		})
		return
	}

	if !currentUser.ComparePassword(request.CurrentPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "原密码错误",
		})
		return
	}

	// 修改密码
	currentUser.Password = request.NewPassword
	rowsAffected := currentUser.Save()
	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "内部错误，请稍后再试",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
