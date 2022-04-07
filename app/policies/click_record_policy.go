package policies

import (
	"gohub/app/models/click_record"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyClickRecord(c *gin.Context, clickRecordModel click_record.ClickRecord) bool {
	return auth.CurrentUser(c).Username == "admin"
}

// func CanViewClickRecord(c *gin.Context, clickRecordModel click_record.ClickRecord) bool {}
// func CanCreateClickRecord(c *gin.Context, clickRecordModel click_record.ClickRecord) bool {}
// func CanUpdateClickRecord(c *gin.Context, clickRecordModel click_record.ClickRecord) bool {}
// func CanDeleteClickRecord(c *gin.Context, clickRecordModel click_record.ClickRecord) bool {}
