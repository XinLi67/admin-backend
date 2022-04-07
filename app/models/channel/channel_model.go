//Package channel 模型
package channel

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type Channel struct {
	models.BaseModel

	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Status      uint64 `gorm:"column:status"`

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

func (channel *Channel) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&channel, ids)
	return result.RowsAffected
}
