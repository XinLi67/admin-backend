package v1

import (
	"gohub/app/http/assemblies"
	"gohub/app/models/announcement"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type AnnouncementsController struct {
	BaseAPIController
}

func (ctrl *AnnouncementsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := announcement.Paginate(c, 0)
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
