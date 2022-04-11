package v1

import (
	"gohub/app/http/assemblies"
	"gohub/app/models/material_group"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type MaterialGroupsController struct {
	BaseAPIController
}

func (ctrl *MaterialGroupsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := material_group.Search(c, 0)
	materialGroups := assemblies.MaterialGroupAssemblyFromModelList(data)
	response.JSON(c, gin.H{
		"data":  materialGroups,
		"pager": pager,
	})
}


func (ctrl *MaterialGroupsController) Show(c *gin.Context) {
	materialGroupModel := material_group.Get(c.Param("id"))
	if materialGroupModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, materialGroupModel)
}

func (ctrl *MaterialGroupsController) Store(c *gin.Context) {

	request := requests.MaterialGroupRequest{}
	if ok := requests.Validate(c, &request, requests.MaterialGroupSave); !ok {
		return
	}

	materialGroupModel := material_group.MaterialGroup{
		Name:        request.Name,
		Description: request.Description,
	}
	materialGroupModel.Create()
	if materialGroupModel.ID > 0 {
		response.Created(c, materialGroupModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *MaterialGroupsController) Update(c *gin.Context) {

	materialGroupModel := material_group.Get(c.Param("id"))
	if materialGroupModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// if ok := policies.CanModifyMaterialGroup(c, materialGroupModel); !ok {
	// 	response.Abort403(c)
	// 	return
	// }

	request := requests.MaterialGroupRequest{}
	bindOk := requests.Validate(c, &request, requests.MaterialGroupSave)
	if !bindOk {
		return
	}

	materialGroupModel.Name = request.Name
	materialGroupModel.Description = request.Description

	rowsAffected := materialGroupModel.Save()
	if rowsAffected > 0 {
		response.Data(c, materialGroupModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *MaterialGroupsController) Delete(c *gin.Context) {

	materialGroupModel := material_group.Get(c.Param("id"))
	if materialGroupModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyMaterialGroup(c, materialGroupModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := materialGroupModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *MaterialGroupsController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	materialGroupModel := material_group.MaterialGroup{}
	if ok := policies.CanModifyMaterialGroup(c, materialGroupModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := materialGroupModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}
