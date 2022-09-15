package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/auth"
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
