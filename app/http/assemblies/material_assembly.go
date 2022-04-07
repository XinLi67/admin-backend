package assemblies

import (
	"gohub/app/models/material"

	"github.com/golang-module/carbon/v2"
)

type MaterialAssembly struct {
	ID              uint64 `json:"id"`
	CreatorId       uint64 `json:"creator_id"`
	MaterialGroupId uint64 `json:"material_group_id"`
	DepartmentId    uint64 `json:"department_id"`
	Type            uint64 `json:"type"`
	Url             string `json:"url"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`

	MaterialGroup MaterialGroupAssembly `json:"group"`
}

func MaterialAssemblyFromModel(data material.Material) *MaterialAssembly {
	return &MaterialAssembly{
		ID:              data.ID,
		CreatorId:       data.CreatorId,
		MaterialGroupId: data.MaterialGroupId,
		DepartmentId:    data.DepartmentId,
		Type:            data.Type,
		Url:             data.Url,
		Title:           data.Title,
		Content:         data.Content,
		CreatedAt:       carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt:       carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),

		MaterialGroup: MaterialGroupAssembly{
			ID:        data.MaterialGroup.ID,
			Name:      data.MaterialGroup.Name,
			CreatedAt: carbon.Time2Carbon(data.MaterialGroup.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.MaterialGroup.UpdatedAt).ToDateTimeString(),
		},
	}
}

func MaterialAssemblyFromModelList(data []material.Material) interface{} {
	Materials := make([]MaterialAssembly, len(data))
	for i, v := range data {
		Materials[i] = MaterialAssembly{
			ID:              v.ID,
			CreatorId:       v.CreatorId,
			MaterialGroupId: v.MaterialGroupId,
			DepartmentId:    v.DepartmentId,
			Type:            v.Type,
			Url:             v.Url,
			Title:           v.Title,
			Content:         v.Content,
			CreatedAt:       carbon.Time2Carbon(v.CreatedAt).ToDateTimeString(),
			UpdatedAt:       carbon.Time2Carbon(v.UpdatedAt).ToDateTimeString(),

			MaterialGroup: MaterialGroupAssembly{
				ID:        v.MaterialGroup.ID,
				Name:      v.MaterialGroup.Name,
				CreatedAt: carbon.Time2Carbon(v.MaterialGroup.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.MaterialGroup.UpdatedAt).ToDateTimeString(),
			},
		}
	}

	return Materials
}
