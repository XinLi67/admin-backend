//Package permission_group 模型
package permission_group

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type PermissionGroup struct {
	models.BaseModel

	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`

	models.CommonTimestampsField
}

func (permissionGroup *PermissionGroup) Create() {
	database.DB.Create(&permissionGroup)
}

func (permissionGroup *PermissionGroup) Save() (rowsAffected int64) {
	result := database.DB.Save(&permissionGroup)
	return result.RowsAffected
}

func (permissionGroup *PermissionGroup) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&permissionGroup)
	return result.RowsAffected
}
