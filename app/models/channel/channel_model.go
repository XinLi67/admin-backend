//Package channel 模型
package channel

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type Channel struct {
	models.BaseModel

	Name        string `gorm:"column:name" json:"name"`
	GuardName   string `gorm:"column:guard_name" json:"guard_name"`
	Description string `gorm:"column:description" json:"description"`

	models.CommonTimestampsField
}

func (channel *Channel) Create() {
	database.DB.Create(&channel)
}

func (channel *Channel) Save() (rowsAffected int64) {
	result := database.DB.Save(&channel)
	return result.RowsAffected
}

func (channel *Channel) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&channel)
	return result.RowsAffected
}
