package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AdvertisingPositionRequest struct {
	Name        string `valid:"name" json:"name,omitempty"`
	ChannelId   uint64 `json:"channel_id,omitempty"`
	Code        string `json:"code,omitempty"`
	Height      uint64 `valid:"height" json:"height,omitempty"`
	Weight      uint64 `valid:"weight" json:"weight,omitempty"`
	Status      uint64 `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
}

func AdvertisingPositionSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":   []string{"required", "min_cn:2", "max_cn:30", "not_exists:advertising_positions,name," + c.Param("id")},
		"height": []string{"numeric"},
		"weight": []string{"numeric"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 30 个字",
			"not_exists:名称已存在",
		},
		"height:": []string{
			"numeric:必须是数字",
		},
		"weight:": []string{
			"numeric:必须是数字",
		},
	}
	return validate(data, rules, messages)
}
