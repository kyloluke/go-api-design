package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/link"
	"net/http"
)

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	links := link.AllCached()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    links,
	})
}
