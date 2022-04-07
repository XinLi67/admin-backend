package policies

import (
	"gohub/app/models/material_group"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyMaterialGroup(c *gin.Context, materialGroupModel material_group.MaterialGroup) bool {
	return auth.CurrentUser(c).Username == "admin"
}

// func CanViewMaterialGroup(c *gin.Context, materialGroupModel material_group.MaterialGroup) bool {}
// func CanCreateMaterialGroup(c *gin.Context, materialGroupModel material_group.MaterialGroup) bool {}
// func CanUpdateMaterialGroup(c *gin.Context, materialGroupModel material_group.MaterialGroup) bool {}
// func CanDeleteMaterialGroup(c *gin.Context, materialGroupModel material_group.MaterialGroup) bool {}
