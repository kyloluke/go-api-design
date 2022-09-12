package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/requests"
	"net/http"
)

type TopicsController struct {
	BaseAPIController
}

func (ctrl *TopicsController) Store(c *gin.Context) {
	// 数据验证
	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	// 数据保存
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"message": "下一步创建model实现 话题创建",
	})
}
