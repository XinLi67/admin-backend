package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Role struct {
		models.BaseModel

		Name        string `gorm:"type:varchar(120);not null;comment:角色名称"`
		GuardName   string `gorm:"type:varchar(30);not null;comment:项目名称"`
		Description string `gorm:"type:varchar(255);null;comment:角色说明"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Role{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Role{})
	}

	migrate.Add("2022_03_24_105211_create_roles_table", up, down)
}
