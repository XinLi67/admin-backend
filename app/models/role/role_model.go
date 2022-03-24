//Package role 模型
package role

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type Role struct {
	models.BaseModel

	Name        string `gorm:"column:name"`
	GuardName   string `gorm:"column:guard_name"`
	Description string `gorm:"column:description"`

	models.CommonTimestampsField
}

func (role *Role) Create() {
	database.DB.Create(&role)
}

func (role *Role) Save() (rowsAffected int64) {
	result := database.DB.Save(&role)
	return result.RowsAffected
}

func (role *Role) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&role)
	return result.RowsAffected
}
