package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ChannelRequest struct {
	Name        string `json:"name" valid:"name"`
	GuardName   string `json:"guard_name" valid:"guard_name"`
	Description string `json:"description"`
}

func ChannelSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":       []string{"required", "min_cn:2", "max_cn:30", "not_exists:channels,name"},
		"guard_name": []string{"min_cn:2", "max_cn:30"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 30 个字",
			"not_exists:名称已存在",
		},
		"guard_name": []string{
			"min_cn:描述长度需至少 2 个字",
			"max_cn:描述长度不能超过 30 个字",
		},
	}
	return validate(data, rules, messages)
}
