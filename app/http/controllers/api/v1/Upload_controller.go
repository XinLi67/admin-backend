package v1

import (
	"gohub/app/requests"
	"gohub/pkg/config"
	"gohub/pkg/file"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type UploadController struct {
	BaseAPIController
}

func (ctrl *UploadController) Upload(c *gin.Context) {

	request := requests.UploadRequest{}
	if ok := requests.Validate(c, &request, requests.Upload); !ok {
		return
	}
	avatar, err := file.SaveUploadAvatar(c, request.File)
	if err != nil {
		response.Abort500(c, "上传失败，请稍后尝试~")
		return
	}

	url := config.GetString("app.url") + avatar
	// url := avatar
	response.Data(c, url)
}
