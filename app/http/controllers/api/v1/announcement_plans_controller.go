package v1

import (
	"gohub/app/http/assemblies"
	"gohub/app/models/announcement_plan"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type AnnouncementPlansController struct {
	BaseAPIController
}

func (ctrl *AnnouncementPlansController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := announcement_plan.Paginate(c, 0)
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
		EndTDate:               request.EndTDate,
		StartTime:              request.StartTime,
		EndTime:                request.EndTime,
		AuditStatus:            request.AuditStatus,
		PresentStatus:          request.PresentStatus,
	}
	announcementPlanModel.Create()
	if announcementPlanModel.ID > 0 {
		response.Created(c, announcementPlanModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
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
	announcementPlanModel.EndTDate = request.EndTDate
	announcementPlanModel.StartTime = request.StartTime
	announcementPlanModel.EndTime = request.EndTime
	announcementPlanModel.AuditStatus = request.AuditStatus
	announcementPlanModel.PresentStatus = request.PresentStatus

	rowsAffected := announcementPlanModel.Save()
	if rowsAffected > 0 {
		response.Data(c, announcementPlanModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
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

	response.Abort500(c, "删除失败，请稍后尝试~")
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
