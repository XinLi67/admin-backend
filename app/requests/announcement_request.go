package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AnnouncementRequest struct {
	AnnouncementPositionId uint64 `json:"announcement_position_id,omitempty"`
	CreatorId              uint64 `json:"creator_id,omitempty"`
	DepartmentId           uint64 `json:"department_id,omitempty"`
	Title                  string `valid:"title" json:"title,omitempty"`
	LongTitle              string `json:"long_title,omitempty"`
	Type                   uint64 `valid:"type" json:"type,omitempty"`
	Banner                 string `json:"banner,omitempty"`
	RedirectTo             uint64 `valid:"redirect_to" json:"redirect_to,omitempty"`
	RedirectParams         string `json:"redirect_params,omitempty"`
	Content                string `json:"content,omitempty"`
	Status                 uint64 `json:"status,omitempty"`
}

func AnnouncementSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":       []string{"min:2", "max:60", "not_exists:announcements,title," + c.Param("id")},
		"type":        []string{"in:0,1"},
		"redirect_to": []string{"in:0,1"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"min:最小长度为2",
			"max:最大长度为60",
			"not_exists:公告名称已存在",
		},
		"type": []string{
			"in:只能为0或1",
		},
		"redirect_to": []string{
			"in:只能为0或1",
		},
	}
	return validate(data, rules, messages)
}
