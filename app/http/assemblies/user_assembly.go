package assemblies

import (
	"gohub/app/models/user"

	"github.com/golang-module/carbon/v2"
)

type UserAssembly struct {
	ID           uint64 `json:"id"`
	DepartmentId uint64 `json:"department_id"`
	UserName     string `json:"username"`
	Name         string `json:"name"`
	Gender       uint64 `json:"gender"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Avatar       string `json:"avatar"`
	Status       uint64 `json:"status,omitempty"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`

	Department DepartmentAssembly `json:"department"`
}

func UserAssemblyFromModel(data user.User) *UserAssembly {
	return &UserAssembly{
		ID:           data.ID,
		DepartmentId: data.DepartmentId,
		UserName:     data.Username,
		Name:         data.Name,
		Gender:       data.Gender,
		Email:        data.Email,
		Phone:        data.Phone,
		Avatar:       data.Avatar,
		Status:       data.Status,
		CreatedAt:    carbon.Time2Carbon(data.CreatedAt).ToDateTimeString(),
		UpdatedAt:    carbon.Time2Carbon(data.UpdatedAt).ToDateTimeString(),

		Department: DepartmentAssembly{
			ID:        data.Department.ID,
			Name:      data.Department.Name,
			CreatedAt: carbon.Time2Carbon(data.Department.CreatedAt).ToDateTimeString(),
			UpdatedAt: carbon.Time2Carbon(data.Department.UpdatedAt).ToDateTimeString(),
		},
	}
}

func UserAssemblyFromModelList(data []user.User, total int) interface{} {
	users := make([]UserAssembly, total)
	for i, v := range data {
		users[i] = UserAssembly{
			ID:           v.ID,
			DepartmentId: v.DepartmentId,
			UserName:     v.Username,
			Name:         v.Name,
			Gender:       v.Gender,
			Email:        v.Email,
			Phone:        v.Phone,
			Avatar:       v.Avatar,
			Status:       v.Status,
			CreatedAt:    carbon.Time2Carbon(v.CreatedAt).ToDateTimeString(),
			UpdatedAt:    carbon.Time2Carbon(v.UpdatedAt).ToDateTimeString(),

			Department: DepartmentAssembly{
				ID:        v.Department.ID,
				Name:      v.Department.Name,
				CreatedAt: carbon.Time2Carbon(v.Department.CreatedAt).ToDateTimeString(),
				UpdatedAt: carbon.Time2Carbon(v.Department.UpdatedAt).ToDateTimeString(),
			},
		}
	}

	return users
}
