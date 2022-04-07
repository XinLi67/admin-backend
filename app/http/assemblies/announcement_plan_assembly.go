package assemblies

import (
	"gohub/app/models/announcement_plan"

	"github.com/golang-module/carbon/v2"
)

type AnnouncementPlanAssembly struct {
	ID                     uint64 `json:"id"`
	Name                   string `json:"name"`
	CreatorId              uint64 `json:"creator_id"`
	AnnouncementId         uint64 `json:"announcement_id"`
	AnnouncementType       uint64 `json:"announcement_type"`
	AnnouncementPositionId uint64 `json:"announcement_position_id"`
	Order                  uint64 `json:"order"`
	SchedulingDate         uint64 `json:"scheduling_date"`
	SchedulingTime         uint64 `json:"scheduling_time"`
	StartDate              string `json:"start_date"`
	EndDate                string `json:"end_date"`
	StartTime              string `json:"start_time"`
	Endime                 string `json:"end_time"`
	AuditStatus            uint64 `json:"audit_status"`
	PresentStatus          uint64 `json:"present_status"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`

	User                 UserAssembly                 `json:"user"`
	AnnouncementPosition AnnouncementPositionAssembly `json:"announcement_position"`
}

func AnnouncementPlanAssemblyFromModel(data announcement_plan.AnnouncementPlan) *AnnouncementPlanAssembly {
	return &AnnouncementPlanAssembly{
		ID:                     data.ID,
		Name:                   data.Name,
		CreatorId:              data.CreatorId,
		AnnouncementId:         data.AnnouncementId,
		AnnouncementType:       data.AnnouncementType,
		AnnouncementPositionId: data.AnnouncementPositionId,
		Order:                  data.Order,
		SchedulingDate:         data.SchedulingDate,
		SchedulingTime:         data.SchedulingTime,
		StartDate:              data.StartDate,
		EndDate:                data.EndDate,
		StartTime:              data.StartTime,
		Endime:                 data.Endime,
		AuditStatus:            data.AuditStatus,
		PresentStatus:          data.PresentStatus,
		CreatedAt:              carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt:              carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),

		User: UserAssembly{
			ID:        data.User.ID,
			UserName:  data.User.Username,
			CreatedAt: carbon.Time2Carbon(data.User.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.User.UpdatedAt).ToDateTimeString(),
		},

		AnnouncementPosition: AnnouncementPositionAssembly{
			ID:        data.AnnouncementPosition.ID,
			Name:      data.AnnouncementPosition.Name,
			CreatedAt: carbon.Time2Carbon(data.AnnouncementPosition.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.AnnouncementPosition.UpdatedAt).ToDateTimeString(),
		},
	}
}

func AnnouncementPlanAssemblyFromModelList(data []announcement_plan.AnnouncementPlan, total int) interface{} {
	AnnouncementPlans := make([]AnnouncementPlanAssembly, total)
	for i, v := range data {
		AnnouncementPlans[i] = AnnouncementPlanAssembly{
			ID:                     v.ID,
			Name:                   v.Name,
			CreatorId:              v.CreatorId,
			AnnouncementId:         v.AnnouncementId,
			AnnouncementType:       v.AnnouncementType,
			AnnouncementPositionId: v.AnnouncementPositionId,
			Order:                  v.Order,
			SchedulingDate:         v.SchedulingDate,
			SchedulingTime:         v.SchedulingTime,
			StartDate:              v.StartDate,
			EndDate:                v.EndDate,
			StartTime:              v.StartTime,
			Endime:                 v.Endime,
			AuditStatus:            v.AuditStatus,
			PresentStatus:          v.PresentStatus,
			CreatedAt:              carbon.Time2Carbon(v.CreatedAt).ToDateTimeString(),
			UpdatedAt:              carbon.Time2Carbon(v.UpdatedAt).ToDateTimeString(),

			User: UserAssembly{
				ID:        v.User.ID,
				UserName:  v.User.Username,
				CreatedAt: carbon.Time2Carbon(v.User.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.User.UpdatedAt).ToDateTimeString(),
			},

			AnnouncementPosition: AnnouncementPositionAssembly{
				ID:        v.AnnouncementPosition.ID,
				Name:      v.AnnouncementPosition.Name,
				CreatedAt: carbon.Time2Carbon(v.AnnouncementPosition.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.AnnouncementPosition.UpdatedAt).ToDateTimeString(),
			},
		}
	}

	return AnnouncementPlans
}
