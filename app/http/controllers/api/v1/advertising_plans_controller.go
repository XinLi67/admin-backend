package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gohub/app/http/assemblies"
	"gohub/app/models/advertising"
	"gohub/app/models/advertising_plan"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
	"gohub/pkg/response"
	"gohub/utils"
)

type AdvertisingPlansController struct {
	BaseAPIController
}

func (ctrl *AdvertisingPlansController) Index(c *gin.Context) {
	audit_status := c.Query("audit_status")
	name := c.Query("name")

	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	var data []advertising_plan.AdvertisingPlan
	var pager paginator.Paging

	if len(audit_status) > 0 || len(name) > 0 {
		data, pager = advertising_plan.Paginate2(c, 0)
	} else {
		data, pager = advertising_plan.Paginate(c, 0)
	}

	AdvertisingPlans := assemblies.AdvertisingPlanAssemblyFromModelList(data, len(data))
	response.JSON(c, gin.H{
		"data":  AdvertisingPlans,
		"pager": pager,
	})
}

func (ctrl *AdvertisingPlansController) Show(c *gin.Context) {
	advertisingPlanModel := advertising_plan.Get(c.Param("id"))
	if advertisingPlanModel.ID == 0 {
		response.Abort404(c)
		return
	}
	advertisingPlanAssembly := assemblies.AdvertisingPlanAssemblyFromModel(advertisingPlanModel)
	response.Data(c, advertisingPlanAssembly)
}

func (ctrl *AdvertisingPlansController) Store(c *gin.Context) {

	request := requests.AdvertisingPlanRequest{}
	if ok := requests.Validate(c, &request, requests.AdvertisingPlanSave); !ok {
		return
	}

	advertisingPlanModel := advertising_plan.AdvertisingPlan{
		Name:                  request.Name,
		CreatorId:             request.CreatorId,
		AdvertisingId:         request.AdvertisingId,
		AdvertisingType:       request.AdvertisingType,
		AdvertisingPositionId: request.AdvertisingPositionId,
		Order:                 request.Order,
		SchedulingDate:        request.SchedulingDate,
		SchedulingTime:        request.SchedulingTime,
		StartDate:             request.StartDate,
		EndDate:               request.EndDate,
		StartTime:             request.StartTime,
		EndTime:               request.EndTime,
		AuditStatus:           request.AuditStatus,
		PresentStatus:         request.PresentStatus,
	}

	advertisingPlanModel.Create()
	if advertisingPlanModel.ID > 0 {
		response.Created(c, advertisingPlanModel)
	} else {
		response.Abort500(c, "??????????????????????????????~")
	}
}

//todo ????????????????????????
func (ctrl *AdvertisingPlansController) BatchStore(c *gin.Context) {

	request := requests.AdvertigingBatchStoreRequest{}
	bindOk := requests.Validate(c, &request, requests.AdvertigingBatchStore)
	if !bindOk {
		return
	}

	var advertisings = request.Advertisings

	advertisingPlanModel := advertising_plan.AdvertisingPlan{
		Name:                  request.AdvertisingPlan.Name,
		CreatorId:             request.AdvertisingPlan.CreatorId,
		AdvertisingId:         request.AdvertisingPlan.AdvertisingId,
		AdvertisingType:       request.AdvertisingPlan.AdvertisingType,
		AdvertisingPositionId: request.AdvertisingPlan.AdvertisingPositionId,
		Order:                 request.AdvertisingPlan.Order,
		SchedulingDate:        request.AdvertisingPlan.SchedulingDate,
		SchedulingTime:        request.AdvertisingPlan.SchedulingTime,
		StartDate:             request.AdvertisingPlan.StartDate,
		EndDate:               request.AdvertisingPlan.EndDate,
		StartTime:             request.AdvertisingPlan.StartTime,
		EndTime:               request.AdvertisingPlan.EndTime,
		AuditStatus:           request.AdvertisingPlan.AuditStatus,
		PresentStatus:         request.AdvertisingPlan.PresentStatus,
	}

	advertisingPlanModel.Create()
	if advertisingPlanModel.ID > 0 {
		fmt.Println("????????????????????????ID:", advertisingPlanModel.ID)
	}

	var len1 = len(advertisings)
	var advertisingModels []advertising.Advertising = make([]advertising.Advertising, len1)

	fmt.Println(len(advertisingModels))

	for i, item := range advertisings {
		advertisingModels[i].AdvertisingNo = item.AdvertisingNo
		advertisingModels[i].AdvertisingPositionId = item.AdvertisingPositionId
		advertisingModels[i].StartTime = item.StartTime
		advertisingModels[i].EndTime = item.EndTime
	}

	for _, item := range advertisingModels {
		database.DB.Model(&advertising.Advertising{}).Where("advertising_no=?", item.AdvertisingNo).Updates(advertising.Advertising{StartTime: item.StartTime, EndTime: item.EndTime, SchedulingTime: item.SchedulingTime, AdvertisingPositionId: item.AdvertisingPositionId})
	}

	response.Created(c, advertisingPlanModel)

}

func (ctrl *AdvertisingPlansController) Update(c *gin.Context) {

	advertisingPlanModel := advertising_plan.Get(c.Param("id"))
	if advertisingPlanModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertisingPlan(c, advertisingPlanModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AdvertisingPlanRequest{}
	bindOk := requests.Validate(c, &request, requests.AdvertisingPlanSave)
	if !bindOk {
		return
	}

	advertisingPlanModel.Name = request.Name
	advertisingPlanModel.CreatorId = request.CreatorId
	advertisingPlanModel.AdvertisingId = request.AdvertisingId
	advertisingPlanModel.AdvertisingType = request.AdvertisingType
	advertisingPlanModel.AdvertisingPositionId = request.AdvertisingPositionId
	advertisingPlanModel.Order = request.Order
	advertisingPlanModel.SchedulingDate = request.SchedulingDate
	advertisingPlanModel.SchedulingTime = request.SchedulingTime
	advertisingPlanModel.StartDate = request.StartDate
	advertisingPlanModel.EndDate = request.EndDate
	advertisingPlanModel.StartTime = request.StartTime
	advertisingPlanModel.EndTime = request.EndTime
	advertisingPlanModel.AuditStatus = request.AuditStatus
	advertisingPlanModel.PresentStatus = request.PresentStatus

	rowsAffected := advertisingPlanModel.Save()
	if rowsAffected > 0 {
		response.Data(c, advertisingPlanModel)
	} else {
		response.Abort500(c, "??????????????????????????????~")
	}
}

func (ctrl *AdvertisingPlansController) Delete(c *gin.Context) {

	advertisingPlanModel := advertising_plan.Get(c.Param("id"))
	if advertisingPlanModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertisingPlan(c, advertisingPlanModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := advertisingPlanModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "??????????????????????????????~")
}

func (ctrl *AdvertisingPlansController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	advertisingPlanModel := advertising_plan.AdvertisingPlan{}
	if ok := policies.CanModifyAdvertisingPlan(c, advertisingPlanModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := advertisingPlanModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}

//????????????
func (ctrl *AdvertisingPlansController) Export(c *gin.Context) {

	listData := advertising_plan.All2()
	f := excelize.NewFile() // ?????????????????????
	//// ??????????????????
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "????????????")
	f.SetCellValue("Sheet1", "C1", "???????????????")
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

	//fruits := getFruits()
	// ??????????????????
	for _, v := range listData {
		line++
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", line), v.ID)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", line), v.Name)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", line), v.CreatorId)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", line), v.AdvertisingId)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", line), v.AdvertisingType)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", line), v.AdvertisingPositionId)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", line), v.Order)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", line), v.SchedulingDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", line), v.SchedulingTime)
		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", line), v.StartDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", line), v.EndDate)
		f.SetCellValue("Sheet1", fmt.Sprintf("L%d", line), v.StartTime)
		f.SetCellValue("Sheet1", fmt.Sprintf("M%d", line), v.EndTime)
		f.SetCellValue("Sheet1", fmt.Sprintf("N%d", line), v.AuditStatus)
		f.SetCellValue("Sheet1", fmt.Sprintf("O%d", line), v.PresentStatus)
	}

	var fileName = utils.RandFileName()
	var fullPath = "D:/" + fileName + ".xlsx"

	// ????????????
	if err := f.SaveAs(fullPath); err != nil {
		fmt.Println(err)
	}

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;fileName=%s", fileName))
	c.Writer.Header().Add("Content-Type", "application/octet-stream;charset=utf-8")

	c.File(fullPath)
}
