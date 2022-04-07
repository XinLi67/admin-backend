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
		CreatorId              uint64 `gorm:"type:bigint unsigned;column:creator_id"`
		DepartmentId           uint64 `gorm:"type:bigint unsigned;column:department_id"`
		Title                  string `gorm:"type:varchar(255);column:title"`
		LongTitle              string `gorm:"type:varchar(255);column:long_title"`
		Type                   uint64 `gorm:"type:tinyint unsigned;column:type"`
		Banner                 string `gorm:"type:varchar(255);column:banner"`
		RedirectTo             uint64 `gorm:"type:tinyint unsigned;column:redirect_to"`
		RedirectParams         string `gorm:"type:varchar(60);column:redirect_params"`
		Content                string `gorm:"type:text;column:content"`
		Status                 uint64 `gorm:"type:tinyint unsigned;column:status"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Announcement{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Announcement{})
	}

	migrate.Add("2022_04_01_170306_create_announcements_table", up, down)
}
