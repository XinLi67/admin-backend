//Package announcement_plan 模型
package announcement_plan

import (
	"gohub/app/models"
	"gohub/app/models/announcement_position"
	"gohub/app/models/user"
	"gohub/pkg/database"
	"time"
)

type AnnouncementPlan struct {
	models.BaseModel

	Name                   string    `gorm:"column:name"`
	CreatorId              uint64    `gorm:"column:creator_id"`
	AnnouncementId         uint64    `gorm:"column:announcement_id"`
	AnnouncementType       uint64    `gorm:"column:announcement_type"`
	AnnouncementPositionId uint64    `gorm:"column:announcement_position_id"`
	Order                  uint64    `gorm:"column:order"`
	SchedulingDate         uint64    `gorm:"column:scheduling_date"`
	SchedulingTime         uint64    `gorm:"column:scheduling_time"`
	StartDate              time.Time `gorm:"column:start_date"`
	EndTDate               time.Time `gorm:"column:end_date"`
	StartTime              time.Time `gorm:"column:start_time"`
	EndTime                time.Time `gorm:"column:end_time"`
	AuditStatus            uint64    `gorm:"column:audit_status"`
	PresentStatus          uint64    `gorm:"column:present_status"`

	User                 user.User                                  `json:"user"`
	AnnouncementPosition announcement_position.AnnouncementPosition `json:"announcement_position"`

	models.CommonTimestampsField
}

func (announcementPlan *AnnouncementPlan) Create() {
	database.DB.Create(&announcementPlan)
}

func (announcementPlan *AnnouncementPlan) Save() (rowsAffected int64) {
	result := database.DB.Save(&announcementPlan)
	return result.RowsAffected
}

func (announcementPlan *AnnouncementPlan) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&announcementPlan)
	return result.RowsAffected
}

func (announcementPlan *AnnouncementPlan) BatchDelete(ids []int) (rowsAffected int64) {
	result := database.DB.Delete(&announcementPlan, ids)
	return result.RowsAffected
}
