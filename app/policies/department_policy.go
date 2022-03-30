// Package policies 用户授权
package policies

import (
	"gohub/app/models/department"
	"gohub/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyDepartment(c *gin.Context, _department department.Department) bool {
	return auth.CurrentUser(c).Username == "admin"
}
