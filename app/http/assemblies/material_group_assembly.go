package assemblies

import (
	"gohub/app/models/material_group"

	"github.com/golang-module/carbon/v2"
)

type MaterialGroupAssembly struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ParentId    uint64 `json:"parent_id"`
	Path        string `json:"path"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func MaterialGroupAssemblyFromModel(data *material_group.MaterialGroup) *MaterialGroupAssembly {
	return &MaterialGroupAssembly{
		ID:          data.ID,
		Name:        data.Name,
		ParentId:    data.ParentId,
		Path:        data.Path,
		Description: data.Description,
		CreatedAt:   carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt:   carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),
	}
}

func MaterialGroupAssemblyFromModelList(data []material_group.MaterialGroup) interface{} {
	MaterialGroups := make([]MaterialGroupAssembly, len(data))
	for i, v := range data {
		MaterialGroups[i] = MaterialGroupAssembly{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			ParentId:    v.ParentId,
			Path:        v.Path,
			CreatedAt:   carbon.Time2Carbon(v.CreatedAt).ToDateTimeString(),
			UpdatedAt:   carbon.Time2Carbon(v.UpdatedAt).ToDateTimeString(),
		}
	}

	return MaterialGroups
}
