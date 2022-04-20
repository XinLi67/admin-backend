package v1

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gohub/app/http/assemblies"
	"gohub/app/models/advertising_plan"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
	"gohub/pkg/response"
	"gohub/utils"

	"github.com/gin-gonic/gin"
)

type AdvertisingPlansController struct {
	BaseAPIController
}

func (ctrl *AdvertisingPlansController) Index(c *gin.Context) {
	audit_status := c.Query("audit_status")
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	var  data []advertising_plan.AdvertisingPlan
	var pager paginator.Paging
	if len(audit_status)<=0{
		data, pager = advertising_plan.Paginate(c, 0)
	}else{
		data, pager = advertising_plan.Paginate2(c, 0,audit_status)
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
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

//todo 批量添加广告计划
func (ctrl *AdvertisingPlansController) BatchStore(c *gin.Context) {

	request := requests.AdvertigingBatchStoreRequest{}
	bindOk := requests.Validate(c, &request, requests.AdvertigingBatchStore)
	if !bindOk {
		return
	}

	var plans = request.Data
	database.DB.Create(&plans)

	for _, plan := range plans {
		fmt.Println("新增广告计划记录ID:", plan.ID)
	}

	response.Data(c, map[string]int{
		"新增广告计划数量": len(plans),
	})

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
		response.Abort500(c, "更新失败，请稍后尝试~")
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

	response.Abort500(c, "删除失败，请稍后尝试~")
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

//数据导出
func (ctrl *AdvertisingPlansController) Export(c *gin.Context) {

	listData := advertising_plan.All2()
	f := excelize.NewFile() // 设置单元格的值
	//// 这里设置表头
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "计划名称")
	f.SetCellValue("Sheet1", "C1", "创建者编号")
	f.SetCellValue("Sheet1", "D1", "广告ID")
	f.SetCellValue("Sheet1", "E1", "广告类型")
	f.SetCellValue("Sheet1", "F1", "广告位ID")
	f.SetCellValue("Sheet1", "G1", "排序号")
	f.SetCellValue("Sheet1", "H1", "排期天数")
	f.SetCellValue("Sheet1", "I1", "排期时间")
	f.SetCellValue("Sheet1", "J1", "开始日期")
	f.SetCellValue("Sheet1", "K1", "结束日期")
	f.SetCellValue("Sheet1", "L1", "开始时间")
	f.SetCellValue("Sheet1", "M1", "结束时间")
	f.SetCellValue("Sheet1", "N1", "审核状态")
	f.SetCellValue("Sheet1", "O1", "当前状态")

	line := 1

	//fruits := getFruits()
	// 循环写入数据
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
	var fullPath = "G:/studyFile/" + fileName + ".xlsx"

	// 保存文件
	if err := f.SaveAs(fullPath); err != nil {
		fmt.Println(err)
	}

	response.Data(c, "文件保存为:"+fullPath)
}
