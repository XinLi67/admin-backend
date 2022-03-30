// Package policies 用户授权
package policies

import (
	"gohub/app/models/menu"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyMenu(c *gin.Context, _menu menu.Menu) bool {
	return auth.CurrentUser(c).Username == "admin"
}
