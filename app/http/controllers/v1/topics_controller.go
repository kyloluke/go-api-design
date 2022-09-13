package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/app/models/topic"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"net/http"
)

type TopicsController struct {
	BaseAPIController
}

func (ctrl *TopicsController) Show(c *gin.Context) {
	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "該当データは見つかりませんでした",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    topicModel,
	})
}

func (ctrl *TopicsController) Index(c *gin.Context) {

	// 分页参数绑定
	type paginateRequest struct {
		Sort    string `valid:"sort" form:"sort"`
		Order   string `valid:"order" form:"order"`
		PerPage string `valid:"per_page" form:"per_page"`
	}
	var request paginateRequest
	if err := c.ShouldBind(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "数据解析错误",
		})
		return
	}

	// 分页参数数据验证
	rules := govalidator.MapData{
		"sort":     []string{"in:id,created_at,updated_at"},
		"order":    []string{"in:asc,desc"},
		"per_page": []string{"numeric_between:2,100"},
	}
	messages := govalidator.MapData{
		"sort":     []string{"in:排序字段必须是 id，created_at,updated_at 其中一个"},
		"order":    []string{"in:排序规则必须是 asc，desc 其中一个"},
		"per_page": []string{"numeric_between:每页数量必须2-100之间"},
	}

	opts := govalidator.Options{
		Data:          &request,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	errors := govalidator.New(opts).ValidateStruct()

	if len(errors) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"errors": errors,
		})
		return
	}

	// 获取分页数据
	topics, pager := topic.Paginate(c, 10)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"data":  topics,
			"pager": pager,
		},
	})
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

	// 2 編集権限のチェック
	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		//response.Abort403()
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "无权操作",
		})
		return
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

func (ctrl *TopicsController) Delete(c *gin.Context) {
	topicModel := topic.Get(c.Param("id"))

	if topicModel.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "帖子没找到，请重新操作",
		})
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "无权限操作",
		})
		return
	}

	rowsAffected := topicModel.Delete()
	if rowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "内部错误，请稍后再试",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
