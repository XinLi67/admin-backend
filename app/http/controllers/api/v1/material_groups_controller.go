package v1

import (
	"encoding/json"
	"gohub/app/http/assemblies"
	"gohub/app/models/material_group"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/database"
	"gohub/pkg/response"
	"strconv"
	"strings"

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
	material_group.MaterialGroup
	Chileren []groupItem `json:"children"`
}

func (ctrl *MaterialGroupsController) GetTree(c *gin.Context) {
	var groupTrees []groupItem
	database.DB.Model(material_group.MaterialGroup{}).Find(&groupTrees)
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
	data := getPath(materialGroupModel, id)
	data = append(data, int(id))
	b, _ := json.Marshal(data)
	result := string(b)
	result = strings.Trim(result, "[]")
	return result
}
func getPath(materialGroup []material_group.MaterialGroup, id uint64) []int {
	var path []int
	for _, v := range materialGroup {
		if v.ID == id {
			if v.ParentId != 0 {
				path = append(path, getPath(materialGroup, v.ParentId)...)
			}
			path = append(path, int(v.ParentId))

		}
	}
	return path
}
