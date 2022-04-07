package policies

import (
	"gohub/app/models/announcement_plan"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyAnnouncementPlan(c *gin.Context, announcementPlanModel announcement_plan.AnnouncementPlan) bool {
	return auth.CurrentUser(c).Username == "admin"
}

// func CanViewAnnouncementPlan(c *gin.Context, announcementPlanModel announcement_plan.AnnouncementPlan) bool {}
// func CanCreateAnnouncementPlan(c *gin.Context, announcementPlanModel announcement_plan.AnnouncementPlan) bool {}
// func CanUpdateAnnouncementPlan(c *gin.Context, announcementPlanModel announcement_plan.AnnouncementPlan) bool {}
// func CanDeleteAnnouncementPlan(c *gin.Context, announcementPlanModel announcement_plan.AnnouncementPlan) bool {}
