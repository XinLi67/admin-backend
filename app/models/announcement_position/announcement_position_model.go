//Package announcement_position 模型
package announcement_position

import (
	"gohub/app/models"
	"gohub/app/models/channel"
	"gohub/pkg/database"
)

type AnnouncementPosition struct {
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

func (announcementPosition *AnnouncementPosition) Create() {
	database.DB.Create(&announcementPosition)
}

func (announcementPosition *AnnouncementPosition) Save() (rowsAffected int64) {
	result := database.DB.Save(&announcementPosition)
	return result.RowsAffected
}

func (announcementPosition *AnnouncementPosition) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&announcementPosition)
	return result.RowsAffected
}

func (announcementPosition *AnnouncementPosition) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&announcementPosition, ids)
	return result.RowsAffected
}
