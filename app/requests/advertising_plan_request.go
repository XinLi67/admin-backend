package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/app/models"
)

type AdvertisingPlanRequest struct {

	models.BaseModel
	Name                  string `valid:"name" json:"name,omitempty"`
	CreatorId             uint64 `json:"creator_id,omitempty"`
	AdvertisingId         uint64 `json:"advertising_id,omitempty"`
	AdvertisingType       uint64 `valid:"advertising_type" json:"advertising_type,omitempty"`
	AdvertisingPositionId uint64 `json:"advertising_position_id,omitempty"`
	Order                 uint64 `json:"order,omitempty"`
	SchedulingDate        uint64 `valid:"scheduling_date" json:"scheduling_date,omitempty"`
	SchedulingTime        uint64 `valid:"scheduling_time" json:"scheduling_time,omitempty"`
	StartDate             string `json:"start_date,omitempty"`
	EndDate               string `json:"end_date,omitempty"`
	StartTime             string `json:"start_time,omitempty"`
	EndTime               string `json:"end_time,omitempty"`
	AuditStatus           uint64 `json:"audit_status,omitempty"`
	PresentStatus         uint64 `json:"present_status,omitempty"`
}

func AdvertisingPlanSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":             []string{"required", "min_cn:2", "max_cn:30", "not_exists:advertising_plans,name," + c.Param("id")},
		"advertising_type": []string{"in:0,1,2,3"},
		"SchedulingDate":   []string{"in:0,1"},
		"SchedulingTime":   []string{"in:0,1"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 30 个字",
			"not_exists:名称已存在",
		},
		"advertising_type": []string{
			"in:只能为0或1或2或3",
		},
		"SchedulingDate": []string{
			"in:只能为0或1",
		},
		"SchedulingTime": []string{
			"in:只能为0或1",
		},
	}
	return validate(data, rules, messages)
}
