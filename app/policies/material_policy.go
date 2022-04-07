package policies

import (
	"gohub/app/models/material"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyMaterial(c *gin.Context, materialModel material.Material) bool {
	return auth.CurrentUser(c).Username == "admin"
}

// func CanViewMaterial(c *gin.Context, materialModel material.Material) bool {}
// func CanCreateMaterial(c *gin.Context, materialModel material.Material) bool {}
// func CanUpdateMaterial(c *gin.Context, materialModel material.Material) bool {}
// func CanDeleteMaterial(c *gin.Context, materialModel material.Material) bool {}
