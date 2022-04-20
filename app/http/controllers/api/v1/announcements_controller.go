package v1

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gohub/app/http/assemblies"
	"gohub/app/models/announcement"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/paginator"
	"gohub/pkg/response"
	"gohub/utils"

	"github.com/gin-gonic/gin"
)

type AnnouncementsController struct {
	BaseAPIController
}

func (ctrl *AnnouncementsController) Index(c *gin.Context) {
	params := c.Query("params")
	status := c.Query("status")
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	var data []announcement.Announcement
	var pager paginator.Paging
	if len(status) > 0 {
		data, pager = announcement.PaginateByStatus(c, 0, status)
	} else if len(params) > 0 {
		data, pager = announcement.Paginate2(c, 0, params)
	} else {
		data, pager = announcement.Paginate(c, 0)
	}

	announcements := assemblies.AnnouncementAssemblyFromModelList(data, len(data))
	response.JSON(c, gin.H{
		"data":  announcements,
		"pager": pager,
	})
}

func (ctrl *AnnouncementsController) Show(c *gin.Context) {
	announcementModel := announcement.Get(c.Param("id"))
	if announcementModel.ID == 0 {
		response.Abort404(c)
		return
	}
	announcementAssembly := assemblies.AnnouncementAssemblyFromModel(announcementModel)
	response.Data(c, announcementAssembly)
}

func (ctrl *AnnouncementsController) Store(c *gin.Context) {

	request := requests.AnnouncementRequest{}
	if ok := requests.Validate(c, &request, requests.AnnouncementSave); !ok {
		return
	}

	announcementModel := announcement.Announcement{
		AnnouncementPositionId: request.AnnouncementPositionId,
		CreatorId:              request.CreatorId,
		DepartmentId:           request.DepartmentId,
		Title:                  request.Title,
		LongTitle:              request.LongTitle,
		Type:                   request.Type,
		Banner:                 request.Banner,
		RedirectTo:             request.RedirectTo,
		RedirectParams:         request.RedirectParams,
		Content:                request.Content,
		Status:                 request.Status,
		AuditReason:            request.AuditReason,
	}
	announcementModel.Create()
	if announcementModel.ID > 0 {
		response.Created(c, announcementModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AnnouncementsController) Update(c *gin.Context) {

	announcementModel := announcement.Get(c.Param("id"))
	if announcementModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAnnouncement(c, announcementModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AnnouncementRequest{}
	bindOk := requests.Validate(c, &request, requests.AnnouncementSave)
	if !bindOk {
		return
	}

	announcementModel.AnnouncementPositionId = request.AnnouncementPositionId
	announcementModel.CreatorId = request.CreatorId
	announcementModel.DepartmentId = request.DepartmentId
	announcementModel.Title = request.Title
	announcementModel.LongTitle = request.LongTitle
	announcementModel.Type = request.Type
	announcementModel.Banner = request.Banner
	announcementModel.RedirectTo = request.RedirectTo
	announcementModel.RedirectParams = request.RedirectParams
	announcementModel.Content = request.Content
	announcementModel.Status = request.Status

	rowsAffected := announcementModel.Save()
	if rowsAffected > 0 {
		response.Data(c, announcementModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AnnouncementsController) Delete(c *gin.Context) {

	announcementModel := announcement.Get(c.Param("id"))
	if announcementModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAnnouncement(c, announcementModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := announcementModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *AnnouncementsController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	announcementModel := announcement.Announcement{}
	if ok := policies.CanModifyAnnouncement(c, announcementModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := announcementModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}

//数据导出
func (ctrl *AnnouncementsController) Export(c *gin.Context) {

	listData := announcement.All2()
	f := excelize.NewFile() // 设置单元格的值
	//// 这里设置表头
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "公告编号")
	f.SetCellValue("Sheet1", "C1", "公告位编号")
	f.SetCellValue("Sheet1", "D1", "创建者编号")
	f.SetCellValue("Sheet1", "E1", "部门ID")
	f.SetCellValue("Sheet1", "F1", "标题")
	f.SetCellValue("Sheet1", "G1", "长公告标题")
	f.SetCellValue("Sheet1", "H1", "类型")
	f.SetCellValue("Sheet1", "I1", "头图链接")
	f.SetCellValue("Sheet1", "J1", "跳转类型")
	f.SetCellValue("Sheet1", "K1", "跳转参数")
	f.SetCellValue("Sheet1", "L1", "内容")
	f.SetCellValue("Sheet1", "M1", "状态")

	line := 1

	//fruits := getFruits()
	// 循环写入数据
	for _, v := range listData {
		line++
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", line), v.ID)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", line), v.AnnouncementNo)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", line), v.AnnouncementPositionId)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", line), v.CreatorId)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", line), v.DepartmentId)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", line), v.Title)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", line), v.LongTitle)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", line), v.Type)
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", line), v.Banner)
		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", line), v.RedirectTo)
		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", line), v.RedirectParams)
		f.SetCellValue("Sheet1", fmt.Sprintf("L%d", line), v.Content)
		f.SetCellValue("Sheet1", fmt.Sprintf("M%d", line), v.Status)

	}

	var fileName = utils.RandFileName()
	var fullPath = "G:/studyFile/" + fileName + ".xlsx"

	// 保存文件
	if err := f.SaveAs(fullPath); err != nil {
		fmt.Println(err)
	}

	response.Data(c, "文件保存为:"+fullPath)
}
