// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 获取请求参数，并做表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist 检测邮箱是否已注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}
	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}

func (sc *SignupController) Signup(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignupRequest{}
	if ok := requests.Validate(c, &request, requests.Signup); !ok {
		return
	}

	// 2. 验证成功，创建数据
	userModel := user.User{
		DepartmentId: request.DepartmentId,
		GuardName:    request.GuardName,
		UserName:     request.UserName,
		Name:         request.Name,
		Gender:       request.Gender,
		Email:        request.Email,
		Phone:        request.Phone,
		Avatar:       request.Avatar,
		Status:       request.Status,
	}

	userModel.Create()

	if userModel.ID > 0 {
		response.Data(c, userModel)
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}
