package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type PermissionGroup struct {
		models.BaseModel

		Name        string `gorm:"type:varchar(120);not null;comment:权限组名称"`
		Description string `gorm:"type:varchar(255);null;comment:权限组说明"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&PermissionGroup{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&PermissionGroup{})
	}

	migrate.Add("2022_03_24_105140_create_permission_groups_table", up, down)
}
