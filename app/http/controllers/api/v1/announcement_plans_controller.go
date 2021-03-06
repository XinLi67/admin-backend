package v1

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gohub/app/http/assemblies"
	"gohub/app/models/announcement_plan"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
	"gohub/pkg/response"
	"gohub/utils"

	"github.com/gin-gonic/gin"
)

type AnnouncementPlansController struct {
	BaseAPIController
}

func (ctrl *AnnouncementPlansController) Index(c *gin.Context) {
	status := c.Query("audit_status")
	name := c.Query("name")

	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	var data []announcement_plan.AnnouncementPlan
	var pager paginator.Paging
	if len(status) > 0 || len(name) > 0 {
		data, pager = announcement_plan.Paginate2(c, 0)
	} else {
		data, pager = announcement_plan.Paginate(c, 0)
	}

	announcementPlans := assemblies.AnnouncementPlanAssemblyFromModelList(data, len(data))
	response.JSON(c, gin.H{
		"data":  announcementPlans,
		"pager": pager,
	})
}

func (ctrl *AnnouncementPlansController) Show(c *gin.Context) {
	announcementPlanModel := announcement_plan.Get(c.Param("id"))
	if announcementPlanModel.ID == 0 {
		response.Abort404(c)
		return
	}
	announcementPlanAssembly := assemblies.AnnouncementPlanAssemblyFromModel(announcementPlanModel)
	response.Data(c, announcementPlanAssembly)
}

func (ctrl *AnnouncementPlansController) Store(c *gin.Context) {

	request := requests.AnnouncementPlanRequest{}
	if ok := requests.Validate(c, &request, requests.AnnouncementPlanSave); !ok {
		return
	}

	announcementPlanModel := announcement_plan.AnnouncementPlan{
		Name:                   request.Name,
		CreatorId:              request.CreatorId,
		AnnouncementId:         request.AnnouncementId,
		AnnouncementType:       request.AnnouncementType,
		AnnouncementPositionId: request.AnnouncementPositionId,
		Order:                  request.Order,
		SchedulingDate:         request.SchedulingDate,
		SchedulingTime:         request.SchedulingTime,
		StartDate:              request.StartDate,
		EndDate:                request.EndDate,
		StartTime:              request.StartTime,
		Endime:                 request.Endime,
		AuditStatus:            request.AuditStatus,
		PresentStatus:          request.PresentStatus,
	}
	announcementPlanModel.Create()
	if announcementPlanModel.ID > 0 {
		response.Created(c, announcementPlanModel)
	} else {
		response.Abort500(c, "??????????????????????????????~")
	}
}

//todo ????????????????????????
func (ctrl *AnnouncementPlansController) BatchStore(c *gin.Context) {

	request := requests.AnnouncementBatchStoreRequest{}
	bindOk := requests.Validate(c, &request, requests.AnnouncementBatchStore)
	if !bindOk {
		return
	}

	var plans = request.Data
	database.DB.Create(&plans)

	for _, plan := range plans {
		fmt.Println("??????????????????ID:", plan.ID)
	}

	response.Data(c, map[string]int{
		"????????????????????????": len(plans),
	})

}

func (ctrl *AnnouncementPlansController) Update(c *gin.Context) {

	announcementPlanModel := announcement_plan.Get(c.Param("id"))
	if announcementPlanModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAnnouncementPlan(c, announcementPlanModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AnnouncementPlanRequest{}
	bindOk := requests.Validate(c, &request, requests.AnnouncementPlanSave)
	if !bindOk {
		return
	}

	announcementPlanModel.Name = request.Name
	announcementPlanModel.CreatorId = request.CreatorId
	announcementPlanModel.AnnouncementId = request.AnnouncementId
	announcementPlanModel.AnnouncementType = request.AnnouncementType
	announcementPlanModel.AnnouncementPositionId = request.AnnouncementPositionId
	announcementPlanModel.Order = request.Order
	announcementPlanModel.SchedulingDate = request.SchedulingDate
	announcementPlanModel.SchedulingTime = request.SchedulingTime
	announcementPlanModel.StartDate = request.StartDate
	announcementPlanModel.EndDate = request.EndDate
	announcementPlanModel.StartTime = request.StartTime
	announcementPlanModel.Endime = request.Endime
	announcementPlanModel.AuditStatus = request.AuditStatus
	announcementPlanModel.PresentStatus = request.PresentStatus

	rowsAffected := announcementPlanModel.Save()
	if rowsAffected > 0 {
		response.Data(c, announcementPlanModel)
	} else {
		response.Abort500(c, "??????????????????????????????~")
	}
}

func (ctrl *AnnouncementPlansController) Delete(c *gin.Context) {

	announcementPlanModel := announcement_plan.Get(c.Param("id"))
	if announcementPlanModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAnnouncementPlan(c, announcementPlanModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := announcementPlanModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "??????????????????????????????~")
}

func (ctrl *AnnouncementPlansController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	announcementPlanModel := announcement_plan.AnnouncementPlan{}
	if ok := policies.CanModifyAnnouncementPlan(c, announcementPlanModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := announcementPlanModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}

//????????????
func (ctrl *AnnouncementPlansController) Export(c *gin.Context) {

	listData:=announcement_plan.All2()
	f := excelize.NewFile()// ?????????????????????
	// ??????????????????
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "??????????????????")
	f.SetCellValue("Sheet1", "C1", "?????????ID")
	f.SetCellValue("Sheet1", "D1", "??????ID")
	f.SetCellValue("Sheet1", "E1", "????????????")
	f.SetCellValue("Sheet1", "F1", "?????????ID")
	f.SetCellValue("Sheet1", "G1", "?????????")
	f.SetCellValue("Sheet1", "H1", "????????????")
	f.SetCellValue("Sheet1", "I1", "????????????")
	f.SetCellValue("Sheet1", "J1", "????????????")
	f.SetCellValue("Sheet1", "K1", "????????????")
	f.SetCellValue("Sheet1", "L1", "????????????")
	f.SetCellValue("Sheet1", "M1", "????????????")
	f.SetCellValue("Sheet1", "N1", "????????????")
	f.SetCellValue("Sheet1", "O1", "????????????")

	line := 1

	// ??????????????????
	for _, v := range listData {
		line++
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", line), v.ID)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", line), v.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", line), v.CreatorId)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", line), v.AnnouncementId)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", line), v.AnnouncementType)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", line), v.AnnouncementPositionId)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", line), v.Order)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", line), v.SchedulingDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", line), v.SchedulingTime)
		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", line), v.StartDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", line), v.EndDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("L%d", line), v.StartTime)
		f.SetCellValue("Sheet1", fmt.Sprintf("M%d", line), v.Endime)
		f.SetCellValue("Sheet1", fmt.Sprintf("N%d", line), v.AuditStatus)
		f.SetCellValue("Sheet1", fmt.Sprintf("O%d", line), v.PresentStatus)

	}

	var fileName=utils.RandFileName()
	var fullPath="G:/studyFile/"+fileName+".xlsx"

	// ????????????
	if err := f.SaveAs(fullPath); err != nil {
		fmt.Println(err)
	}

	c.Writer.Header().Add("Content-Disposition",fmt.Sprintf("attachment;fileName=%s",fileName))
	c.Writer.Header().Add("Content-Type", "application/octet-stream;charset=utf-8")

	c.File(fullPath)
}
