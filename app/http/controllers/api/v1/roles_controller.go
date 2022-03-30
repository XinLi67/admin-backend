package v1

import (
	"gohub/app/models/role"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type RolesController struct {
	BaseAPIController
}

func (ctrl *RolesController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := role.Paginate(c, 0)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *RolesController) Show(c *gin.Context) {
	roleModel := role.Get(c.Param("id"))
	if roleModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, roleModel)
}

func (ctrl *RolesController) Store(c *gin.Context) {

	request := requests.RoleRequest{}
	if ok := requests.Validate(c, &request, requests.RoleSave); !ok {
		return
	}

	roleModel := role.Role{
		Name:        request.Name,
		GuardName:   request.GuardName,
		Description: request.Description,
	}
	roleModel.Create()
	if roleModel.ID > 0 {
		response.Created(c, roleModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *RolesController) Update(c *gin.Context) {

	roleModel := role.Get(c.Param("id"))
	if roleModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyRole(c, roleModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.RoleRequest{}
	bindOk := requests.Validate(c, &request, requests.RoleSave)
	if !bindOk {
		return
	}

	roleModel.Name = request.Name
	roleModel.GuardName = request.GuardName
	roleModel.Description = request.Description
	rowsAffected := roleModel.Save()
	if rowsAffected > 0 {
		response.Data(c, roleModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *RolesController) Delete(c *gin.Context) {

	roleModel := role.Get(c.Param("id"))
	if roleModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyRole(c, roleModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := roleModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *RolesController) GetGuardNameRoles(c *gin.Context) {
	roles := role.All()
	response.Data(c, roles)
}

func (ctrl *RolesController) UpdateGuardName(c *gin.Context) {
	roles := role.All()
	response.Data(c, roles)
}

func (ctrl *RolesController) GetRolePermissions(c *gin.Context) {
	roles := role.All()
	response.Data(c, roles)
}

func (ctrl *RolesController) AssignPermissions(c *gin.Context) {
	roles := role.All()
	response.Data(c, roles)
}
