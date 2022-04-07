package v1

import (
	"gohub/app/models/menu"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type MenusController struct {
	BaseAPIController
}

type menuList []menu.Menu

type menuItem struct {
	menu.Menu
	Chileren []menuItem `json:"children"`
}

func (ctrl *MenusController) MyMenu(c *gin.Context) {
	menus := menu.All()

	mymenu := menuList(menus)

	menuTree := mymenu.processToTree(0, 0)

	response.Data(c, menuTree)
}

func (ctrl *MenusController) Index(c *gin.Context) {
	menus := menu.All()
	mymenu := menuList(menus)

	menuTree := mymenu.processToTree(0, 0)

	response.Data(c, menuTree)
}

func (ctrl *MenusController) Show(c *gin.Context) {
	menuModel := menu.Get(c.Param("id"))
	if menuModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, menuModel)
}

func (ctrl *MenusController) Store(c *gin.Context) {

	request := requests.MenuRequest{}
	if ok := requests.Validate(c, &request, requests.MenuSave); !ok {
		return
	}

	menuModel := menu.Menu{
		ParentId:       request.ParentId,
		Name:           request.Name,
		Icon:           request.Icon,
		Uri:            request.Uri,
		IsLink:         request.IsLink,
		PermissionName: request.PermissionName,
		GuardName:      request.GuardName,
		Sequence:       request.Sequence,
	}
	menuModel.Create()
	if menuModel.ID > 0 {
		response.Created(c, menuModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *MenusController) Update(c *gin.Context) {

	menuModel := menu.Get(c.Param("id"))
	if menuModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyMenu(c, menuModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.MenuRequest{}
	bindOk := requests.Validate(c, &request, requests.MenuUpdate)
	if !bindOk {
		return
	}

	menuModel.ParentId = request.ParentId
	menuModel.Name = request.Name
	menuModel.Icon = request.Icon
	menuModel.Uri = request.Uri
	menuModel.IsLink = request.IsLink
	menuModel.PermissionName = request.PermissionName
	menuModel.GuardName = request.GuardName
	menuModel.Sequence = request.Sequence

	rowsAffected := menuModel.Save()
	if rowsAffected > 0 {
		response.Data(c, menuModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *MenusController) Delete(c *gin.Context) {

	menuModel := menu.Get(c.Param("id"))
	if menuModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyMenu(c, menuModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := menuModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *MenusController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	menuModel := menu.Menu{}
	if ok := policies.CanModifyMenu(c, menuModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := menuModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}

func (m *menuList) processToTree(pid uint64, level uint64) []menuItem {
	var menuTree []menuItem
	if level == 10 {
		return menuTree
	}

	list := m.findChildren(pid)
	if len(list) == 0 {
		return menuTree
	}

	for _, v := range list {
		child := m.processToTree(v.ID, level+1)
		menuTree = append(menuTree, menuItem{v, child})
	}

	return menuTree
}

func (m *menuList) findChildren(pid uint64) []menu.Menu {
	child := []menu.Menu{}

	for _, v := range *m {
		if v.ParentId == pid {
			child = append(child, v)
		}
	}

	return child
}
