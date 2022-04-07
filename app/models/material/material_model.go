//Package material 模型
package material

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type Material struct {
	models.BaseModel

	CreatorId       uint64 `gorm:"column:creator_id"`
	MaterialGroupId uint64 `gorm:"column:material_group_id"`
	DepartmentId    uint64 `gorm:"column:department_id"`
	Type            uint64 `gorm:"column:type"`
	Url             string `gorm:"column:url"`
	Title           string `gorm:"column:title"`
	Content         string `gorm:"column:content"`

	models.CommonTimestampsField
}

func (material *Material) Create() {
	database.DB.Create(&material)
}

func (material *Material) Save() (rowsAffected int64) {
	result := database.DB.Save(&material)
	return result.RowsAffected
}

func (material *Material) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&material)
	return result.RowsAffected
}

func (material *Material) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&material, ids)
	return result.RowsAffected
}
