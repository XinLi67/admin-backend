//Package advertising 模型
package advertising

import (
	"gohub/app/models"
	"gohub/app/models/advertising_position"
	"gohub/app/models/channel"
	"gohub/app/models/user"
	"gohub/pkg/database"
)

type Advertising struct {
	models.BaseModel

	AdvertisingNo         uint64 `gorm:"column:advertising_no"`
	AdvertisingPositionId uint64 `gorm:"column:advertising_position_id"`
	CreatorId             uint64 `gorm:"column:creator_id"`
	DepartmentId          uint64 `gorm:"column:department_id"`
	ChannelId             uint64 `gorm:"column:channel_id"`
	Title                 string `gorm:"column:title"`
	Type                  uint64 `gorm:"column:type"`
	RedirectTo            uint64 `gorm:"column:redirect_to"`
	MaterialId            uint64 `gorm:"column:material_id"`
	MaterialType          uint64 `gorm:"column:material_type"`
	Size                  string `gorm:"column:size"`
	RedirectParams        string `gorm:"column:redirect_params"`
	Description           string `gorm:"column:description"`
	Status                uint64 `gorm:"column:status"`
	AuditReason           string `gorm:"column:audit_reason"`
	PushContent           string `gorm:"column:push_content"`
	PushTitle             string `gorm:"column:push_title"`
	AdvertisingCreativity string `gorm:"column:advertising_creativity"`
	StartTime             string `gorm:"type:varchar(20);column:start_time"`
	EndTime               string `gorm:"type:varchar(20);column:end_time"`
	SchedulingTime        uint64 `gorm:"column:scheduling_time"`
	Url                   string `gorm:"type:varchar(255);column:url"`
	Url2                  string `gorm:"type:varchar(255);column:url2"`
	Url3                  string `gorm:"type:varchar(255);column:url3"`

	User                *user.User                                `json:"user" gorm:"foreignkey:id;references:CreatorId"`
	AdvertisingPosition *advertising_position.AdvertisingPosition `json:"advertising_position"`
	Channel             *channel.Channel                          `json:"channel"`

	models.CommonTimestampsField
}

func (advertising *Advertising) Create() {
	database.DB.Create(&advertising)
}

func (advertising *Advertising) Save() (rowsAffected int64) {
	result := database.DB.Save(&advertising)
	return result.RowsAffected
}

func (advertising *Advertising) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&advertising)
	return result.RowsAffected
}

func (advertising *Advertising) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&advertising, ids)
	return result.RowsAffected
}
