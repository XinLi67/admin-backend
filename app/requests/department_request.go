package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type DepartmentRequest struct {
	ParentId    uint64 `json:"parent_id" valid:"parent_id"`
	Name        string `json:"name" valid:"name"`
	Phone       string `json:"phone"`
	LinkMan     string `json:"link_man"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

func DepartmentSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"parent_id": []string{"required", "exists:departments,id"},
		"name":      []string{"required", "min_cn:2", "max_cn:60", "not_exists:departments,name"},
	}
	messages := govalidator.MapData{
		"parent_id": []string{
			"required: 部门信息为必输项",
			"exists: 查询不到部门信息",
		},
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 60 个字",
			"not_exists:名称已存在",
		},
	}

	return validate(data, rules, messages)
}
