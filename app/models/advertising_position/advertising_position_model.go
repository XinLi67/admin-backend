//Package advertising_position 模型
package advertising_position

import (
	"gohub/app/models"
	"gohub/app/models/channel"
	"gohub/pkg/database"
)

type AdvertisingPosition struct {
	models.BaseModel

	Name        string `gorm:"column:name"`
	ChannelId   uint64 `gorm:"column:channel_id"`
	Code        string `gorm:"column:code"`
	Height      uint64 `gorm:"column:height"`
	Weight      uint64 `gorm:"column:weight"`
	Status      uint64 `gorm:"column:status"`
	Description string `gorm:"column:description"`

	Channel channel.Channel `json:"channel"`

	models.CommonTimestampsField
}

func (advertisingPosition *AdvertisingPosition) Create() {
	database.DB.Create(&advertisingPosition)
}

func (advertisingPosition *AdvertisingPosition) Save() (rowsAffected int64) {
	result := database.DB.Save(&advertisingPosition)
	return result.RowsAffected
}

func (advertisingPosition *AdvertisingPosition) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&advertisingPosition)
	return result.RowsAffected
}

func (advertisingPosition *AdvertisingPosition) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&advertisingPosition, ids)
	return result.RowsAffected
}
