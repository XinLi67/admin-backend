package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/app/models/announcement_plan"
)

type AnnouncementBatchStoreRequest struct {
	Data []announcement_plan.AnnouncementPlan `valid:"data" json:"data"`
}

func AnnouncementBatchStore(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"data": []string{"required"},
	}

	messages := govalidator.MapData{
		"data": []string{
			"required:参数data为必填项",
		},
	}
	return validate(data, rules, messages)
}

