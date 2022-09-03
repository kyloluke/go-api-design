package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/user"
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
	data, pager := user.Paginate(c, 10)

	//response.JSON(c, gin.H{"data": data, "pager": pager})
	c.JSON(http.StatusOK, gin.H{
		"message": true,
		"data":    gin.H{"data": data, "pager": pager},
	})
}
