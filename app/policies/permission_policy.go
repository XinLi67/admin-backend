// Package policies 用户授权
package policies

import (
	"gohub/app/models/permission"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyPermission(c *gin.Context, _permission permission.Permission) bool {
	return auth.CurrentUser(c).Username == "admin"
}
