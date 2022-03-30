package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type PermissionRequest struct {
	PermissionGroupId uint64 `json:"pg_id" valid:"permission_group_id"`
	Name              string `json:"name" valid:"name"`
	Icon              string `json:"icon"`
	GuardName         string `json:"gurad_name"`
	DisplayName       string `json:"display_name" valid:"display_name"`
	Description       string `json:"description"`
	Sequence          uint64 `json:"sequence"`
}

func PermissionSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"permission_group_id": []string{"required", "exists:permission_groups,id"},
		"name":                []string{"required", "min_cn:2", "max_cn:8"},
		"display_name":        []string{"required", "min_cn:2", "max_cn:30"},
	}
	messages := govalidator.MapData{
		"permission_group_id": []string{
			"required:权限组为必填项",
			"not_exists:权限组不存在",
		},
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 8 个字",
		},
		"display_name": []string{
			"required:显示名称为必填项",
			"min_cn:描述长度需至少 2 个字",
			"max_cn:描述长度不能超过 30 个字",
		},
	}
	return validate(data, rules, messages)
}
