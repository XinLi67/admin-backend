package v1

import (
	"encoding/json"
	"gohub/app/http/assemblies"
	"gohub/app/models/material_group"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/response"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MaterialGroupsController struct {
	BaseAPIController
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
	//根据ParentId获取Path
	Path := GetPath(request.ParentId)
	materialGroupModel := material_group.MaterialGroup{
		Name:        request.Name,
		Description: request.Description,
		ParentId:    request.ParentId,
		Path:        Path,
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

	if ok := policies.CanModifyMaterialGroup(c, materialGroupModel); !ok {
		response.Abort403(c)
		return
	}

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
	id := c.Param("id")
	materialGroupModel := material_group.Get(id)
	if materialGroupModel.ID == 0 {
		response.Abort404(c)
		return
	}
	count := material_group.GetCountById(id)
	if count > 0 {

		response.JSON(c, gin.H{
			"message": "不能删除",
		})
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

//文件夹查找
func (ctrl *MaterialGroupsController) GetDocumentById(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	data, pager := material_group.GetDocumentById(c, 0, c.Param("id"))
	materialGroups := assemblies.MaterialGroupAssemblyFromModelList(data)
	response.JSON(c, gin.H{
		"data":  materialGroups,
		"pager": pager,
	})
}

//文件导航获取
type groupItem struct {
	ID       uint64      `json:"id"`
	Name     string      `json:"name"`
	ParentId uint64      `json:"parent_id"`
	Path     string      `json:"path"`
	Chileren []groupItem `json:"children"`
}

func GroupAssemblyFromModelList(data []material_group.MaterialGroup) []groupItem {
	group := make([]groupItem, len(data))
	for i, v := range data {
		group[i] = groupItem{
			ID:       v.ID,
			Name:     v.Name,
			ParentId: v.ParentId,
			Path:     v.Path}
	}
	return group
}
func (ctrl *MaterialGroupsController) GetTree(c *gin.Context) {
	data1 := material_group.All()
	groupTrees := GroupAssemblyFromModelList(data1)
	id := c.Param("id")
	intNum, _ := strconv.Atoi(id)
	data := treeDate(groupTrees, uint64(intNum))
	response.Data(c, data)
}
func treeDate(groupItem1 []groupItem, id uint64) []groupItem {
	var groupTrees []groupItem
	for _, v := range groupItem1 {
		if v.ID == id {
			if v.ParentId != 0 {
				v.Chileren = append(groupTrees, treeDate(groupItem1, v.ParentId)...)

			}
			groupTrees = append(groupTrees, v)

		}
	}
	return groupTrees
}

//获取PATH
func GetPath(id uint64) string {
	materialGroupModel := material_group.All()
	data := PathDate(materialGroupModel, id)
	data = append(data, int(id))
	b, _ := json.Marshal(data)
	result := string(b)
	result = strings.Trim(result, "[]")
	return result
}
func PathDate(materialGroup []material_group.MaterialGroup, id uint64) []int {
	var path []int
	for _, v := range materialGroup {
		if v.ID == id {
			if v.ParentId != 0 {
				path = append(path, PathDate(materialGroup, v.ParentId)...)
			}
			path = append(path, int(v.ParentId))

		}
	}
	return path
}

func  bottomDate(materialGroup []material_group.MaterialGroup) []material_group.MaterialGroup {
	var materialGroup1 []material_group.MaterialGroup
	for _, v := range materialGroup {
		b := strconv.Itoa(int(v.ID))
		c := string(b)
		count := material_group.GetCountById(c)
		if count == 0 {
			materialGroup1 = append(materialGroup1, v)
		}
	}
	return materialGroup1
}

func (ctrl *MaterialGroupsController) GetBottom(c *gin.Context) {
	materialGroupModel := material_group.All()
	data := bottomDate(materialGroupModel)
	response.Data(c, data)
}