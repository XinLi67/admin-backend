package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type AdvertisingPlan struct {
		models.BaseModel

		Name                  string `gorm:"type:varchar(60);column:name"`
		CreatorId             uint64 `gorm:"type:bigint unsigned;column:creator_id"`
		AdvertisingId         uint64 `gorm:"type:bigint unsigned;column:advertising_id"`
		AdvertisingType       uint64 `gorm:"type:tinyint unsigned;column:advertising_type"`
		AdvertisingPositionId uint64 `gorm:"type:bigint unsigned;column:advertising_position_id"`
		Order                 uint64 `gorm:"type:smallint unsigned;column:order"`
		SchedulingDate        uint64 `gorm:"type:tinyint unsigned;column:scheduling_date"`
		SchedulingTime        uint64 `gorm:"type:tinyint unsigned;column:scheduling_time"`
		StartDate             string `gorm:"type:varchar(10);column:start_date"`
		EndDate               string `gorm:"type:varchar(10);column:end_date"`
		StartTime             string `gorm:"type:varchar(10);column:start_time"`
		EndTime               string `gorm:"type:varchar(10);column:end_time"`
		AuditStatus           uint64 `gorm:"type:tinyint unsigned;column:audit_status"`
		PresentStatus         uint64 `gorm:"type:tinyint unsigned;column:present_status"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AdvertisingPlan{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AdvertisingPlan{})
	}

	migrate.Add("2022_04_01_170337_create_advertising_plans_table", up, down)
}
