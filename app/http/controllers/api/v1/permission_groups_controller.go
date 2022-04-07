package v1

import (
	"gohub/app/models/permission_group"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type PermissionGroupsController struct {
	BaseAPIController
}

func (ctrl *PermissionGroupsController) All(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := permission_group.Paginate(c, 0)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *PermissionGroupsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := permission_group.Paginate(c, 0)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *PermissionGroupsController) Show(c *gin.Context) {
	permissionGroupModel := permission_group.Get(c.Param("id"))
	if permissionGroupModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, permissionGroupModel)
}

func (ctrl *PermissionGroupsController) Store(c *gin.Context) {

	request := requests.PermissionGroupRequest{}
	if ok := requests.Validate(c, &request, requests.PermissionGroupSave); !ok {
		return
	}

	permissionGroupModel := permission_group.PermissionGroup{
		Name:        request.Name,
		Description: request.Description,
	}
	permissionGroupModel.Create()
	if permissionGroupModel.ID > 0 {
		response.Created(c, permissionGroupModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *PermissionGroupsController) Update(c *gin.Context) {

	permissionGroupModel := permission_group.Get(c.Param("id"))
	if permissionGroupModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyPermissionGroup(c, permissionGroupModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.PermissionGroupRequest{}
	bindOk := requests.Validate(c, &request, requests.PermissionGroupSave)
	if !bindOk {
		return
	}

	permissionGroupModel.Name = request.Name
	permissionGroupModel.Description = request.Description

	rowsAffected := permissionGroupModel.Save()
	if rowsAffected > 0 {
		response.Data(c, permissionGroupModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *PermissionGroupsController) Delete(c *gin.Context) {

	permissionGroupModel := permission_group.Get(c.Param("id"))
	if permissionGroupModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyPermissionGroup(c, permissionGroupModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := permissionGroupModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *PermissionGroupsController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	permissionGroupModel := permission_group.PermissionGroup{}
	if ok := policies.CanModifyPermissionGroup(c, permissionGroupModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := permissionGroupModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}
