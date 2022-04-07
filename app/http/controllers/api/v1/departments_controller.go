package v1

import (
	"gohub/app/models/department"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type DepartmentsController struct {
	BaseAPIController
}

type departmentList []department.Department

type departmentItem struct {
	department.Department
	Chileren []departmentItem `json:"children"`
}

func (ctrl *DepartmentsController) MyDepartment(c *gin.Context) {
	departments := department.All()

	mydepartment := departmentList(departments)

	departmentTree := mydepartment.processToTree(0, 0)

	response.Data(c, departmentTree)
}

func (ctrl *DepartmentsController) Index(c *gin.Context) {
	departments := department.All()
	mydepartment := departmentList(departments)

	departmentTree := mydepartment.processToTree(0, 0)

	response.Data(c, departmentTree)
}

func (ctrl *DepartmentsController) Show(c *gin.Context) {
	departmentModel := department.Get(c.Param("id"))
	if departmentModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, departmentModel)
}

func (ctrl *DepartmentsController) Store(c *gin.Context) {

	request := requests.DepartmentRequest{}
	if ok := requests.Validate(c, &request, requests.DepartmentSave); !ok {
		return
	}

	departmentModel := department.Department{
		ParentId:    request.ParentId,
		Name:        request.Name,
		Phone:       request.Phone,
		LinkMan:     request.LinkMan,
		Address:     request.Address,
		Description: request.Description,
	}
	departmentModel.Create()
	if departmentModel.ID > 0 {
		response.Created(c, departmentModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *DepartmentsController) Update(c *gin.Context) {

	departmentModel := department.Get(c.Param("id"))
	if departmentModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyDepartment(c, departmentModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.DepartmentRequest{}
	bindOk := requests.Validate(c, &request, requests.DepartmentSave)
	if !bindOk {
		return
	}

	departmentModel.ParentId = request.ParentId
	departmentModel.Name = request.Name
	departmentModel.Phone = request.Phone
	departmentModel.LinkMan = request.LinkMan
	departmentModel.Address = request.Address
	departmentModel.Description = request.Description

	rowsAffected := departmentModel.Save()
	if rowsAffected > 0 {
		response.Data(c, departmentModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *DepartmentsController) Delete(c *gin.Context) {

	departmentModel := department.Get(c.Param("id"))
	if departmentModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyDepartment(c, departmentModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := departmentModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *DepartmentsController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	departmentModel := department.Department{}
	if ok := policies.CanModifyDepartment(c, departmentModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := departmentModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}

func (m *departmentList) processToTree(pid uint64, level uint64) []departmentItem {
	var departmentTree []departmentItem
	if level == 10 {
		return departmentTree
	}

	list := m.findChildren(pid)
	if len(list) == 0 {
		return departmentTree
	}

	for _, v := range list {
		child := m.processToTree(v.ID, level+1)
		departmentTree = append(departmentTree, departmentItem{v, child})
	}

	return departmentTree
}

func (m *departmentList) findChildren(pid uint64) []department.Department {
	child := []department.Department{}

	for _, v := range *m {
		if v.ParentId == pid {
			child = append(child, v)
		}
	}

	return child
}
