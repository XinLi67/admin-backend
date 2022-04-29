package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Announcement struct {
		models.BaseModel

		AnnouncementNo         uint64 `gorm:"type:bigint unsigned;column:announcement_no"`
		AnnouncementPositionId uint64 `gorm:"type:bigint unsigned;column:announcement_position_id"`
		UserId                 uint64 `gorm:"type:bigint unsigned;column:user_id"`
		Title                  string `gorm:"type:varchar(255);column:title"`
		Type                   uint64 `gorm:"type:tinyint unsigned;column:type"`
		ChannelId              uint64 `gorm:"type:bigint unsigned;column:channel_id"`
		RedirectTo             uint64 `gorm:"type:tinyint unsigned;column:redirect_to"`
		RedirectParams         string `gorm:"type:varchar(60);column:redirect_params"`
		Content                string `gorm:"type:text;column:content"`
		Status                 uint64 `gorm:"type:tinyint unsigned;column:status"`
		AuditReason            string `gorm:"type:text;column:audit_reason"`
		SchedulingType         uint64 `gorm:"type:tinyint unsigned;column:scheduling_type"`
		StartDate              string `gorm:"column:start_date;index;" json:"start_date,omitempty"`
		EndDate                string `gorm:"column:end_date;index;" json:"end_date,omitempty"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Announcement{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Announcement{})
	}

	migrate.Add("2022_04_29_171733_add_announcements_table", up, down)
}
