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

func PeimissionGroupAssemblyFromModel(data *permission_group.PermissionGroup) *PermissionGroupAssembly {
	return &PermissionGroupAssembly{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt: carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),
	}
}
