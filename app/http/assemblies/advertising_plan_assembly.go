package assemblies

import (
	"gohub/app/models/advertising_plan"

	"github.com/golang-module/carbon/v2"
)

type AdvertisingPlanAssembly struct {
	ID                    uint64 `json:"id"`
	Name                  string `json:"name"`
	CreatorId             uint64 `json:"creator_id"`
	AdvertisingId         uint64 `json:"ajdvertising_id"`
	AdvertisingType       uint64 `json:"advertising_type"`
	AdvertisingPositionId uint64 `json:"advertising_position_id"`
	Order                 uint64 `json:"order"`
	SchedulingDate        uint64 `json:"scheduling_date"`
	SchedulingTime        uint64 `json:"scheduling_time"`
	StartDate             string `json:"start_date"`
	EndDate               string `json:"end_date"`
	StartTime             string `json:"start_time"`
	EndTime               string `json:"end_time"`
	AuditStatus           uint64 `json:"audit_status"`
	PresentStatus         uint64 `json:"present_status"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`

	User                UserAssembly                `json:"user"`
	AdvertisingPosition AdvertisingPositionAssembly `json:"advertising_position"`
}

func AdvertisingPlanAssemblyFromModel(data advertising_plan.AdvertisingPlan) *AdvertisingPlanAssembly {
	return &AdvertisingPlanAssembly{
		ID:                    data.ID,
		Name:                  data.Name,
		CreatorId:             data.CreatorId,
		AdvertisingId:         data.AdvertisingId,
		AdvertisingType:       data.AdvertisingType,
		AdvertisingPositionId: data.AdvertisingPositionId,
		Order:                 data.Order,
		SchedulingDate:        data.SchedulingDate,
		SchedulingTime:        data.SchedulingTime,
		StartDate:             data.StartDate,
		EndDate:               data.EndDate,
		StartTime:             data.StartTime,
		EndTime:               data.EndTime,
		AuditStatus:           data.AuditStatus,
		PresentStatus:         data.PresentStatus,
		CreatedAt:             carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt:             carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),

		User: UserAssembly{
			ID:        data.User.ID,
			UserName:  data.User.Username,
			CreatedAt: carbon.Time2Carbon(data.User.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.User.UpdatedAt).ToDateTimeString(),
		},

		AdvertisingPosition: AdvertisingPositionAssembly{
			ID:        data.AdvertisingPosition.ID,
			Name:      data.AdvertisingPosition.Name,
			CreatedAt: carbon.Time2Carbon(data.AdvertisingPosition.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.AdvertisingPosition.UpdatedAt).ToDateTimeString(),
		},
	}
}

func AdvertisingPlanAssemblyFromModelList(data []advertising_plan.AdvertisingPlan, total int) interface{} {
	AdvertisingPlans := make([]AdvertisingPlanAssembly, total)
	for i, v := range data {
		AdvertisingPlans[i] = AdvertisingPlanAssembly{
			ID:                    v.ID,
			Name:                  v.Name,
			CreatorId:             v.CreatorId,
			AdvertisingId:         v.AdvertisingId,
			AdvertisingType:       v.AdvertisingType,
			AdvertisingPositionId: v.AdvertisingPositionId,
			Order:                 v.Order,
			SchedulingDate:        v.SchedulingDate,
			SchedulingTime:        v.SchedulingTime,
			StartDate:             v.StartDate,
			EndDate:               v.EndDate,
			StartTime:             v.StartTime,
			EndTime:               v.EndTime,
			AuditStatus:           v.AuditStatus,
			PresentStatus:         v.PresentStatus,
			CreatedAt:             carbon.Time2Carbon(v.CreatedAt).ToDateTimeString(),
			UpdatedAt:             carbon.Time2Carbon(v.UpdatedAt).ToDateTimeString(),

			User: UserAssembly{
				ID:        v.User.ID,
				UserName:  v.User.Username,
				CreatedAt: carbon.Time2Carbon(v.User.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.User.UpdatedAt).ToDateTimeString(),
			},

			AdvertisingPosition: AdvertisingPositionAssembly{
				ID:        v.AdvertisingPosition.ID,
				Name:      v.AdvertisingPosition.Name,
				CreatedAt: carbon.Time2Carbon(v.AdvertisingPosition.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.AdvertisingPosition.UpdatedAt).ToDateTimeString(),
			},
		}
	}

	return AdvertisingPlans
}
