//Package permission 模型
package permission

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type Permission struct {
	models.BaseModel

	PermissionGroupId uint64 `gorm:"column:permission_group_id"`
	Name              string `gorm:"column:name"`
	Icon              string `gorm:"column:icon"`
	GuardName         string `gorm:"column:guard_name"`
	DisplayName       string `gorm:"column:display_name"`
	Description       string `gorm:"column:description"`
	Sequence          uint64 `gorm:"column:sequence"`

	models.CommonTimestampsField
}

func (permission *Permission) Create() {
	database.DB.Create(&permission)
}

func (permission *Permission) Save() (rowsAffected int64) {
	result := database.DB.Save(&permission)
	return result.RowsAffected
}

func (permission *Permission) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&permission)
	return result.RowsAffected
}
