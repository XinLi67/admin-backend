package v1

import (
	"fmt"
	"gohub/app/http/assemblies"
	"gohub/app/models/advertising_position"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/paginator"
	"gohub/pkg/response"
	"gohub/utils"

	"github.com/xuri/excelize/v2"

	"github.com/gin-gonic/gin"
)

type AdvertisingPositionsController struct {
	BaseAPIController
}

func (ctrl *AdvertisingPositionsController) Index(c *gin.Context) {
	status := c.Query("status")
	params := c.Query("params")

	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	var data []advertising_position.AdvertisingPosition
	var pager paginator.Paging
	if len(status) > 0 && len(params) > 0 {
		data, pager = advertising_position.PaginateByStatusAndParams(c, 0, status, params)
	} else if len(params) > 0 {
		data, pager = advertising_position.PaginateByName(c, 0, params)
	} else if len(status) > 0 {
		data, pager = advertising_position.PaginateByStatus(c, 0, status)
	} else {
		data, pager = advertising_position.Paginate(c, 0)
	}

	advertisingPositions := assemblies.AdvertisingPositionAssemblyFromModelList(data, len(data))
	response.JSON(c, gin.H{
		"data":  advertisingPositions,
		"pager": pager,
	})
}

func (ctrl *AdvertisingPositionsController) Show(c *gin.Context) {
	advertisingPositionModel := advertising_position.Get(c.Param("id"))
	if advertisingPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}
	advertisingPositionAssembly := assemblies.AdvertisingPositionAssemblyFromModel(advertisingPositionModel)
	response.Data(c, advertisingPositionAssembly)
}

func (ctrl *AdvertisingPositionsController) Store(c *gin.Context) {

	request := requests.AdvertisingPositionRequest{}
	if ok := requests.Validate(c, &request, requests.AdvertisingPositionSave); !ok {
		return
	}

	advertisingPositionModel := advertising_position.AdvertisingPosition{
		Name:        request.Name,
		ChannelId:   request.ChannelId,
		Code:        request.Code,
		Height:      request.Height,
		Weight:      request.Weight,
		Status:      request.Status,
		Description: request.Description,
	}
	advertisingPositionModel.Create()
	if advertisingPositionModel.ID > 0 {
		response.Created(c, advertisingPositionModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AdvertisingPositionsController) Update(c *gin.Context) {

	advertisingPositionModel := advertising_position.Get(c.Param("id"))
	if advertisingPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertisingPosition(c, advertisingPositionModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AdvertisingPositionRequest{}
	bindOk := requests.Validate(c, &request, requests.AdvertisingPositionSave)
	if !bindOk {
		return
	}

	advertisingPositionModel.Name = request.Name
	advertisingPositionModel.ChannelId = request.ChannelId
	advertisingPositionModel.Code = request.Code
	advertisingPositionModel.Height = request.Height
	advertisingPositionModel.Weight = request.Weight
	advertisingPositionModel.Status = request.Status
	advertisingPositionModel.Description = request.Description

	rowsAffected := advertisingPositionModel.Save()
	if rowsAffected > 0 {
		response.Data(c, advertisingPositionModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AdvertisingPositionsController) Delete(c *gin.Context) {

	advertisingPositionModel := advertising_position.Get(c.Param("id"))
	if advertisingPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertisingPosition(c, advertisingPositionModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := advertisingPositionModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *AdvertisingPositionsController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	advertisingPositionModel := advertising_position.AdvertisingPosition{}
	if ok := policies.CanModifyAdvertisingPosition(c, advertisingPositionModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := advertisingPositionModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}

//数据导出
func (ctrl *AdvertisingPositionsController) Export(c *gin.Context) {

	listData := advertising_position.All2()
	f := excelize.NewFile() // 设置单元格的值
	//// 这里设置表头
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "广告位名称")
	f.SetCellValue("Sheet1", "C1", "渠道编号")
	f.SetCellValue("Sheet1", "D1", "广告位代码")
	f.SetCellValue("Sheet1", "E1", "高度")
	f.SetCellValue("Sheet1", "F1", "宽度")
	f.SetCellValue("Sheet1", "G1", "状态")
	f.SetCellValue("Sheet1", "H1", "描述")

	line := 1

	//fruits := getFruits()
	// 循环写入数据
	for _, v := range listData {
		line++
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", line), v.ID)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", line), v.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", line), v.ChannelId)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", line), v.Code)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", line), v.Height)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", line), v.Weight)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", line), v.Status)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", line), v.Description)
	}

	var fileName = utils.RandFileName()
	var fullPath = "G:/studyFile/" + fileName + ".xlsx"

	// 保存文件
	if err := f.SaveAs(fullPath); err != nil {
		fmt.Println(err)
	}

	response.Data(c, "文件保存为:"+fullPath)
}
