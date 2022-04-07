package policies

import (
	"gohub/app/models/announcement"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyAnnouncement(c *gin.Context, announcementModel announcement.Announcement) bool {
	return auth.CurrentUser(c).Username == "admin"
}

// func CanViewAnnouncement(c *gin.Context, announcementModel announcement.Announcement) bool {}
// func CanCreateAnnouncement(c *gin.Context, announcementModel announcement.Announcement) bool {}
// func CanUpdateAnnouncement(c *gin.Context, announcementModel announcement.Announcement) bool {}
// func CanDeleteAnnouncement(c *gin.Context, announcementModel announcement.Announcement) bool {}
