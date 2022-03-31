package assemblies

import (
	"gohub/app/models/department"

	"github.com/golang-module/carbon/v2"
)

type DepartmentAssembly struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func DepartmentAssemblyFromModel(data *department.Department) *DepartmentAssembly {
	return &DepartmentAssembly{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt: carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),
	}
}
