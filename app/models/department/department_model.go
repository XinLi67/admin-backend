//Package department 模型
package department

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type Department struct {
	models.BaseModel

	ParentId    uint64 `gorm:"column:parent_id"`
	Name        string `gorm:"column:name"`
	Phone       string `gorm:"column:phone"`
	LinkMan     string `gorm:"column:link_man"`
	Address     string `gorm:"column:address"`
	Description string `gorm:"column:description"`

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
