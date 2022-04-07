package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AnnouncementRequest struct {
	AnnouncementNo         uint64 `valid:"announcement_no" json:"announcement_no,omitempty"`
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
		"announcement_no": []string{"required", "not_exists:announcements,announcement_no"},
		"title":           []string{"min:2", "max:60"},
		"type":            []string{"numeric_between:-1,3"},
		"redirect_to":     []string{"numeric_between:-1,2"},
	}
	messages := govalidator.MapData{
		"announcement_no": []string{
			"required:公告代码为必填项",
			"not_exists:公告代码已存在",
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
	}
	return validate(data, rules, messages)
}
