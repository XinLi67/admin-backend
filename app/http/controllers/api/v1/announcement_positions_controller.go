package v1

import (
	"fmt"
	"gohub/app/http/assemblies"
	"gohub/app/models/announcement_position"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/paginator"
	"gohub/pkg/response"
	"gohub/utils"
	"github.com/xuri/excelize/v2"
	"github.com/gin-gonic/gin"
)

type AnnouncementPositionsController struct {
	BaseAPIController
}

func (ctrl *AnnouncementPositionsController) Index(c *gin.Context) {
	status := c.Query("status")
	name := c.Query("name")

	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	var data []announcement_position.AnnouncementPosition
	var pager paginator.Paging
	if len(status) > 0 || len(name) > 0 {
		data, pager = announcement_position.Paginate2(c, 0)
	} else{
		data, pager = announcement_position.Paginate(c, 0)
	}

	announcementPositions := assemblies.AnnouncementPositionAssemblyFromModelList(data, len(data))
	response.JSON(c, gin.H{
		"data":  announcementPositions,
		"pager": pager,
	})
}

func (ctrl *AnnouncementPositionsController) Show(c *gin.Context) {
	announcementPositionModel := announcement_position.Get(c.Param("id"))
	if announcementPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}
	announcementPositionAssembly := assemblies.AnnouncementPositionAssemblyFromModel(announcementPositionModel)
	response.Data(c, announcementPositionAssembly)
}

func (ctrl *AnnouncementPositionsController) Store(c *gin.Context) {

	request := requests.AnnouncementPositionRequest{}
	if ok := requests.Validate(c, &request, requests.AnnouncementPositionSave); !ok {
		return
	}

	announcementPositionModel := announcement_position.AnnouncementPosition{
		Name:        request.Name,
		ChannelId:   request.ChannelId,
		Code:        request.Code,
		Height:      request.Height,
		Weight:      request.Weight,
		Status:      request.Status,
		Description: request.Description,
	}
	announcementPositionModel.Create()
	if announcementPositionModel.ID > 0 {
		response.Created(c, announcementPositionModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AnnouncementPositionsController) Update(c *gin.Context) {

	announcementPositionModel := announcement_position.Get(c.Param("id"))
	if announcementPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAnnouncementPosition(c, announcementPositionModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AnnouncementPositionRequest{}
	bindOk := requests.Validate(c, &request, requests.AnnouncementPositionSave)
	if !bindOk {
		return
	}

	announcementPositionModel.Name = request.Name
	announcementPositionModel.ChannelId = request.ChannelId
	announcementPositionModel.Code = request.Code
	announcementPositionModel.Height = request.Height
	announcementPositionModel.Weight = request.Weight
	announcementPositionModel.Status = request.Status
	announcementPositionModel.Description = request.Description

	rowsAffected := announcementPositionModel.Save()
	if rowsAffected > 0 {
		response.Data(c, announcementPositionModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AnnouncementPositionsController) Delete(c *gin.Context) {

	announcementPositionModel := announcement_position.Get(c.Param("id"))
	if announcementPositionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAnnouncementPosition(c, announcementPositionModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := announcementPositionModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *AnnouncementPositionsController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	announcementPositionModel := announcement_position.AnnouncementPosition{}
	if ok := policies.CanModifyAnnouncementPosition(c, announcementPositionModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := announcementPositionModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}

//数据导出
func (ctrl *AnnouncementPositionsController) Export(c *gin.Context) {

	listData := announcement_position.All2()
	f := excelize.NewFile() // 设置单元格的值
	//// 这里设置表头
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "公告位名称")
	f.SetCellValue("Sheet1", "C1", "渠道ID")
	f.SetCellValue("Sheet1", "D1", "公告位编码")
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

	c.Writer.Header().Add("Content-Disposition",fmt.Sprintf("attachment;fileName=%s",fileName))
	c.Writer.Header().Add("Content-Type", "application/octet-stream;charset=utf-8")

	c.File(fullPath)
}
