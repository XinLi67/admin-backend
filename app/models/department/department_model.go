//Package department 模型
package department

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type Department struct {
	models.BaseModel

	ParentId    uint64 `gorm:"column:parent_id" json:"parent_id"`
	Name        string `gorm:"column:name" json:"name"`
	Phone       string `gorm:"column:phone" json:"phone"`
	LinkMan     string `gorm:"column:link_man" json:"link_man"`
	Address     string `gorm:"column:address" json:"address"`
	Description string `gorm:"column:description" json:"description"`

	models.CommonTimestampsField
}

func (department *Department) Create() {
	database.DB.Create(&department)
}

func (department *Department) Save() (rowsAffected int64) {
	result := database.DB.Save(&department)
	return result.RowsAffected
}

func (department *Department) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&department)
	return result.RowsAffected
}

func (department *Department) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&department, ids)
	return result.RowsAffected
}
