//Package user 模型
package user

import (
	"gohub/app/models"
	"gohub/app/models/department"
	"gohub/pkg/database"
	"gohub/pkg/hash"
)

type User struct {
	models.BaseModel

	DepartmentId uint64 `gorm:"column:department_id"`
	GuardName    string `gorm:"column:guard_name"`
	Username     string `gorm:"column:username"`
	Name         string `gorm:"column:name"`
	Gender       uint64 `gorm:"column:gender"`
	Email        string `gorm:"column:email"`
	Phone        string `gorm:"column:phone"`
	Avatar       string `gorm:"column:avatar"`
	Password     string `gorm:"column:password"`
	Status       uint64 `gorm:"column:status"`

	Department department.Department `json:"department"`

	models.CommonTimestampsField
}

func (user *User) Create() {
	database.DB.Create(&user)
}

func (user *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&user)
	return result.RowsAffected
}

func (user *User) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&user)
	return result.RowsAffected
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (user *User) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&user, ids)
	return result.RowsAffected
}
