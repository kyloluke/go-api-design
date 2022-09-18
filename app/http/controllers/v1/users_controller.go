package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/config"
	"gohub/pkg/file"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"mime/multipart"
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

func (ctrl *UsersController) UpdateAvatar(c *gin.Context) {
	type updateAvatarRequest struct {
		Avatar *multipart.FileHeader `valid:"avatar" form:"avatar"`
	}
	var request = updateAvatarRequest{}

	// 数据绑定
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请确认数据格式是否正确",
		})
		return
	}

	// 数据验证
	rules := govalidator.MapData{
		"file:avatar": []string{"required", "ext:png,jpg,jpeg", "size:20971520"},
	}

	messages := govalidator.MapData{
		"file:avatar": []string{
			"required:请上传文件",
			"ext:文件格式只支持 png,jpg,jpeg",
			"size:文件大小不超过20MB",
		},
	}

	opts := govalidator.Options{
		Request:       c.Request,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}
	errs := govalidator.New(opts).Validate()

	if len(errs) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "数据验证失败",
			"errors":  errs,
		})
		return
	}

	// 保存文件
	//fmt.Printf("valud: %#v\n", request.Avatar)
	avatar, err := file.SaveUploadAvatar(c, request.Avatar)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "头像上传失败",
		})
		return
	}
	currentUser := auth.CurrentUser(c)
	currentUser.Avatar = config.GetString("app.url") + avatar
	currentUser.Save()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    currentUser,
	})
}
