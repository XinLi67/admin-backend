package policies

import (
	"gohub/app/models/announcement_position"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyAnnouncementPosition(c *gin.Context, announcementPositionModel announcement_position.AnnouncementPosition) bool {
	return auth.CurrentUser(c).Username == "admin"
}

// func CanViewAnnouncementPosition(c *gin.Context, announcementPositionModel announcement_position.AnnouncementPosition) bool {}
// func CanCreateAnnouncementPosition(c *gin.Context, announcementPositionModel announcement_position.AnnouncementPosition) bool {}
// func CanUpdateAnnouncementPosition(c *gin.Context, announcementPositionModel announcement_position.AnnouncementPosition) bool {}
// func CanDeleteAnnouncementPosition(c *gin.Context, announcementPositionModel announcement_position.AnnouncementPosition) bool {}
