// Package policies 用户授权
package policies

import (
	"gohub/app/models/role"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyRole(c *gin.Context, _role role.Role) bool {
	return auth.CurrentUser(c).Username == "admin"
}
