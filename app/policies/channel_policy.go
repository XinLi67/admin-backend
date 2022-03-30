// Package policies 用户授权
package policies

import (
	"gohub/app/models/channel"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyChannel(c *gin.Context, _channel channel.Channel) bool {
	return auth.CurrentUser(c).Username == "admin"
}
