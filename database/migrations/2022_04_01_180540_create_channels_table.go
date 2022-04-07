package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Channel struct {
		models.BaseModel

		Status uint64 `gorm:"type:tinyint unsigned;column:status"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Channel{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Channel{})
	}

	migrate.Add("2022_04_01_180540_create_channels_table", up, down)
}
