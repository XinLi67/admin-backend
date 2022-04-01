package v1

import (
	"gohub/app/http/assemblies"
	"gohub/app/models/user"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/config"
	"gohub/pkg/file"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 0)

	users := assemblies.UserAssemblyFromModelList(data)

	response.JSON(c, gin.H{
		"data":  users,
		"pager": pager,
	})
}

// 查看用户详情
func (ctrl *UsersController) Show(c *gin.Context) {
	userModel := user.Get(c.Param("id"))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}

	userAssembly := assemblies.UserAssemblyFromModel(userModel)

	response.Data(c, userAssembly)
}

func (ctrl *UsersController) Store(c *gin.Context) {

	request := requests.UserCreateRequest{}
	if ok := requests.Validate(c, &request, requests.UserCreate); !ok {
		return
	}

	userModel := user.User{
		DepartmentId: request.DepartmentId,
		Username:     request.Username,
		Name:         request.Name,
		Gender:       request.Gender,
		Email:        request.Email,
		Phone:        request.Phone,
		Avatar:       request.Avatar,
		Status:       request.Status,
	}

	userModel.Create()
	if userModel.ID > 0 {
		response.Created(c, userModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) Update(c *gin.Context) {

	userModel := user.Get(c.Param("id"))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyUser(c, userModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.UserUpdateProfileRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateProfile); !ok {
		return
	}

	userModel.Email = request.Email
	userModel.Phone = request.Phone
	userModel.Name = request.Name
	userModel.Gender = request.Gender
	rowsAffected := userModel.Save()
	if rowsAffected > 0 {
		response.Data(c, userModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) Delete(c *gin.Context) {

	userModel := user.Get(c.Param("id"))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyUser(c, userModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := userModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *UsersController) UpdateProfile(c *gin.Context) {

	request := requests.UserUpdateProfileRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateProfile); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Name = request.Name
	currentUser.Email = request.Email
	currentUser.Phone = request.Phone
	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Data(c, currentUser)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdateEmail(c *gin.Context) {

	request := requests.UserUpdateEmailRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateEmail); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Email = request.Email
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		// 失败，显示错误提示
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdatePhone(c *gin.Context) {

	request := requests.UserUpdatePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePhone); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Phone = request.Phone
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdatePassword(c *gin.Context) {

	request := requests.UserUpdatePasswordRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePassword); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	// 验证原始密码是否正确
	_, err := auth.Attempt(currentUser.Name, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Unauthorized(c, "原密码不正确")
	} else {
		// 更新密码为新密码
		currentUser.Password = request.NewPassword
		currentUser.Save()

		response.Success(c)
	}
}

func (ctrl *UsersController) UpdateAvatar(c *gin.Context) {

	request := requests.UserUpdateAvatarRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateAvatar); !ok {
		return
	}

	avatar, err := file.SaveUploadAvatar(c, request.Avatar)
	if err != nil {
		response.Abort500(c, "上传头像失败，请稍后尝试~")
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Avatar = config.GetString("app.url") + avatar
	currentUser.Save()

	response.Data(c, currentUser)
}
