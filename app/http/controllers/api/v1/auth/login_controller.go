package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/jwt"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// LoginController 用户控制器
type LoginController struct {
	v1.BaseAPIController
}

// LoginByPassword 多种方法登录，支持手机号、email 和用户名
func (lc *LoginController) LoginByPassword(c *gin.Context) {
	// 1. 验证表单
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.Attempt(request.Username, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Unauthorized(c, err.Error())

	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name, request.GuardName)
		response.JSON(c, gin.H{
			"data": map[string]string{
				"token": token,
			},
		})
	}
}

func (lc *LoginController) Logout(c *gin.Context) {
	err := jwt.NewJWT().JoinBlackList(c)
	if err != nil {
		response.Abort500(c)
	}

	response.Success(c)
}

// RefreshToken 刷新 Access Token
func (lc *LoginController) RefreshToken(c *gin.Context) {

	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
