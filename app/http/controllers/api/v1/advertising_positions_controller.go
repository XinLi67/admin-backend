package v1

import (
	"gohub/app/http/assemblies"
	"gohub/app/models/advertising_position"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type AdvertisingPositionsController struct {
	BaseAPIController
}

func (ctrl *AdvertisingPositionsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := advertising_position.Paginate(c, 0)
	advertisingPositions := assemblies.AdvertisingPositionAssemblyFromModelList(data, len(data))
	response.JSON(c, gin.H{
		"data":  advertisingPositions,
		"pager": pager,
	})
}

func (ctrl *AdvertisingPositionsController) Show(c *gin.Context) {
	advertisingPositionModel := advertising_position.Get(c.Param("id"))
	if advertisingPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}
	advertisingPositionAssembly := assemblies.AdvertisingPositionAssemblyFromModel(advertisingPositionModel)
	response.Data(c, advertisingPositionAssembly)
}

func (ctrl *AdvertisingPositionsController) Store(c *gin.Context) {

	request := requests.AdvertisingPositionRequest{}
	if ok := requests.Validate(c, &request, requests.AdvertisingPositionSave); !ok {
		return
	}

	advertisingPositionModel := advertising_position.AdvertisingPosition{
		Name:        request.Name,
		ChannelId:   request.ChannelId,
		Code:        request.Code,
		Height:      request.Height,
		Weight:      request.Weight,
		Status:      request.Status,
		Description: request.Description,
	}
	advertisingPositionModel.Create()
	if advertisingPositionModel.ID > 0 {
		response.Created(c, advertisingPositionModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AdvertisingPositionsController) Update(c *gin.Context) {

	advertisingPositionModel := advertising_position.Get(c.Param("id"))
	if advertisingPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertisingPosition(c, advertisingPositionModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AdvertisingPositionRequest{}
	bindOk := requests.Validate(c, &request, requests.AdvertisingPositionSave)
	if !bindOk {
		return
	}

	advertisingPositionModel.Name = request.Name
	advertisingPositionModel.ChannelId = request.ChannelId
	advertisingPositionModel.Code = request.Code
	advertisingPositionModel.Height = request.Height
	advertisingPositionModel.Weight = request.Weight
	advertisingPositionModel.Status = request.Status
	advertisingPositionModel.Description = request.Description

	rowsAffected := advertisingPositionModel.Save()
	if rowsAffected > 0 {
		response.Data(c, advertisingPositionModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AdvertisingPositionsController) Delete(c *gin.Context) {

	advertisingPositionModel := advertising_position.Get(c.Param("id"))
	if advertisingPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertisingPosition(c, advertisingPositionModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := advertisingPositionModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *AdvertisingPositionsController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	advertisingPositionModel := advertising_position.AdvertisingPosition{}
	if ok := policies.CanModifyAdvertisingPosition(c, advertisingPositionModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := advertisingPositionModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}
