package v1

import (
	"gohub/app/models/permission"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type PermissionsController struct {
	BaseAPIController
}

func (ctrl *PermissionsController) All(c *gin.Context) {
	permissions := permission.All()

	allPermissions := make([]string, len(permissions))
	for i := range allPermissions {
		allPermissions[i] = permissions[i].Name
	}

	response.Data(c, allPermissions)
}

func (ctrl *PermissionsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := permission.Paginate(c, 0)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *PermissionsController) Show(c *gin.Context) {
	permissionModel := permission.Get(c.Param("id"))
	if permissionModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, permissionModel)
}

func (ctrl *PermissionsController) Store(c *gin.Context) {

	request := requests.PermissionRequest{}
	if ok := requests.Validate(c, &request, requests.PermissionSave); !ok {
		return
	}

	permissionModel := permission.Permission{
		PermissionGroupId: request.PermissionGroupId,
		Name:              request.Name,
		Icon:              request.Icon,
		GuardName:         request.GuardName,
		DisplayName:       request.DisplayName,
		Description:       request.Description,
		Sequence:          request.Sequence,
	}
	permissionModel.Create()
	if permissionModel.ID > 0 {
		response.Created(c, permissionModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *PermissionsController) Update(c *gin.Context) {

	permissionModel := permission.Get(c.Param("id"))
	if permissionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyPermission(c, permissionModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.PermissionRequest{}
	bindOk := requests.Validate(c, &request, requests.PermissionSave)
	if !bindOk {
		return
	}

	permissionModel.PermissionGroupId = request.PermissionGroupId
	permissionModel.Name = request.Name
	permissionModel.Icon = request.Icon
	permissionModel.GuardName = request.GuardName
	permissionModel.DisplayName = request.DisplayName
	permissionModel.Description = request.Description
	permissionModel.Sequence = request.Sequence

	rowsAffected := permissionModel.Save()
	if rowsAffected > 0 {
		response.Data(c, permissionModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *PermissionsController) Delete(c *gin.Context) {

	permissionModel := permission.Get(c.Param("id"))
	if permissionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyPermission(c, permissionModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := permissionModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *PermissionsController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	permissionModel := permission.Permission{}
	if ok := policies.CanModifyPermission(c, permissionModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := permissionModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}
