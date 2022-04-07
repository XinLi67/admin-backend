package v1

import (
	"gohub/app/http/assemblies"
	"gohub/app/models/announcement_position"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type AnnouncementPositionsController struct {
	BaseAPIController
}

func (ctrl *AnnouncementPositionsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := announcement_position.Paginate(c, 0)
	announcementPositions := assemblies.AnnouncementPositionAssemblyFromModelList(data, len(data))
	response.JSON(c, gin.H{
		"data":  announcementPositions,
		"pager": pager,
	})
}

func (ctrl *AnnouncementPositionsController) Show(c *gin.Context) {
	announcementPositionModel := announcement_position.Get(c.Param("id"))
	if announcementPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}
	announcementPositionAssembly := assemblies.AnnouncementPositionAssemblyFromModel(announcementPositionModel)
	response.Data(c, announcementPositionAssembly)
}

func (ctrl *AnnouncementPositionsController) Store(c *gin.Context) {

	request := requests.AnnouncementPositionRequest{}
	if ok := requests.Validate(c, &request, requests.AnnouncementPositionSave); !ok {
		return
	}

	announcementPositionModel := announcement_position.AnnouncementPosition{
		Name:        request.Name,
		ChannelId:   request.ChannelId,
		Code:        request.Code,
		Height:      request.Height,
		Weight:      request.Weight,
		Status:      request.Status,
		Description: request.Description,
	}
	announcementPositionModel.Create()
	if announcementPositionModel.ID > 0 {
		response.Created(c, announcementPositionModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AnnouncementPositionsController) Update(c *gin.Context) {

	announcementPositionModel := announcement_position.Get(c.Param("id"))
	if announcementPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAnnouncementPosition(c, announcementPositionModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AnnouncementPositionRequest{}
	bindOk := requests.Validate(c, &request, requests.AnnouncementPositionSave)
	if !bindOk {
		return
	}

	announcementPositionModel.Name = request.Name
	announcementPositionModel.ChannelId = request.ChannelId
	announcementPositionModel.Code = request.Code
	announcementPositionModel.Height = request.Height
	announcementPositionModel.Weight = request.Weight
	announcementPositionModel.Status = request.Status
	announcementPositionModel.Description = request.Description

	rowsAffected := announcementPositionModel.Save()
	if rowsAffected > 0 {
		response.Data(c, announcementPositionModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AnnouncementPositionsController) Delete(c *gin.Context) {

	announcementPositionModel := announcement_position.Get(c.Param("id"))
	if announcementPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAnnouncementPosition(c, announcementPositionModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := announcementPositionModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *AnnouncementPositionsController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	announcementPositionModel := announcement_position.AnnouncementPosition{}
	if ok := policies.CanModifyAnnouncementPosition(c, announcementPositionModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := announcementPositionModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}
