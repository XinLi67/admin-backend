package v1

import (
	"gohub/app/models/channel"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type TestersController struct {
	BaseAPIController
}

func (ctrl *TestersController) BatchDelete(c *gin.Context) {
	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	channelModel := &channel.Channel{}
	rowsAffected := channelModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}
