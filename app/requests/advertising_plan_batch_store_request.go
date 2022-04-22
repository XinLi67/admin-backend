package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AdvertigingBatchStoreRequest struct {
	AdvertisingPlan AdvertisingPlanRequest `valid:"advertising_plan" json:"advertising_plan"`
	Advertisings []AdvertisingRequest `valid:"advertisings" json:"advertisings"`

}

func AdvertigingBatchStore(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"advertising_plan": []string{},
	}

	messages := govalidator.MapData{
		"advertising_plan": []string{
			"required:参数advertising_plan为必填项",
		},
	}
	return validate(data, rules, messages)
}

