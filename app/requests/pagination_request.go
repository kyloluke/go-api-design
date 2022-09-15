package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PaginationRequest struct {
	Order   string `valid:"order" form:"order"`
	Sort    string `valid:"sort" form:"sort"`
	PerPage string `valid:"per_page" form:"per_page"`
}

func Pagination(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"order":    []string{"in:asc,desc"},
		"sort":     []string{"in:id,created_at,updated_at"},
		"per_page": []string{"numeric_between:2,100"},
	}
	messages := govalidator.MapData{
		"order":    []string{"in:排序规则仅支持asc，desc"},
		"sort":     []string{"in:排序字段仅支持 id, created_at, updated_at"},
		"per_page": []string{"numeric_between:每页条数的值介于2-100之间"},
	}
	return validate(data, rules, messages)
}
