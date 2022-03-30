package v1

import (
	"gohub/app/models/channel"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type ChannelsController struct {
	BaseAPIController
}

func (ctrl *ChannelsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := channel.Paginate(c, 0)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *ChannelsController) Show(c *gin.Context) {
	channelModel := channel.Get(c.Param("id"))
	if channelModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, channelModel)
}

func (ctrl *ChannelsController) Store(c *gin.Context) {

	request := requests.ChannelRequest{}
	if ok := requests.Validate(c, &request, requests.ChannelSave); !ok {
		return
	}

	channelModel := channel.Channel{
		Name:        request.Name,
		GuardName:   request.GuardName,
		Description: request.Description,
	}
	channelModel.Create()
	if channelModel.ID > 0 {
		response.Created(c, channelModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *ChannelsController) Update(c *gin.Context) {

	channelModel := channel.Get(c.Param("id"))
	if channelModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyChannel(c, channelModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.ChannelRequest{}
	bindOk := requests.Validate(c, &request, requests.ChannelSave)
	if !bindOk {
		return
	}

	channelModel.Name = request.Name
	channelModel.GuardName = request.GuardName
	channelModel.Description = request.Description
	rowsAffected := channelModel.Save()
	if rowsAffected > 0 {
		response.Data(c, channelModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ChannelsController) Delete(c *gin.Context) {

	channelModel := channel.Get(c.Param("id"))
	if channelModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyChannel(c, channelModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := channelModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
