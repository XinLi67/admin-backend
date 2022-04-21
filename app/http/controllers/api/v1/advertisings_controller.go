package v1

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gohub/app/http/assemblies"
	"gohub/app/models/advertising"
	"gohub/app/models/advertising_plan"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/paginator"
	"gohub/pkg/response"
	"gohub/utils"

	"github.com/gin-gonic/gin"
)

type AdvertisingsController struct {
	BaseAPIController
}

func (ctrl *AdvertisingsController) Index(c *gin.Context) {

	params := c.Query("params")
	status := c.Query("status")
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	// data, pager := advertising.Paginate(c, 0)
	//data, pager := advertising.Search(c, 0)
	var data []advertising.Advertising
	var pager paginator.Paging
	if len(status) > 0 && len(params) > 0 {
		data, pager = advertising.PaginateByStatusAndParams(c, 0, status, params)
	} else if len(params) > 0 {
		data, pager = advertising.Paginate2(c, 0, params)
	} else if len(status) > 0 {
		data, pager = advertising.PaginateByStatus(c, 0, status)
	} else {
		data, pager = advertising.Paginate(c, 0)
	}

	advertisings := assemblies.AdvertisingAssemblyFromModelList(data, len(data))
	response.JSON(c, gin.H{
		"data":  advertisings,
		"pager": pager,
	})
}

func (ctrl *AdvertisingsController) IndexByAdvertisingPosId(c *gin.Context) {
	listData := advertising_plan.GetAll(c.Param("id"))
	if len(listData) <= 0 {
		response.Abort404(c)
		return
	}
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	//从缓存中取数据
	response.Data(c, advertising_plan.AllCached(c.Param("id")))

	//分页实现
	//data, pager := advertising.Paginate2(c, 10,  c.Param("id"))
	//advertisings := assemblies.AdvertisingAssemblyFromModelList(data, len(data))
	//response.JSON(c, gin.H{
	//	"data":  advertisings,
	//	"pager": pager,
	//})
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
		PushContent:           request.PushContent,
		PushTitle:             request.PushTitle,
		AdvertisingCreativity: request.AdvertisingCreativity,
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
	advertisingModel.PushContent = request.PushContent
	advertisingModel.PushTitle = request.PushTitle
	advertisingModel.AdvertisingCreativity = request.AdvertisingCreativity

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

//数据导出
func (ctrl *AdvertisingsController) Export(c *gin.Context) {

	listData := advertising.All2()
	f := excelize.NewFile() // 设置单元格的值
	//// 这里设置表头
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "广告编号")
	f.SetCellValue("Sheet1", "C1", "广告位编号")
	f.SetCellValue("Sheet1", "D1", "创建者编号")
	f.SetCellValue("Sheet1", "E1", "部门ID")
	f.SetCellValue("Sheet1", "F1", "标题")
	f.SetCellValue("Sheet1", "G1", "类型")
	f.SetCellValue("Sheet1", "H1", "跳转")
	f.SetCellValue("Sheet1", "I1", "素材ID")
	f.SetCellValue("Sheet1", "J1", "素材类型")
	f.SetCellValue("Sheet1", "K1", "尺寸")
	f.SetCellValue("Sheet1", "L1", "跳转参数")
	f.SetCellValue("Sheet1", "M1", "描述")
	f.SetCellValue("Sheet1", "N1", "状态")

	line := 1

	//fruits := getFruits()
	// 循环写入数据
	for _, v := range listData {
		line++
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", line), v.ID)
		//fmt.Println("AdvertisingNo:",v.AdvertisingNo)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", line), v.AdvertisingNo)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", line), v.AdvertisingPositionId)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", line), v.CreatorId)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", line), v.DepartmentId)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", line), v.Title)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", line), v.Type)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", line), v.RedirectTo)
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", line), v.MaterialId)
		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", line), v.MaterialType)
		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", line), v.Size)
		f.SetCellValue("Sheet1", fmt.Sprintf("L%d", line), v.RedirectParams)
		f.SetCellValue("Sheet1", fmt.Sprintf("M%d", line), v.Description)
		f.SetCellValue("Sheet1", fmt.Sprintf("N%d", line), v.Status)
	}

	var fileName = utils.RandFileName()
	var fullPath = "G:/studyFile/" + fileName + ".xlsx"

	// 保存文件
	if err := f.SaveAs(fullPath); err != nil {
		fmt.Println(err)
	}

	response.Data(c, "文件保存为:"+fullPath)
}
