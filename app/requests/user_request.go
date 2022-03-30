package requests

import (
	"gohub/app/requests/validators"
	"gohub/pkg/auth"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type UserCreateRequest struct {
	DepartmentId uint64 `json:"department_id" valid:"department_id"`
	GuardName    string `json:"guard_name" valid:"guard_name" `
	Username     string `json:"username" valid:"username"`
	Name         string `json:"name" valid:"name"`
	Gender       uint64 `json:"gender" valid:"gender"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Avatar       string `json:"avatar"`
	Status       uint64 `json:"status;omitempty"`
}

func UserCreate(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"department_id": []string{"required", "exists:departments,id"},
		"guard_name":    []string{"required"},
		"username":      []string{"required", "alpha_num", "between:3,20", "not_exists:users,username"},
		"name":          []string{"required", "min:2"},
		"gender":        []string{"required", "in:1,2"},
	}

	messages := govalidator.MapData{
		"department_id": []string{
			"required: 部门为必填项，参数名称 department_id",
			"exists: 部门信息不存在",
		},
		"guard_name": []string{
			"required: 项目为必填项，参数名称 guard_name",
		},
		"username": []string{
			"required: 用户名为必填项",
			"alpha_num: 用户名格式错误，只允许数字和英文",
			"between: 用户名长度需在 3~20 之间",
			"not_exists: 用户名已被占用",
		},
		"name": []string{
			"required: 姓名为必填项",
			"min: 姓名长度需大于 2",
		},
		"gender": []string{
			"required: 性别为必填项",
			"in: 性别只能是男或女",
		},
	}

	errs := validate(data, rules, messages)

	return errs
}

type UserUpdateProfileRequest struct {
	Name  string `valid:"name" json:"name"`
	Email string `valid:"email" json:"email"`
	Phone string `valid:"phone" json:"phone"`
}

func UserUpdateProfile(data interface{}, c *gin.Context) map[string][]string {

	// 查询用户名重复时，过滤掉当前用户 ID
	uid := auth.CurrentUID(c)
	rules := govalidator.MapData{
		"name": []string{"required", "alpha_num", "between:3,20", "not_exists:users,name," + uid},
		"email": []string{
			"required", "min:4",
			"max:30",
			"email",
		},
		"phone": []string{
			"required",
			"digits:11",
		},
	}

	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
			"not_exists:用户名已被占用",
		},
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	return validate(data, rules, messages)
}

type UserUpdateEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func UserUpdateEmail(data interface{}, c *gin.Context) map[string][]string {

	currentUser := auth.CurrentUser(c)
	rules := govalidator.MapData{
		"email": []string{
			"required", "min:4",
			"max:30",
			"email",
			"not_exists:users,email," + currentUser.GetStringID(),
			"not_in:" + currentUser.Email,
		},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
			"not_in:新的 Email 与老 Email 一致",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*UserUpdateEmailRequest)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs
}

type UserUpdatePhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func UserUpdatePhone(data interface{}, c *gin.Context) map[string][]string {

	currentUser := auth.CurrentUser(c)

	rules := govalidator.MapData{
		"phone": []string{
			"required",
			"digits:11",
			"not_exists:users,phone," + currentUser.GetStringID(),
			"not_in:" + currentUser.Phone,
		},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
			"not_exists:手机号已被占用",
			"not_in:新的手机与老手机号一致",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*UserUpdatePhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

type UserUpdatePasswordRequest struct {
	Password           string `valid:"password" json:"password,omitempty"`
	NewPassword        string `valid:"new_password" json:"new_password,omitempty"`
	NewPasswordConfirm string `valid:"new_password_confirm" json:"new_password_confirm,omitempty"`
}

func UserUpdatePassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"password":             []string{"required", "min:6"},
		"new_password":         []string{"required", "min:6"},
		"new_password_confirm": []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"new_password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"new_password_confirm": []string{
			"required:确认密码框为必填项",
			"min:确认密码长度需大于 6",
		},
	}

	// 确保 comfirm 密码正确
	errs := validate(data, rules, messages)
	_data := data.(*UserUpdatePasswordRequest)
	errs = validators.ValidatePasswordConfirm(_data.NewPassword, _data.NewPasswordConfirm, errs)

	return errs
}

type UserUpdateAvatarRequest struct {
	Avatar *multipart.FileHeader `valid:"avatar" form:"avatar"`
}

func UserUpdateAvatar(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// size 的单位为 bytes
		// - 1024 bytes 为 1kb
		// - 1048576 bytes 为 1mb
		// - 20971520 bytes 为 20mb
		"file:avatar": []string{"ext:png,jpg,jpeg", "size:20971520", "required"},
	}
	messages := govalidator.MapData{
		"file:avatar": []string{
			"ext:ext头像只能上传 png, jpg, jpeg 任意一种的图片",
			"size:头像文件最大不能超过 20MB",
			"required:必须上传图片",
		},
	}

	return validateFile(c, data, rules, messages)
}
