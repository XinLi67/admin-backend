package v1

import (
	"gohub/app/http/assemblies"
	"gohub/app/models/material"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type MaterialsController struct {
	BaseAPIController
}

func (ctrl *MaterialsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := material.Paginate(c, 0)
	materials := assemblies.MaterialAssemblyFromModelList(data)
	response.JSON(c, gin.H{
		"data":  materials,
		"pager": pager,
	})
}

func (ctrl *MaterialsController) Show(c *gin.Context) {
	materialModel := material.Get(c.Param("id"))
	if materialModel.ID == 0 {
		response.Abort404(c)
		return
	}
	materialAssembly := assemblies.MaterialAssemblyFromModel(materialModel)
	response.Data(c, materialAssembly)
}

func (ctrl *MaterialsController) Store(c *gin.Context) {

	request := requests.MaterialRequest{}
	if ok := requests.Validate(c, &request, requests.MaterialSave); !ok {
		return
	}

	materialModel := material.Material{
		CreatorId:       request.CreatorId,
		MaterialGroupId: request.MaterialGroupId,
		DepartmentId:    request.DepartmentId,
		Type:            request.Type,
		Url:             request.Url,
		Title:           request.Title,
		Content:         request.Content,
	}
	materialModel.Create()
	if materialModel.ID > 0 {
		response.Created(c, materialModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *MaterialsController) Update(c *gin.Context) {

	materialModel := material.Get(c.Param("id"))
	if materialModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyMaterial(c, materialModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.MaterialRequest{}
	bindOk := requests.Validate(c, &request, requests.MaterialSave)
	if !bindOk {
		return
	}

	materialModel.CreatorId = request.CreatorId
	materialModel.MaterialGroupId = request.MaterialGroupId
	materialModel.DepartmentId = request.DepartmentId
	materialModel.Type = request.Type
	materialModel.Url = request.Url
	materialModel.Title = request.Title
	materialModel.Content = request.Content

	rowsAffected := materialModel.Save()
	if rowsAffected > 0 {
		response.Data(c, materialModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *MaterialsController) Delete(c *gin.Context) {

	materialModel := material.Get(c.Param("id"))
	if materialModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyMaterial(c, materialModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := materialModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *MaterialsController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	materialModel := material.Material{}
	if ok := policies.CanModifyMaterial(c, materialModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := materialModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}
