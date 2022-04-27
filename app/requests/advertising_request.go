package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/app/models"
	"mime/multipart"
)

type AdvertisingRequest struct {
	models.BaseModel

	AdvertisingNo         uint64 `json:"advertising_no"`
	AdvertisingPositionId uint64 `json:"advertising_position_id,omitempty"`
	CreatorId             uint64 `json:"creator_id,omitempty"`
	DepartmentId          uint64 `json:"department_id,omitempty"`
	Title                 string `valid:"title" json:"title,omitempty"`
	Type                  uint64 `valid:"type" json:"type,omitempty"`
	RedirectTo            uint64 `valid:"redirect_to" json:"redirect_to,omitempty"`
	MaterialId            uint64 `json:"material_id,omitempty"`
	MaterialType          uint64 `valid:"material_type" json:"material_type,omitempty"`
	Size                  string `json:"size,omitempty"`
	RedirectParams        string `json:"redirect_params,omitempty"`
	Description           string `json:"description,omitempty"`
	Status                uint64 `json:"status,omitempty"`
	AuditReason           string `json:"audit_reason,omitempty"`
	PushContent           string `json:"push_content,omitempty"`
	PushTitle             string `json:"push_title,omitempty"`
	AdvertisingCreativity string `json:"advertising_creativity,omitempty"`
	StartTime             string `json:"start_time,omitempty"`
	EndTime               string `json:"end_time,omitempty"`
	SchedulingTime        uint64 `json:"scheduling_time,omitempty"`
	Url                   *multipart.FileHeader `json:"url  form:"url"`
	Url2                  *multipart.FileHeader `json:"url2 form:"url2"`
	Url3                  *multipart.FileHeader `json:"url3 form:"url3"`
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
