package v1

import (
	"gohub/app/http/assemblies"
	"gohub/app/models/advertising"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type AdvertisingsController struct {
	BaseAPIController
}

func (ctrl *AdvertisingsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	// data, pager := advertising.Paginate(c, 0)
	data, pager := advertising.Search(c, 0)
	advertisings := assemblies.AdvertisingAssemblyFromModelList(data, len(data))
	response.JSON(c, gin.H{
		"data":  advertisings,
		"pager": pager,
	})
}

func (ctrl *AdvertisingsController) Show(c *gin.Context) {
	advertisingModel := advertising.Get(c.Param("id"))
	if advertisingModel.ID == 0 {
		response.Abort404(c)
		return
	}
	advertisingAssembly := assemblies.AdvertisingAssemblyFromModel(advertisingModel)
	response.Data(c, advertisingAssembly)
}

func (ctrl *AdvertisingsController) Store(c *gin.Context) {

	request := requests.AdvertisingRequest{}
	if ok := requests.Validate(c, &request, requests.AdvertisingSave); !ok {
		return
	}

	advertisingModel := advertising.Advertising{
		AdvertisingPositionId: request.AdvertisingPositionId,
		CreatorId:             request.CreatorId,
		DepartmentId:          request.DepartmentId,
		Title:                 request.Title,
		Type:                  request.Type,
		RedirectTo:            request.RedirectTo,
		MaterialId:            request.MaterialId,
		MaterialType:          request.Materialtype,
		Size:                  request.Size,
		RedirectParams:        request.RedirectParams,
		Description:           request.Description,
		Status:                request.Status,
	}
	advertisingModel.Create()
	if advertisingModel.ID > 0 {
		response.Created(c, advertisingModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AdvertisingsController) Update(c *gin.Context) {

	advertisingModel := advertising.Get(c.Param("id"))
	if advertisingModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertising(c, advertisingModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AdvertisingRequest{}
	bindOk := requests.Validate(c, &request, requests.AdvertisingSave)
	if !bindOk {
		return
	}

	advertisingModel.AdvertisingPositionId = request.AdvertisingPositionId
	advertisingModel.CreatorId = request.CreatorId
	advertisingModel.DepartmentId = request.DepartmentId
	advertisingModel.Title = request.Title
	advertisingModel.Type = request.Type
	advertisingModel.RedirectTo = request.RedirectTo
	advertisingModel.MaterialId = request.MaterialId
	advertisingModel.MaterialType = request.Materialtype
	advertisingModel.Size = request.Size
	advertisingModel.RedirectParams = request.RedirectParams
	advertisingModel.Description = request.Description
	advertisingModel.Status = request.Status

	rowsAffected := advertisingModel.Save()
	if rowsAffected > 0 {
		response.Data(c, advertisingModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AdvertisingsController) Delete(c *gin.Context) {

	advertisingModel := advertising.Get(c.Param("id"))
	if advertisingModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertising(c, advertisingModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := advertisingModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *AdvertisingsController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	advertisingModel := advertising.Advertising{}
	if ok := policies.CanModifyAdvertising(c, advertisingModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := advertisingModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}
