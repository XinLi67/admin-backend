package assemblies

import (
	"gohub/app/models/permission"

	"github.com/golang-module/carbon/v2"
)

type PermissionAssembly struct {
	ID                uint64 `json:"id"`
	PermissionGroupId uint64 `json:"permission_group_id"`
	Name              string `json:"name"`
	Icon              string `json:"icon"`
	GuardName         string `json:"guard_name"`
	DisplayName       string `json:"display_name"`
	Description       string `json:"description"`
	Sequence          uint64 `json:"sequence"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`

	PermissionGroup PermissionGroupAssembly `json:"group"`
}

func PermissionAssemblyFromModel(data permission.Permission) *PermissionAssembly {
	return &PermissionAssembly{
		ID:                data.ID,
		PermissionGroupId: data.PermissionGroupId,
		Name:              data.Name,
		Icon:              data.Icon,
		GuardName:         data.GuardName,
		DisplayName:       data.DisplayName,
		Description:       data.Description,
		Sequence:          data.Sequence,

		CreatedAt: carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt: carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),

		PermissionGroup: PermissionGroupAssembly{
			ID:        data.PermissionGroup.ID,
			Name:      data.PermissionGroup.Name,
			CreatedAt: carbon.Time2Carbon(data.PermissionGroup.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.PermissionGroup.UpdatedAt).ToDateTimeString(),
		},
	}
}

func PermissionAssemblyFromModelList(data []permission.Permission) interface{} {
	users := make([]PermissionAssembly, len(data))
	for i, v := range data {
		users[i] = PermissionAssembly{
			ID:                v.ID,
			PermissionGroupId: v.PermissionGroupId,
			Name:              v.Name,
			Icon:              v.Icon,
			GuardName:         v.GuardName,
			DisplayName:       v.DisplayName,
			Description:       v.Description,
			CreatedAt:         carbon.Time2Carbon(v.CreatedAt).ToDateTimeString(),
			UpdatedAt:         carbon.Time2Carbon(v.UpdatedAt).ToDateTimeString(),

			PermissionGroup: PermissionGroupAssembly{
				ID:        v.PermissionGroup.ID,
				Name:      v.PermissionGroup.Name,
				CreatedAt: carbon.Time2Carbon(v.PermissionGroup.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.PermissionGroup.UpdatedAt).ToDateTimeString(),
			},
		}
	}

	return users
}
