package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type MaterialRequest struct {
	CreatorId       uint64 `json:"creator_id,omitempty"`
	MaterialGroupId uint64 `json:"material_group_id,omitempty"`
	DepartmentId    uint64 `json:"department_id,omitempty"`
	Type            uint64 `valid:"type" json:"type,omitempty"`
	Url             string `json:"url,omitempty"`
	Title           string `valid:"title" json:"title,omitempty"`
	Content         string `json:"content,omitempty"`
}

func MaterialSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"type":  []string{"in:0,1"},
		"title": []string{"min:2", "max:30"},
	}
	messages := govalidator.MapData{
		"type": []string{
			"in:只能为0或1",
		},
		"title": []string{
			"min:最小长度为2",
			"max:最大长度为30",
		},
	}
	return validate(data, rules, messages)
}
