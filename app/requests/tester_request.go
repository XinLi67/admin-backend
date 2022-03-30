// Package requests 处理请求数据和表单验证
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type BatchDeleteRequest struct {
	Ids []int `json:"ids" valid:"ids"`
}

func BatchDelete(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"ids": []string{"required"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"ids": []string{
			"required:IDS为必填项，参数名称 ids",
		},
	}
	return validate(data, rules, messages)
}
