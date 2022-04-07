package policies

import (
	"gohub/app/models/advertising"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyAdvertising(c *gin.Context, advertisingModel advertising.Advertising) bool {
	return auth.CurrentUser(c).Username == "admin"
}

// func CanViewAdvertising(c *gin.Context, advertisingModel advertising.Advertising) bool {}
// func CanCreateAdvertising(c *gin.Context, advertisingModel advertising.Advertising) bool {}
// func CanUpdateAdvertising(c *gin.Context, advertisingModel advertising.Advertising) bool {}
// func CanDeleteAdvertising(c *gin.Context, advertisingModel advertising.Advertising) bool {}
