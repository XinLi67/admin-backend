package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type RolePermission struct {
		models.BaseModel

		RoleId       uint64 `gorm:"type:bigint unsigned;not null;index;comment:角色ID"`
		PermissionId uint64 `gorm:"type:bigint unsigned;not null;index;comment:权限ID"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&RolePermission{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&RolePermission{})
	}

	migrate.Add("2022_03_24_112704_create_role_permissions_table", up, down)
}
