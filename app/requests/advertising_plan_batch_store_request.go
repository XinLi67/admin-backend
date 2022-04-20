package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/app/models/advertising_plan"
)

type AdvertigingBatchStoreRequest struct {
	Data []advertising_plan.AdvertisingPlan `valid:"data" json:"data"`
}

func AdvertigingBatchStore(data interface{}, c *gin.Context) map[string][]string {

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

