package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/topic"
	"gohub/app/requests"
	"gohub/pkg/auth"
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

	// 保存数据
	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.CurrentUID(c),
	}

	topicModel.Create()

	if topicModel.ID > 0 {
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"data":    topicModel,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": "话题更新失败，请稍后再试",
	})

}
