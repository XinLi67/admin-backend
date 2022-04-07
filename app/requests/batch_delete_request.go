package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type BatchDeleteRequest struct {
	Ids []int `valid:"ids" json:"ids"`
}

func BatchDelete(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"ids": []string{"required"},
	}

	messages := govalidator.MapData{
		"ids": []string{
			"required:IDS为必填项,参数名称 ids",
		},
	}
	return validate(data, rules, messages)
}
