//Package announcement 模型
package announcement

import (
	"gohub/app/models"
	"gohub/app/models/announcement_position"
	"gohub/app/models/channel"
	"gohub/app/models/user"
	"gohub/pkg/database"
)

type Announcement struct {
	models.BaseModel

	AnnouncementNo         uint64 `gorm:"column:announcement_no"`
	AnnouncementPositionId uint64 `gorm:"column:announcement_position_id"`
	ChannelId              uint64 `gorm:"column:channel_id"`
	UserId                 uint64 `gorm:"column:user_id"`
	Title                  string `gorm:"column:title"`
	Type                   uint64 `gorm:"column:type"`

	RedirectTo           uint64                                      `gorm:"column:redirect_to"`
	RedirectParams       string                                      `gorm:"column:redirect_params"`
	Content              string                                      `gorm:"column:content"`
	Status               uint64                                      `gorm:"column:status"`
	AuditReason          string                                      `gorm:"column:audit_reason"`
	SchedulingType       uint64                                      `gorm:"column:scheduling_type"`
	StartDate            string                                      `gorm:"column:start_date;index;" json:"start_date,omitempty"`
	EndDate              string                                      `gorm:"column:end_date;index;" json:"end_date,omitempty"`
	User                 *user.User                                  `json:"user" gorm:"foreignkey:id;references:UserId"`
	AnnouncementPosition *announcement_position.AnnouncementPosition `json:"announcement_position"`
	Channel              *channel.Channel                            `json:"channel"`
	models.CommonTimestampsField
}

func (announcement *Announcement) Create() {
	database.DB.Create(&announcement)
}

func (announcement *Announcement) Save() (rowsAffected int64) {
	result := database.DB.Save(&announcement)
	return result.RowsAffected
}

//更新审核状态
func (announcement *Announcement) UpdateStatus(status string, auditReason string) (rowsAffected int64) {
	result := database.DB.Model(&announcement).Updates(map[string]interface{}{"status": status, "audit_reason": auditReason})
	return result.RowsAffected
}

func (announcement *Announcement) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&announcement)
	return result.RowsAffected
}

func (announcement *Announcement) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&announcement, ids)
	return result.RowsAffected
}
