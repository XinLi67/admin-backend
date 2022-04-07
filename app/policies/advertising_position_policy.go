package policies

import (
	"gohub/app/models/advertising_position"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyAdvertisingPosition(c *gin.Context, advertisingPositionModel advertising_position.AdvertisingPosition) bool {
	return auth.CurrentUser(c).Username == "admin"
}

// func CanViewAdvertisingPosition(c *gin.Context, advertisingPositionModel advertising_position.AdvertisingPosition) bool {}
// func CanCreateAdvertisingPosition(c *gin.Context, advertisingPositionModel advertising_position.AdvertisingPosition) bool {}
// func CanUpdateAdvertisingPosition(c *gin.Context, advertisingPositionModel advertising_position.AdvertisingPosition) bool {}
// func CanDeleteAdvertisingPosition(c *gin.Context, advertisingPositionModel advertising_position.AdvertisingPosition) bool {}
