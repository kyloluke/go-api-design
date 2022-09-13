package policies

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/topic"
)

func CanModifyTopic(c *gin.Context, topic topic.Topic) bool {
	return c.GetString("current_user_id") == topic.UserID
}
