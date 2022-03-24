package dto

import "gohub/app/models/user"

type UserDTO struct {
	ID           uint64 `json:"id"`
	DepartmentId uint64 `json:"department_id"`
	UserName     string `json:"username"`
	Name         string `json:"name"`
	Gender       uint64 `json:"gender"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Avatar       string `json:"avatar"`
	Status       uint64 `json:"status,omitempty"`
}

func UserDTOFromModel(data *user.User) *UserDTO {
	return &UserDTO{
		ID:           data.ID,
		DepartmentId: data.DepartmentId,
		UserName:     data.Username,
		Name:         data.Name,
		Gender:       data.Gender,
		Email:        data.Email,
		Phone:        data.Phone,
		Avatar:       data.Avatar,
		Status:       data.Status,
	}
}

func UserDTOFromModelList(data []user.User, total int64) interface{} {
	users := make([]UserDTO, total)
	for i, v := range data {
		users[i] = UserDTO{
			ID:           v.ID,
			DepartmentId: v.DepartmentId,
			UserName:     v.Username,
			Name:         v.Name,
			Gender:       v.Gender,
			Email:        v.Email,
			Phone:        v.Phone,
			Avatar:       v.Avatar,
			Status:       v.Status,
		}
	}

	return users
}
