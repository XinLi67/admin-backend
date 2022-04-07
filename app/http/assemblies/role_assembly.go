package assemblies

import (
	"gohub/app/models/role"

	"github.com/golang-module/carbon/v2"
)

type RoleAssembly struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	GuardName   string `json:"guard_name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func RoleAssemblyFromModel(data role.Role) *RoleAssembly {
	return &RoleAssembly{
		ID:          data.ID,
		Name:        data.Name,
		GuardName:   data.GuardName,
		Description: data.Description,
		CreatedAt:   carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt:   carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),
	}
}

func RoleAssemblyFromModelList(data []role.Role) interface{} {
	Roles := make([]RoleAssembly, len(data))
	for i, v := range data {
		Roles[i] = RoleAssembly{
			ID:          v.ID,
			Name:        v.Name,
			GuardName:   v.GuardName,
			Description: v.Description,
			CreatedAt:   carbon.Time2Carbon(v.CreatedAt).ToDateTimeString(),
			UpdatedAt:   carbon.Time2Carbon(v.UpdatedAt).ToDateTimeString(),
		}
	}

	return Roles
}
