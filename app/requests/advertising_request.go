package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AdvertisingRequest struct {
	AdvertisingNo         uint64 `valid:"advertising_no" json:"advertising_no,omitempty"`
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
}

func AdvertisingSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"advertising_no": []string{"required", "not_exists:advertisings,advertising_no"},
		"title":          []string{"min:2", "max:60"},
		"type":           []string{"numeric_between:-1,3"},
		"redirect_to":    []string{"numeric_between:-1,2"},
		"material_type":  []string{"numeric_between:-1,2"},
	}
	messages := govalidator.MapData{
		"advertising_no": []string{
			"required:广告代码为必填项",
			"not_exists:广告代码已存在",
		},
		"title": []string{
			"min:最小长度为2",
			"max:最大长度为60",
		},
		"type": []string{
			"numeric_between:只能为0或1或2",
		},
		"redirect_to": []string{
			"numeric_between:只能为0或1",
		},
		"material_type": []string{
			"numeric_between:只能为0或1",
		},
	}
	return validate(data, rules, messages)
}
