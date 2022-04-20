package policies

import (
	// "fmt"
	"gohub/app/models/material"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

//权限查询
func CanModifyMaterial(c *gin.Context, materialModel material.Material) bool {
	return auth.CurrentUser(c).Username == "admin"
	// CreatorId := fmt.Sprintf("%d", materialModel.CreatorId)
	// return auth.CurrentUID(c) == CreatorId
}

// func CanViewMaterial(c *gin.Context, materialModel material.Material) bool {}
// func CanCreateMaterial(c *gin.Context, materialModel material.Material) bool {}
// func CanUpdateMaterial(c *gin.Context, materialModel material.Material) bool {}
// func CanDeleteMaterial(c *gin.Context, materialModel material.Material) bool {}
