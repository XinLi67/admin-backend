//Package material_group 模型
package material_group

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type MaterialGroup struct {
	models.BaseModel

	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	ParentId    uint64 `gorm:"column:parent_id"`
	Path        string `gorm:"column:path" `
	models.CommonTimestampsField
}

func (materialGroup *MaterialGroup) Create() {
	database.DB.Create(&materialGroup)
}

func (materialGroup *MaterialGroup) Save() (rowsAffected int64) {
	result := database.DB.Save(&materialGroup)
	return result.RowsAffected
}

func (materialGroup *MaterialGroup) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&materialGroup)
	return result.RowsAffected
}

func (materialGroup *MaterialGroup) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&materialGroup, ids)
	return result.RowsAffected
}
