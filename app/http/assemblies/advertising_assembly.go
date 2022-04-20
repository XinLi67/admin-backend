package assemblies

import (
	"gohub/app/models/advertising"

	"github.com/golang-module/carbon/v2"
)

type AdvertisingAssembly struct {
	ID                    uint64 `json:"id"`
	AdvertisingNo         uint64 `json:"advertising_no"`
	AdvertisingPositionId uint64 `json:"advertising_position_id"`
	CreatorId             uint64 `json:"creator_id"`
	DepartmentId          uint64 `json:"department_id"`
	Title                 string `json:"title"`
	Type                  uint64 `json:"type"`
	RedirectTo            uint64 `json:"redirect_to"`
	MaterialId            uint64 `json:"material_id"`
	MaterialType          uint64 `json:"material_type"`
	Size                  string `json:"size"`
	RedirectParams        string `json:"redirect_params"`
	Description           string `json:"description"`
	Status                uint64 `json:"status"`
	AuditReason           string `json:"audit_reason"`
	PushContent           string `json:"push_content"`
	PushTitle             string `json:"push_title"`
	AdvertisingCreativity string `json:"advertising_creativity"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`

	User                UserAssembly                `json:"user"`
	AdvertisingPosition AdvertisingPositionAssembly `json:"advertising_position"`
}

func AdvertisingAssemblyFromModel(data advertising.Advertising) *AdvertisingAssembly {
	return &AdvertisingAssembly{
		ID:                    data.ID,
		AdvertisingNo:         data.AdvertisingNo,
		AdvertisingPositionId: data.AdvertisingPositionId,
		CreatorId:             data.CreatorId,
		DepartmentId:          data.DepartmentId,
		Title:                 data.Title,
		Type:                  data.Type,
		RedirectTo:            data.RedirectTo,
		MaterialId:            data.MaterialId,
		MaterialType:          data.MaterialType,
		Size:                  data.Size,
		RedirectParams:        data.RedirectParams,
		Description:           data.Description,
		Status:                data.Status,
		AuditReason:		   data.AuditReason,
		PushContent:           data.PushContent,
		PushTitle:             data.PushTitle,
		AdvertisingCreativity: data.AdvertisingCreativity,
		CreatedAt:             carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt:             carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),

		AdvertisingPosition: AdvertisingPositionAssembly{
			ID:        data.AdvertisingPosition.ID,
			Name:      data.AdvertisingPosition.Name,
			CreatedAt: carbon.Time2Carbon(data.AdvertisingPosition.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.AdvertisingPosition.UpdatedAt).ToDateTimeString(),
		},
		User: UserAssembly{
			ID:        data.User.ID,
			UserName:  data.User.Username,
			CreatedAt: carbon.Time2Carbon(data.User.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.User.UpdatedAt).ToDateTimeString(),
		},
	}
}

func AdvertisingAssemblyFromModelList(data []advertising.Advertising, total int) interface{} {
	Advertisings := make([]AdvertisingAssembly, total)
	for i, v := range data {
		Advertisings[i] = AdvertisingAssembly{
			ID:                    v.ID,
			AdvertisingNo:         v.AdvertisingNo,
			AdvertisingPositionId: v.AdvertisingPositionId,
			CreatorId:             v.CreatorId,
			DepartmentId:          v.DepartmentId,
			Title:                 v.Title,
			Type:                  v.Type,
			RedirectTo:            v.RedirectTo,
			MaterialId:            v.MaterialId,
			MaterialType:          v.MaterialType,
			Size:                  v.Size,
			RedirectParams:        v.RedirectParams,
			Description:           v.Description,
			Status:                v.Status,
			AuditReason:		   v.AuditReason,
			PushContent:           v.PushContent,
			PushTitle:             v.PushTitle,
			AdvertisingCreativity: v.AdvertisingCreativity,
			CreatedAt:             carbon.Time2Carbon(v.CreatedAt).ToDateTimeString(),
			UpdatedAt:             carbon.Time2Carbon(v.UpdatedAt).ToDateTimeString(),

			AdvertisingPosition: AdvertisingPositionAssembly{
				ID:        v.AdvertisingPosition.ID,
				Name:      v.AdvertisingPosition.Name,
				CreatedAt: carbon.Time2Carbon(v.AdvertisingPosition.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.AdvertisingPosition.UpdatedAt).ToDateTimeString(),
			},
			User: UserAssembly{
				ID:        v.User.ID,
				UserName:  v.User.Username,
				CreatedAt: carbon.Time2Carbon(v.User.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.User.UpdatedAt).ToDateTimeString(),
			},
		}
	}

	return Advertisings
}
