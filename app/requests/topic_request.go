package requests

import (
	"github.com/thedevsaddam/govalidator"
)

type TopicRequest struct {
	Title      string `valid:"title" json:"title,omitempty"`
	Body       string `valid:"body" json:"body,omitempty"`
	CategoryID string `valid:"category_id" json:"category_id,omitempty"`
}

func TopicSave(data interface{}) map[string][]string {
	rules := govalidator.MapData{
		"title":       []string{"required", "min_cn:2", "max_cn:40"},
		"body":        []string{"required", "min_cn:2", "max_cn:5000"},
		"category_id": []string{"required", "exists:categories,id"},
	}

	messages := govalidator.MapData{
		"title":       []string{"required:话题标题为必填项", "min_cn:话题标题最少2个字符", "max_cn:话题标题最多40个字符"},
		"body":        []string{"required:话题内容为必填项", "min_cn:话题内容最少2个字符", "max_cn:话题内容最多5000字符"},
		"category_id": []string{"required:话题分类为必填项", "exists:话题分类不存在"},
	}
	//validate(data, rules, messages)
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	return govalidator.New(opts).ValidateStruct()
}
