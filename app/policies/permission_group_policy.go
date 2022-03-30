// Package policies 用户授权
package policies

import (
	"gohub/app/models/permission_group"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyPermissionGroup(c *gin.Context, _permissionGroup permission_group.PermissionGroup) bool {
	return auth.CurrentUser(c).Username == "admin"
}
