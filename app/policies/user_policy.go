// Package policies 用户授权
package policies

import (
	"gohub/app/models/user"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyUser(c *gin.Context, _user user.User) bool {
	return auth.CurrentUID(c) == _user.GetStringID() || auth.CurrentUser(c).Username == "admin"
}
