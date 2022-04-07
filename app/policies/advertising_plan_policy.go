package policies

import (
	"gohub/app/models/advertising_plan"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyAdvertisingPlan(c *gin.Context, advertisingPlanModel advertising_plan.AdvertisingPlan) bool {
	return auth.CurrentUser(c).Username == "admin"
}

// func CanViewAdvertisingPlan(c *gin.Context, advertisingPlanModel advertising_plan.AdvertisingPlan) bool {}
// func CanCreateAdvertisingPlan(c *gin.Context, advertisingPlanModel advertising_plan.AdvertisingPlan) bool {}
// func CanUpdateAdvertisingPlan(c *gin.Context, advertisingPlanModel advertising_plan.AdvertisingPlan) bool {}
// func CanDeleteAdvertisingPlan(c *gin.Context, advertisingPlanModel advertising_plan.AdvertisingPlan) bool {}
