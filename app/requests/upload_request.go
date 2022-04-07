package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"mime/multipart"
)


type UploadRequest struct {
	File *multipart.FileHeader `valid:"file" form:"file"`
}


func Upload(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// size 的单位为 bytes
		// - 1024 bytes 为 1kb
		// - 1048576 bytes 为 1mb
		// - 20971520 bytes 为 20mb
		"file:file": []string{"ext:png,jpg,jpeg,mp4", "size:20971520", "required"},
	}
	messages := govalidator.MapData{
		"file:file": []string{
			"ext:ext文件只能上传 png, jpg, jpeg ,,mp4任意一种的文件",
			"size:头像文件最大不能超过 20MB",
			"required:必须上传文件",
		},
	}

	return validateFile(c, data, rules, messages)
}
