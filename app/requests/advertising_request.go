package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AdvertisingRequest struct {
	AdvertisingPositionId uint64 `json:"advertising_position_id,omitempty"`
	CreatorId             uint64 `json:"creator_id,omitempty"`
	DepartmentId          uint64 `json:"department_id,omitempty"`
	Title                 string `valid:"title" json:"title,omitempty"`
	Type                  uint64 `valid:"type" json:"type,omitempty"`
	RedirectTo            uint64 `valid:"redirect_to" json:"redirect_to,omitempty"`
	MaterialId            uint64 `json:"material_id,omitempty"`
	Materialtype          uint64 `valid:"material_type" json:"material_type,omitempty"`
	Size                  string `json:"size,omitempty"`
	RedirectParams        string `json:"redirect_params,omitempty"`
	Description           string `json:"description,omitempty"`
	Status                uint64 `json:"status,omitempty"`
	AuditReason           string `json:"audit_reason"`
	PushContent           string `json:"push_content,omitempty"`
	PushTitle             string `json:"push_title,omitempty"`
	AdvertisingCreativity string `json:"advertising_creativity,omitempty"`
}

func AdvertisingSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":         []string{"min:2", "max:60", "not_exists:advertisings,title," + c.Param("id")},
		"type":          []string{"in:0,1,2,3"},
		"redirect_to":   []string{"in:0,1"},
		"material_type": []string{"in:0,1"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"min:最小长度为2",
			"max:最大长度为60",
			"not_exists:广告名称已存在",
		},
		"type": []string{
			"in:只能为0或1或2或3",
		},
		"redirect_to": []string{
			"in:只能为0或1",
		},
		"material_type": []string{
			"in:只能为0或1",
		},
	}
	return validate(data, rules, messages)
}
