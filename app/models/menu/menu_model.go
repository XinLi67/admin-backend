//Package menu 模型
package menu

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type Menu struct {
	models.BaseModel

	ParentId       uint64 `gorm:"column:parent_id"`
	Name           string `gorm:"column:name"`
	Icon           string `gorm:"column:icon"`
	Uri            string `gorm:"column:uri"`
	IsLink         bool   `gorm:"column:is_link"`
	PermissionName string `gorm:"column:permission_name"`
	GuardName      string `gorm:"column:guard_name"`
	Sequence       uint64 `gorm:"column:sequence"`

	models.CommonTimestampsField
}

func (menu *Menu) Create() {
	database.DB.Create(&menu)
}

func (menu *Menu) Save() (rowsAffected int64) {
	result := database.DB.Save(&menu)
	return result.RowsAffected
}

func (menu *Menu) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&menu)
	return result.RowsAffected
}
