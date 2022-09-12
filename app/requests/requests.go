package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/response"
)

// ValidatorFunc 用于控制器的回调函数用
// ValidatePhoneIsExist() ValidateEmailIsExist() 等函数名称

type ValidatorFunc func(interface{}) map[string][]string

func Validate(c *gin.Context, data interface{}, handler ValidatorFunc) bool {
	// 解析json，将数据绑定到 结构体字段中
	// todo ShouldBind 并不能通用于 url的 query 形式 和 json 形式
	if err := c.ShouldBindJSON(data); err != nil {
		//c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		//	"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
		//	"error":   err.Error(),
		//})
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")

		fmt.Println(err.Error())
		return false
	}

	// 调用表单验证
	errs := handler(data)

	if len(errs) > 0 {
		//c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		//	"message": "表单验证不通过，具体请查看errors",
		//	"errors":  errs,
		//})
		response.ValidationError(c, errs)
		return false
	}

	// 商法有错，请求会被直接返回，入国通过则
	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	return govalidator.New(opts).ValidateStruct()

}
