package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type ClickRecord struct {
		models.BaseModel

		AdvertisingId uint64 `gorm:"type:bigint unsigned;column:advertising_id"`
		CustomerId    uint64 `gorm:"type:bigint unsigned;column:customer_id"`
		BrowsingTime  uint64 `gorm:"type:bigint unsigned;column:browsing_time"`
		StartTime     string `gorm:"type:varchar(10);column:start_time"`
		EndTime       string `gorm:"type:varchar(10);column:end_time"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ClickRecord{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&ClickRecord{})
	}

	migrate.Add("2022_04_01_170404_create_click_records_table", up, down)
}