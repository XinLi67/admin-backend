//Package announcement 模型
package announcement

import (
	"gohub/app/models"
	"gohub/app/models/announcement_position"
	"gohub/app/models/user"
	"gohub/pkg/database"
)

type Announcement struct {
	models.BaseModel

	AnnouncementNo         uint64 `gorm:"column:announcement_no"`
	AnnouncementPositionId uint64 `gorm:"column:announcement_position_id"`
	CreatorId              uint64 `gorm:"column:creator_id"`
	DepartmentId           uint64 `gorm:"column:department_id"`
	Title                  string `gorm:"column:title"`
	LongTitle              string `gorm:"column:long_title"`
	Type                   uint64 `gorm:"column:type"`
	Banner                 string `gorm:"column:banner"`
	RedirectTo             uint64 `gorm:"column:redirect_to"`
	RedirectParams         string `gorm:"column:redirect_params"`
	Content                string `gorm:"column:content"`
	Status                 uint64 `gorm:"column:status"`

	User                 user.User                                  `json:"user" gorm:"foreignkey:id"`
	AnnouncementPosition announcement_position.AnnouncementPosition `json:"announcement_position"`

	models.CommonTimestampsField
}

func (announcement *Announcement) Create() {
	database.DB.Create(&announcement)
}

func (announcement *Announcement) Save() (rowsAffected int64) {
	result := database.DB.Save(&announcement)
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
