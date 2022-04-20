package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ChannelRequest struct {
	Name        string `json:"name" valid:"name"`
	Description string `json:"description"`
	Status      uint64 `json:"status"`
}

func ChannelSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:2", "max_cn:30", "not_exists:channels,name," + c.Param("id")},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 30 个字",
			"not_exists:名称已存在",
		},
	}
	return validate(data, rules, messages)
}
