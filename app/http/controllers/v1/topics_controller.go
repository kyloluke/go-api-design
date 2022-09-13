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

func (ctrl *TopicsController) Update(c *gin.Context) {
	// 1. ターゲット話題は存在するか判定する
	topicModel := topic.Get(c.Param("id"))

	if topicModel.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "該当データが見つかりませんでした",
		})
	}

	// 2 バリデーション
	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel.Title = request.Title
	topicModel.Body = request.Body
	topicModel.CategoryID = request.CategoryID

	rowsAffected := topicModel.Save()
	if rowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "内部错误，请稍后再试",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    topicModel,
	})

}
