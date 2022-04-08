package assemblies

import (
	"gohub/app/models/permission_group"

	"github.com/golang-module/carbon/v2"
)

type PermissionGroupAssembly struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func PermissionGroupAssemblyFromModel(data *permission_group.PermissionGroup) *PermissionGroupAssembly {
	return &PermissionGroupAssembly{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt: carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),
	}
}

func PermissionGroupAssemblyFromModelList(data []permission_group.PermissionGroup) interface{} {
	permissionGroups := make([]PermissionGroupAssembly, len(data))
	for i, v := range data {
		permissionGroups[i] = PermissionGroupAssembly{
			ID:        v.ID,
			Name:      v.Name,
			CreatedAt: carbon.Time2Carbon(v.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(v.UpdatedAt).ToDateTimeString(),
		}
	}

	return permissionGroups
}
