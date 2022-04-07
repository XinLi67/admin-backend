package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"
	"time"

	"gorm.io/gorm"
)

func init() {

	type AnnouncementPlan struct {
		models.BaseModel

		Name                   string    `gorm:"type:varchar(60);column:name"`
		CreatorId              uint64    `gorm:"type:bigint unsigned;column:creator_id"`
		AnnouncementId         uint64    `gorm:"type:bigint unsigned;column:announcement_id"`
		AnnouncementType       uint64    `gorm:"type:tinyint unsigned;column:announcement_type"`
		AnnouncementPositionId uint64    `gorm:"type:bigint unsigned;column:announcement_position_id"`
		Order                  uint64    `gorm:"type:smallint unsigned;column:order"`
		SchedulingDate         uint64    `gorm:"type:tinyint unsigned;column:scheduling_date"`
		SchedulingTime         uint64    `gorm:"type:tinyint unsigned;column:scheduling_time"`
		StartDate              time.Time `gorm:"type:datetime(3);column:start_date"`
		EndTDate               time.Time `gorm:"type:datetime(3);column:end_date"`
		StartTime              time.Time `gorm:"type:datetime(3);column:start_time"`
		EndTime                time.Time `gorm:"type:datetime(3);column:end_time"`
		AuditStatus            uint64    `gorm:"type:tinyint unsigned;column:audit_status"`
		PresentStatus          uint64    `gorm:"type:tinyint unsigned;column:present_status"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AnnouncementPlan{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AnnouncementPlan{})
	}

	migrate.Add("2022_04_01_170344_create_announcement_plans_table", up, down)
}
