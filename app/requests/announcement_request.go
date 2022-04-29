package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	// "time"
)

type AnnouncementRequest struct {
	AnnouncementPositionId uint64 `json:"announcement_position_id,omitempty"`
	UserId                 uint64 `json:"user_id,omitempty"`
	Title                  string `valid:"title" json:"title,omitempty"`
	Type                   uint64 `valid:"type" json:"type,omitempty"`
	RedirectTo             uint64 `valid:"redirect_to" json:"redirect_to,omitempty"`
	RedirectParams         string `json:"redirect_params,omitempty"`
	Content                string `json:"content,omitempty"`
	Status                 uint64 `json:"status,omitempty"`
	AuditReason            string `json:"audit_reason,omitempty"`
	SchedulingType         uint64 `json:"scheduling_type,omitempty"`
	StartDate              string `json:"start_date,omitempty"`
	EndDate                string `json:"end_date,omitempty"`
}

func AnnouncementSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":       []string{"min:2", "max:60", "not_exists:announcements,title," + c.Param("id")},
		"type":        []string{"in:0,1,2"},
		"redirect_to": []string{"in:0,1"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"min:最小长度为2",
			"max:最大长度为60",
			"not_exists:公告名称已存在",
		},
		"type": []string{
			"in:只能为0或1或2",
		},
		"redirect_to": []string{
			"in:只能为0或1",
		},
	}
	return validate(data, rules, messages)
}
