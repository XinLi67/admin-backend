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

		Name        string `gorm:"type:varchar(120);not null;comment:角色名称"`
		GuardName   string `gorm:"type:varchar(30);not null;comment:项目名称"`
		Description string `gorm:"type:varchar(255);null;comment:角色说明"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Channel{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Channel{})
	}

	migrate.Add("2022_03_29_151856_create_channels_table", up, down)
}
