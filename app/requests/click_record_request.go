package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ClickRecordRequest struct {
	AdvertisingId uint64 `json:"advertising_id,omitempty"`
	CustomerId    uint64 `json:"customer_id,omitempty"`
	BrowsingTime  uint64 `valid:"browsing_time" json:"browsing_time,omitempty"`
	StartTime     string `json:"start_time,omitempty"`
	EndTime       string `json:"end_time,omitempty"`
}

func ClickRecordSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"browsing_time": []string{"numeric"},
	}
	messages := govalidator.MapData{
		"browsing_time:": []string{
			"numeric:必须是数字",
		},
	}
	return validate(data, rules, messages)
}
