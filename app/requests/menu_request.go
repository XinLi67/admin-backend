package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type MenuRequest struct {
	ParentId       uint64 `json:"parent_id" valid:"parent_id"`
	Name           string `json:"name" valid:"name"`
	Icon           string `json:"icon" `
	Uri            string `json:"uri" `
	IsLink         bool   `json:"is_link" `
	PermissionName string `json:"permission_name" valid:"permission_name"`
	GuardName      string `json:"guard_name" valid:"guard_name"`
	Sequence       uint64 `json:"sequence" valid:"sequence"`
}

func MenuSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"parent_id":  []string{"exists_or_zero:menus,id"},
		"name":       []string{"required", "min_cn:2", "max_cn:8", "not_exists:menus,name"},
		"guard_name": []string{"required", "min:2", "max:30"},
	}
	messages := govalidator.MapData{
		"parent_id": []string{
			"exists_or_zero:上级名称不存在",
		},
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 8 个字",
			"not_exists:名称已存在",
		},
		"guard_name": []string{
			"required:名称为必填项",
			"min:名称长度需至少 2 个字",
			"max:名称长度不能超过 30 个字",
		},
	}
	return validate(data, rules, messages)
}

func MenuUpdate(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"parent_id":  []string{"exists_or_zero:menus,id"},
		"name":       []string{"required", "min_cn:2", "max_cn:8"},
		"guard_name": []string{"required", "min:2", "max:30"},
	}
	messages := govalidator.MapData{
		"parent_id": []string{
			"exists_or_zero:上级名称不存在",
		},
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 8 个字",
		},
		"guard_name": []string{
			"required:名称为必填项",
			"min:名称长度需至少 2 个字",
			"max:名称长度不能超过 30 个字",
		},
	}
	return validate(data, rules, messages)
}
