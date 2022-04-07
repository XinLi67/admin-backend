package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AnnouncementPlanRequest struct {
	Name                   string `valid:"name" json:"name,omitempty"`
	CreatorId              uint64 `json:"creator_id,omitempty"`
	AnnouncementId         uint64 `json:"announcement_id,omitempty"`
	AnnouncementType       uint64 `valid:"announcement_type" json:"announcement_type,omitempty"`
	AnnouncementPositionId uint64 `json:"announcement_position_id,omitempty"`
	Order                  uint64 `json:"order,omitempty"`
	SchedulingDate         uint64 `valid:"scheduling_date" json:"scheduling_date,omitempty"`
	SchedulingTime         uint64 `valid:"scheduling_time" json:"scheduling_time,omitempty"`
	StartDate              string `json:"start_date,omitempty"`
	EndDate                string `json:"end_date,omitempty"`
	StartTime              string `json:"start_time,omitempty"`
	Endime                 string `json:"end_time,omitempty"`
	AuditStatus            uint64 `json:"audit_status,omitempty"`
	PresentStatus          uint64 `json:"present_status,omitempty"`
}

func AnnouncementPlanSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":              []string{"required", "min_cn:2", "max_cn:30", "not_exists:announcement_plans,name"},
		"announcement_type": []string{"numeric_between:-1,3"},
		"SchedulingDate":    []string{"numeric_between:-1,2"},
		"SchedulingTime":    []string{"numeric_between:-1,2"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 30 个字",
			"not_exists:名称已存在",
		},
		"announcement_type": []string{
			"numeric_between:只能为0或1或2",
		},
		"SchedulingDate": []string{
			"numeric_between:只能为0或1",
		},
		"SchedulingTime": []string{
			"numeric_between:只能为0或1",
		},
	}
	return validate(data, rules, messages)
}
