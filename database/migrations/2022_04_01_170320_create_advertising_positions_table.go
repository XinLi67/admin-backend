package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type AdvertisingPosition struct {
		models.BaseModel

		Name        string `gorm:"type:varchar(60);column:name"`
		ChannelId   uint64 `gorm:"type:smallint unsigned;column:channel_id"`
		Code        string `gorm:"type:varchar(60);column:code"`
		Height      uint64 `gorm:"type:int unsigned;column:height"`
		Weight      uint64 `gorm:"type:int unsigned;column:weight"`
		Status      uint64 `gorm:"type:tinyint unsigned;column:status"`
		Description string `gorm:"type:text;column:description"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&AdvertisingPosition{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&AdvertisingPosition{})
	}

	migrate.Add("2022_04_01_170320_create_advertising_positions_table", up, down)
}
